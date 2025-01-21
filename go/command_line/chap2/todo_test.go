package main_test

import (
	"log"
	"os/exec"
	"testing"
)

var bin_name = "todo"
var test_file = "test.txt"

func TestApplication(t *testing.T) {
	//build app
	build := exec.Command("go", "build", "-o", bin_name)
	if err := build.Run(); err != nil {
		log.Fatal(err)
	}
	//run app

	run := exec.Command("./" + bin_name)
	//get the combined stdout and stderr
	stdout, err := run.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	//check the output
	log.Println(string(stdout))
}
