package main

import (
	"os"
	"testing"

	"github.com/go/command_line/chap3/mdp/markdown"
)

const (
	markdownFile   = "./testdata/test.md"
	goldenHtmlFile = "./testdata/test.md.html"
	testHtmlFile   = "./testdata/test.md.html.tmp"
)

func TestParseContent(t *testing.T) {
	md := &markdown.MarkDown{}
	buf, err := md.Open(markdownFile)
	if err != nil {
		t.Fatal(err)
	}
	md.ParseContent(buf)
	expected, err := os.ReadFile(goldenHtmlFile)
	if err != nil {
		t.Fatal(err)
	}
	if string(md.Text) != string(expected) {
		t.Errorf("expected \n%s\n, got \n%s\n", expected, md.Text)
	}
}
