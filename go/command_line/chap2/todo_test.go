package main_test

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

var bin_name = "todo"
var test_file = "test.txt"

func TestMain(m *testing.M) {
	//build app
	build := exec.Command("go", "build", "-o", bin_name)
	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "cannot build the app: %s . %v", bin_name, err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stdout, "Build successful.Running tests...")
	//run tests
	exitVal := m.Run()
	//clean up
	os.Remove(bin_name)
	os.Remove(test_file)
	os.Exit(exitVal)
}

func TestTodo(t *testing.T) {
	task := "test task number 1"
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	//run the app
	cmdPath := filepath.Join(dir, bin_name)
	t.Run("AddNewTask", func(t *testing.T) {
		cmd := exec.Command(cmdPath)
		stdin, err := cmd.StdinPipe()
		if err != nil {
			log.Fatal(err)
		}

		go func() {
			defer stdin.Close()
			io.WriteString(stdin, "add "+task+"\nquit\n")
		}()

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cmdPath)
		cmdStdin, err := cmd.StdinPipe()
		if err != nil {
			t.Fatal(err)
		}
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		go func() {
			defer cmdStdin.Close()
			io.WriteString(cmdStdin, "list\n")
			io.WriteString(cmdStdin, "quit\n")
		}()
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(string(out), task) {
			t.Errorf("expected %q to contain %q", string(out), task)
		}
	})

}
