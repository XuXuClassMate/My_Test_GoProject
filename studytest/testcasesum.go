package main

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type Topic struct {
	Title    string  `xml:"title"`
	Children []Topic `xml:"children>topics>topic"`
}

type Sheet struct {
	RootTopic Topic `xml:"topic"`
}

type Workbook struct {
	Sheets []Sheet `xml:"sheet"`
}

// Recursively count the number of leaf (end) topics and classify them by priority
func countTopicsByPriority(topic Topic, p0Count *int, p1Count *int) int {
	// If the current topic has no children, it's a leaf node
	if len(topic.Children) == 0 {
		if strings.Contains(strings.ToLower(topic.Title), "p0") {
			*p0Count++
		} else if strings.Contains(strings.ToLower(topic.Title), "p1") {
			*p1Count++
		}
		return 1
	}
	// Otherwise, recursively count leaf nodes in its children
	totalCount := 0
	for _, child := range topic.Children {
		totalCount += countTopicsByPriority(child, p0Count, p1Count)
	}
	return totalCount
}

func countTestCasesByPriority(xmindFile string) (int, int, int, error) {
	// Open the XMind zip file
	zipReader, err := zip.OpenReader(xmindFile)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("failed to open xmind file: %w", err)
	}
	defer zipReader.Close()

	var contentXML []byte

	// Extract content.xml from the zip file
	for _, file := range zipReader.File {
		if file.Name == "content.xml" {
			f, err := file.Open()
			if err != nil {
				return 0, 0, 0, fmt.Errorf("failed to open content.xml: %w", err)
			}
			defer f.Close()

			contentXML, err = io.ReadAll(f)
			if err != nil {
				return 0, 0, 0, fmt.Errorf("failed to read content.xml: %w", err)
			}
			break
		}
	}

	if contentXML == nil {
		return 0, 0, 0, fmt.Errorf("content.xml not found in xmind file")
	}

	// Parse content.xml
	var workbook Workbook
	if err := xml.Unmarshal(contentXML, &workbook); err != nil {
		return 0, 0, 0, fmt.Errorf("failed to parse content.xml: %w", err)
	}

	// Count all leaf topics by priority in all sheets
	totalCount, p0Count, p1Count := 0, 0, 0
	for _, sheet := range workbook.Sheets {
		totalCount += countTopicsByPriority(sheet.RootTopic, &p0Count, &p1Count)
	}

	return totalCount, p0Count, p1Count, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <xmind-file>")
		return
	}

	xmindFile := os.Args[1]
	totalCount, p0Count, p1Count, err := countTestCasesByPriority(xmindFile)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Total test cases: %d, including P0 test cases: %d and P1 test cases: %d\n", totalCount, p0Count, p1Count)
}
