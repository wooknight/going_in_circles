package xmltojson

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var user = dat{
	PName:     "Choot",
	Street:    "puss ave",
	City:      "Dick, VA",
	StreetNum: 666,
}

func TestWriteJson(t *testing.T) {
	outJSON, _ := ioutil.TempFile(os.TempDir(), "out.json")
	writeJson(user, outJSON)
	log.Println(outJSON.Name())
	outJSON.Seek(0, 0)
	tmpUser, err := readJson(outJSON)
	if err != nil {
		t.Fatalf("Could not read file %s : Error : %v", outJSON.Name(), err)
	}
	if *tmpUser != user {
		t.Fatalf("Data mismatch : Expected %v : Got %v", user, *tmpUser)
	}
}

func TestWriteXML(t *testing.T) {
	outXML, _ := ioutil.TempFile(os.TempDir(), "out.xml")
	writeXML(user, outXML)
	log.Println(outXML.Name())
	outXML.Seek(0, 0)
	tmpUser, err := readJson(outXML)
	if err != nil {
		t.Fatalf("Could not read file %s : Error : %v", outXML.Name(), err)
	}
	if *tmpUser != user {
		t.Fatalf("Data mismatch : Expected %v : Got %v", user, *tmpUser)
	}
}
