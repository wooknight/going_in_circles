package main_test

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

const success = "\u2713"
const failed = "\u2717"

var (
	binName  = "todo"
	filename = ".todo.json"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool....")
	if runtime.GOOS == "windows" {
		binName += ".exe"
	}
	build := exec.Command("go", "build", "-o", binName)
	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool : %s : %s", binName, err)
		os.Exit(1)
	}
	if os.Getenv("TODO_FILE_NAME") != "" {
		filename = os.Getenv("TODO_FILE_NAME")
	}
	fmt.Println("Running tests ...")
	result := m.Run()
	fmt.Println("cleaning up ...")
	os.Remove(binName)
	os.Remove(filename)
	os.Exit(result)
}

func TestToDoCLI(t *testing.T) {
	t.Log("We are testing the CLI")
	{
		task := "Task 1"
		task2 := "test task number 2"
		dir, err := os.Getwd()
		if err != nil {
			t.Fatal(err)
		}
		cmdPath := filepath.Join(dir, binName)
		t.Log("We are adding")
		{
			t.Run("AddNewTaskFromArguments", func(t *testing.T) {
				cmd := exec.Command(cmdPath, "-add", task)
				if err := cmd.Run(); err != nil {
					t.Fatal(err)
				}
			})

			t.Run("AddNewTaskFromSTDIN", func(t *testing.T) {
				cmd := exec.Command(cmdPath, "-add")
				cmdStdIn, err := cmd.StdinPipe()
				if err != nil {
					t.Fatalf("%s %v", failed, err)
				}
				io.WriteString(cmdStdIn, task2)
				cmdStdIn.Close()
				if err := cmd.Run(); err != nil {
					t.Fatal(err)
				}
			})
		}
		t.Log("We are testing the default case")
		{
			t.Run("ListTasks", func(t *testing.T) {
				cmd := exec.Command(cmdPath, "-list")
				out, err := cmd.CombinedOutput()
				if err != nil {
					t.Fatal(err)
				}
				expected := fmt.Sprintf("  1: %s\n  2: %s  3: %s\n  4: %s\n", task, task2, task, task2)
				if expected != string(out) {
					t.Errorf("%s Expected = %q , got %q instead\n", failed, expected, string(out))
				}
			})
		}
		t.Log("We are testing delete")
		{
			t.Run("DeleteTasks", func(t *testing.T) {
				{
					cmd := exec.Command(cmdPath, "-del", "1")
					_, err := cmd.CombinedOutput()
					if err != nil {
						t.Fatal(err)
					}
				}
				{
					cmd := exec.Command(cmdPath, "-list")
					out, err := cmd.CombinedOutput()
					if err != nil {
						t.Fatal(err)
					}
					expected := fmt.Sprintf("  1: %s\n", task)
					if expected != string(out) {
						t.Errorf("%s Expected = %q , got %q instead\n", failed, expected, string(out))
					}
				}
			})
		}
		t.Log("We are testing the verbose case")
		{
			t.Run("VerboseTasks", func(t *testing.T) {
				cmd := exec.Command(cmdPath, "-verbose")
				out, err := cmd.CombinedOutput()
				if err != nil {
					t.Fatal(err)
				}
				fmt.Println(string(out))
				// expected := fmt.Sprintf("  1: %s\n", task)
				// if expected != string(out) {
				// 	t.Errorf("%s Expected = %q , got %q instead\n", failed, expected, string(out))
				// }
			})
		}
		t.Log("We are testing the unfinished case")
		{
			t.Run("UnfinishedTasks", func(t *testing.T) {
				{
					cmd := exec.Command(cmdPath, "-unfinished")
					out, err := cmd.CombinedOutput()
					if err != nil {
						t.Fatal(err)
					}
					expected := fmt.Sprintf("  1: %s\n", task)
					if expected != string(out) {
						t.Errorf("%s Expected = %q , got %q instead\n", failed, expected, string(out))
					}
				}
				{
					cmd := exec.Command(cmdPath, "-complete", "0")
					_, err := cmd.CombinedOutput()
					if err != nil {
						t.Fatal(err)
					}

					cmd = exec.Command(cmdPath, "-unfinished")
					out, err := cmd.CombinedOutput()
					if err != nil {
						t.Fatal(err)
					}
					expected := "\n"
					if expected != string(out) {
						t.Errorf("%s Expected = %q , got %q instead\n", failed, expected, string(out))
					}
				}

			})
		}

	}
}
