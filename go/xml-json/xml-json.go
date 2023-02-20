package xmltojson

import (
	"encoding/json"
	"encoding/xml"
	"io"
)

type dat struct {
	PName     string `xml:"pname,attr" json:"pname"`
	Street    string `xml:"street,attr" json:"street"`
	City      string `xml:"city,attr" json:"city"`
	StreetNum int    `xml:"num,attr" json:"num,string"`
}

func writeJson(user dat, w io.Writer) error {
	encode := json.NewEncoder(w)
	err := encode.Encode(user)
	if err != nil {
		return err
	}
	return nil
}

func writeXML(user dat, w io.Writer) error {
	// myDat, err := xml.Marshal(user)
	// if err != nil {
	// 	return err
	// }
	encode := xml.NewEncoder(w)
	err := encode.Encode(user)
	if err != nil {
		return err
	}
	return nil
}

func readJson(r io.Reader) (*dat, error) {
	user := dat{}
	decode := json.NewDecoder(r)
	err := decode.Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func readXML(r io.Reader) (*dat, error) {
	user := dat{}
	decode := xml.NewDecoder(r)
	err := decode.Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil

}
