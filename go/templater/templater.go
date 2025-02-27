package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
)

var header = `<!DOCTYPE html>
<html>
<head>
    <title>%s Report</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        table { width: 100% ; border-collapse: collapse; margin-top: 20px; }
        th, td { border: 1px solid #ddd; padding: 10px; text-align: left; }
        th { background-color: #f4f4f4; }
    </style>
</head>
<body>
    <h2>%s Report</h2>
    <table>
	<tr>`
var valStr = `        {{range .}}
        <tr>
		%s
</tr>
        {{end}}
    </table>
</body>
</html>
`

func splitOnCaps(s string) []string {
	re := regexp.MustCompile(`[A-Z][a-z]*`)
	return re.FindAllString(s, -1)
}

func getField(s string) (string, string, bool) {
	parts := strings.Fields(s)
	if len(parts) > 1 && strings.TrimSpace(parts[0]) != "}" {
		if len(parts) > 2 {
			typeStr := strings.TrimSpace(parts[1])
			if typeStr == "int64" || typeStr == "int" || typeStr == "float64" || typeStr == "float32" {
				return strings.TrimSpace(parts[0]), "$", true
			} else {
				return strings.TrimSpace(parts[0]), "", true
			}
		}
		return strings.TrimSpace(parts[0]), "", true
	}
	return "", "", false
}

func main() {
	inputFile := flag.String("input", "", "input file")
	if inputFile == nil {
		panic("input file is required")
	}
	flag.Parse()
	ProcessFile(*inputFile)
}

func ProcessFile(inputFile string) {
	var rowHeader, values []string
	println("Processing file: ", inputFile)
	fp, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	regexStruct := regexp.MustCompile(`type (\w+) struct`)

	read := bufio.NewScanner(fp)
	for read.Scan() {
		match := regexStruct.FindStringSubmatch(read.Text())
		if len(match) > 1 {
			structName := strings.Join(splitOnCaps(match[1]), " ")
			fmt.Println("Struct Name:", structName) // Extract the struct name
			tmplFile, err := os.Create(path.Join("templates", fmt.Sprintf("%s.html", structName)))
			if err != nil {
				panic(err)
			}

			fmt.Fprintf(tmplFile, header, structName, structName, structName)
			for read.Scan() {
				if fieldName, prefix, ok := getField(read.Text()); ok {
					rowHeader = append(rowHeader, fmt.Sprintf("<th>%s</th>", strings.Join(splitOnCaps(fieldName), " ")))
					values = append(values, fmt.Sprintf("<td>%s{{.%s}}</td>", prefix, fieldName))
				} else {
					fmt.Fprintf(tmplFile, "%s\n</tr>", strings.Join(rowHeader, ""))
					fmt.Fprintf(tmplFile, fmt.Sprintf(valStr, strings.Join(values, "")))
					rowHeader = nil
					values = nil
					tmplFile.Close()
					break
				}
			}
		}
	}
}
