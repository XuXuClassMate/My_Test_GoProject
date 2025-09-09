# Test case statistics tool

## 1、How to compile:
```
go build testcasesum.go
```

## 2. How to run and log output

```
➜  TestCaseStatisticsTool git:(main) ✗ ./testcasesum 测试点统计格式Demo.xmind
15:20:21 Start parsing XMind file: 测试点统计格式Demo.xmind
15:20:21 Found content.json and is parsing it....

======== Test Case Summary ========
📊 Total test cases: 11
🔥 P0 test cases: 1 (9.1%)
⚡ P1 test cases: 2 (18.2%)
===================================
15:20:21 Done in 0.002s

```
