package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

const (
	inputFile  = "./testdata/test1.md"
	resultFile = "./test1.md.html"
	goldenFile = "./testdata/test1.md.html"
)

func TestParseContent(t *testing.T) {
	input, err := ioutil.ReadFile(inputFile)
	if err != nil {
		t.Fatal(err)
	}
	result, err := parseContent(input, "")
	if err != nil {
		t.Fatal(err)
	}
	expected, err := ioutil.ReadFile(goldenFile)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(expected, result) {
		t.Logf("Expected : \n%s\n-------\n", expected)
		t.Logf("Got : \n%s\n-------\n", result)
		t.Error("Content did not match")
	}
}

func TestRun(t *testing.T) {
	var mockStdout bytes.Buffer
	if err := run(inputFile, "", &mockStdout, true); err != nil {
		t.Fatal(err)
	}
	resultFile := strings.TrimSpace(mockStdout.String())
	result, err := ioutil.ReadFile(resultFile)
	if err != nil {
		t.Fatal(err)
	}
	expected, err := ioutil.ReadFile(goldenFile)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(expected, result) {
		t.Logf("expected : \n%s\n", expected)
		t.Logf("result : \n%s\n", result)
		t.Error("Result content does not match golden file")
	}
	os.Remove(resultFile)
}
