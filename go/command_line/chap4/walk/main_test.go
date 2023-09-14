package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var (
	buffer    bytes.Buffer
	logBuffer bytes.Buffer
)

func TestRun(t *testing.T) {
	testCases := []struct {
		name        string
		cfg         config
		extNoDelete string
		nDelete     int
		nNoDelete   int
		expected    string
	}{
		{name: "DeleteExtensionNoMatch",
			cfg:         config{ext: ".log", del: true},
			extNoDelete: ".gz", nDelete: 0, nNoDelete: 10,
			expected: "",
		},
		{name: "DeleteExtensionMatch",
			cfg:         config{ext: ".log", del: true},
			extNoDelete: "", nDelete: 10, nNoDelete: 0,
			expected: "",
		},
		{name: "DeleteExtensionMixed",
			cfg:         config{ext: ".log", del: true},
			extNoDelete: ".gz", nDelete: 5, nNoDelete: 5,
			expected: "",
		},
	}

	for _, tc := range testCases {
		tc.cfg.wLog = &logBuffer
		t.Run(tc.name, func(t *testing.T) {
			var buffer bytes.Buffer
			tempDir, cleanup := createTempDir(t, map[string]int{
				tc.cfg.ext:     tc.nDelete,
				tc.extNoDelete: tc.nNoDelete,
			})

			err := run(tempDir, &buffer, tc.cfg)
			if err != nil {
				t.Fatal(err)
			}
			cleanup()
			res := buffer.String()
			if tc.expected != res {
				t.Errorf("Expected %q, got %q instead\n", tc.expected, res)
			}
			expLogLines := tc.nDelete + 1
			lines := bytes.Split(logBuffer.Bytes(), []byte("\n"))
			if len(lines) != expLogLines {
				for idx, val := range lines {
					t.Logf("\n\nlog buffer data %d\n\n%+v\n\n", idx, string(val))
				}

				t.Errorf("Expected %d log lines , got %d instead\n", expLogLines, len(lines))
			}
			filesLeft, err := ioutil.ReadDir(tempDir)
			if err != nil {
				t.Error(err)
			}
			if len(filesLeft) != tc.nNoDelete {
				t.Errorf("Expected %d files left ,got %d instead ", tc.nNoDelete, len(filesLeft))
			}
		})
	}
}

func createTempDir(t *testing.T, files map[string]int) (dirName string, cleanup func()) {
	t.Helper()
	tempdir, err := ioutil.TempDir("", "walktest")
	if err != nil {
		t.Fatal(err)
	}
	for k, n := range files {
		tempdir, _ = ioutil.TempDir("", "walktest")
		for j := 1; j <= n; j++ {
			fname := fmt.Sprintf("file%d%s", j, k)
			fpath := filepath.Join(tempdir, fname)
			if err := ioutil.WriteFile(fpath, []byte(fmt.Sprintf("dummy %s - %d", k, j)), 0644); err != nil {
				t.Fatal(err)
			}
		}

	}
	return tempdir, func() { os.RemoveAll(tempdir) }
}

func TestRunDelExtension(t *testing.T) {
	testCases := []struct {
		name        string
		cfg         config
		extNoDelete string
		nDelete     int
		nNoDelete   int
		expected    string
	}{
		{name: "DeleteExtensionNoMatch",
			cfg:         config{ext: ".log", del: true},
			extNoDelete: ".gz", nDelete: 0, nNoDelete: 10,
			expected: ""},
		{
			name:        "DeleteExtensionMixed",
			cfg:         config{ext: ".log", del: true},
			extNoDelete: ".gz", nDelete: 5, nNoDelete: 5,
		},
		{name: "DeleteExtensionMatch",
			cfg:         config{ext: ".log", del: true},
			extNoDelete: "", nDelete: 10, nNoDelete: 0,
		},
	}

	for _, tc := range testCases {
		tc.cfg.wLog = &logBuffer
		t.Run(tc.name, func(t *testing.T) {
			var buffer bytes.Buffer

			tempDir, cleanup := createTempDir(t, map[string]int{
				tc.cfg.ext:     tc.nDelete,
				tc.extNoDelete: tc.nNoDelete,
			})
			defer cleanup()
			if err := run(tempDir, &buffer, tc.cfg); err != nil {
				t.Fatal(err)
			}
			res := buffer.String()
			if tc.expected != res {
				t.Errorf("Expected %q , got %q instead \n", tc.expected, res)
			}
			filesLeft, err := ioutil.ReadDir(tempDir)
			if err != nil {
				t.Error(err)
			}
			if len(filesLeft) != tc.nNoDelete {
				t.Errorf("Expected %d files left, got %d instead\n", tc.nNoDelete, len(filesLeft))
			}
		})
	}
}
