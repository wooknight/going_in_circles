package markdown

import (
	"bytes"
	"fmt"
	"os"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const (
	header = `<!DOCTYPE html>
	<html>
	<head>
	<meta http-equiv="content-type" content="text/html"; charset="utf-8">
	<title>Markdown Preview Tool</title>
	</head>
	<body>
	<h1>Markdown</h1>
	`

	footer = `</body></html>`
)

type MarkDown struct {
	Text []byte
}

func (m *MarkDown) Open(file string) ([]byte, error) {

	input, err := os.ReadFile(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}
	return input, nil
}

func (m *MarkDown) RenderHTML(outfile string) error {
	return os.WriteFile(outfile, m.Text, 0644)
}

func (m *MarkDown) ParseContent(r []byte) {
	//parse the markdown
	output := blackfriday.Run(r)
	//sanitize the output
	body := bluemonday.UGCPolicy().SanitizeBytes(output)
	var buffer bytes.Buffer
	buffer.WriteString(header)
	buffer.Write(body)
	buffer.WriteString(footer)
	m.Text = buffer.Bytes()
}
