package main

import "net/http"

func main() {
	data, err := http.Get("https://read.amazon.com/notebook?ref_=kcr_notebook_lib&language=en-US")
	if err != nil {
		panic(err)
	}
	var dat []byte
	cnt, err := data.Body.Read(dat)
	if err != nil {
		panic(err)
	}
	defer data.Body.Close()
	println(data.StatusCode)
	println(cnt)
	println(string(dat))

	println("hello world")
}
