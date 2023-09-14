package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type config struct {
	ext  string
	size int64
	list bool
	del  bool
	dry  bool
	wLog io.Writer
}

var f = os.Stdout
var err error

func main() {
	root := flag.String("root", ".", "Root directory to start")
	list := flag.Bool("list", false, "List files only")
	ext := flag.String("ext", "", "File extension to filter out")
	size := flag.Int64("size", 0, "minimum file size")
	del := flag.Bool("del", false, "Delete files")
	dry := flag.Bool("dry-run", false, "Dry Run")
	logFile := flag.String("log", ".", "Log deletes to file")
	if *logFile != "" {
		f, err = os.Open(*logFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer f.Close()
	}
	flag.Parse()
	c := config{
		ext:  *ext,
		size: *size,
		list: *list,
		del:  *del,
		dry:  *dry,
		wLog: f,
	}
	if err := run(*root, os.Stdout, c); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(root string, out io.Writer, cfg config) error {
	delLogger := log.New(cfg.wLog, "DELETED FILE:", log.LstdFlags)
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filterOut(path, cfg.ext, cfg.size, info) {
			return nil
		}
		if cfg.list {
			return listFile(path, out)
		}
		if cfg.del {
			return delFile(path, cfg.dry, out, delLogger)
		}

		return listFile(path, out)
	})
}
