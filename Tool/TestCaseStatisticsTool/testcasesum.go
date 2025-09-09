package main

import (
	"archive/zip"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// ---------------- Log Configuration ----------------
func init() {
	log.SetFlags(log.Ltime) // The log only displays the time
}

// ---------------- XMind JSON structure ----------------
type XMindJSON struct {
	RootTopic *Topic `json:"rootTopic"`
}

type Topic struct {
	Title    string   `json:"title"`
	Children Children `json:"children"`
}

type Children struct {
	Attached []*Topic `json:"attached"`
}

// ---------------- XML Structure (Old Versions of XMind) ----------------
type XMap struct {
	Sheets []Sheet `xml:"sheet"`
}
type Sheet struct {
	Topic TopicXML `xml:"topic"`
}
type TopicXML struct {
	Title  string     `xml:"title"`
	Topics []TopicXML `xml:"topics>topic"`
}

// ---------------- Statistical Logic ----------------
func countJSONTopic(t *Topic, total, p0, p1 *int) {
	if t == nil {
		return
	}
	// Leaf node: no child nodes
	if len(t.Children.Attached) == 0 {
		*total++
		title := strings.ToLower(t.Title)
		if strings.Contains(title, "p0") {
			*p0++
		} else if strings.Contains(title, "p1") {
			*p1++
		}
	}
	for _, child := range t.Children.Attached {
		countJSONTopic(child, total, p0, p1)
	}
}

func countXMLTopic(t TopicXML, total, p0, p1 *int) {
	if len(t.Topics) == 0 {
		*total++
		title := strings.ToLower(t.Title)
		if strings.Contains(title, "p0") {
			*p0++
		} else if strings.Contains(title, "p1") {
			*p1++
		}
	}
	for _, child := range t.Topics {
		countXMLTopic(child, total, p0, p1)
	}
}

// ---------------- Main process ----------------
func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide the path to the .xmind file")
	}
	xmindFile := os.Args[1]
	start := time.Now()

	log.Printf("Start parsing XMind file: %s", filepath.Base(xmindFile))

	r, err := zip.OpenReader(xmindFile)
	if err != nil {
		log.Fatalf("Unable to open XMind file: %v", err)
	}
	defer r.Close()

	var total, p0, p1 int
	found := false

	for _, f := range r.File {
		if filepath.Base(f.Name) == "content.json" {
			found = true
			log.Println("Found content.json and is parsing it....")
			rc, _ := f.Open()
			data, _ := io.ReadAll(rc)
			rc.Close()

			var arr []XMindJSON
			if err := json.Unmarshal(data, &arr); err != nil {
				log.Fatalf("Parsing JSON failed: %v", err)
			}
			for _, sheet := range arr {
				if sheet.RootTopic != nil {
					countJSONTopic(sheet.RootTopic, &total, &p0, &p1)
				}
			}
			break
		}
	}

	if !found {
		for _, f := range r.File {
			if filepath.Base(f.Name) == "content.xml" {
				log.Println("Found content.json, parsing...")
				rc, _ := f.Open()
				data, _ := io.ReadAll(rc)
				rc.Close()

				var xm XMap
				if err := xml.Unmarshal(data, &xm); err != nil {
					log.Fatalf("Failed to parse XML: %v", err)
				}
				for _, sheet := range xm.Sheets {
					countXMLTopic(sheet.Topic, &total, &p0, &p1)
				}
				break
			}
		}
	}

	// helper for percent string
	getPercent := func(part, total int) string {
		if total == 0 {
			return "0.0%"
		}
		return fmt.Sprintf("%.1f%%", float64(part)*100.0/float64(total))
	}

	// Print colored summary
	fmt.Println()
	fmt.Println("======== Test Case Summary ========")
	fmt.Printf("📊 Total test cases: \033[1;34m%d\033[0m\n", total)
	fmt.Printf("🔥 P0 test cases: \033[1;31m%d\033[0m (%s)\n", p0, getPercent(p0, total))
	fmt.Printf("⚡ P1 test cases: \033[1;33m%d\033[0m (%s)\n", p1, getPercent(p1, total))
	fmt.Println("===================================")

	log.Printf("Done in %.3fs", time.Since(start).Seconds())

}
