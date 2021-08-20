package webhook

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func ProcessData(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Fatalf("Encountered an error : %v\n", err)
	}
	log.Printf("Request data %s", string(dump))
	switch r.Method {
	case "GET":
		// curl http://127.0.0.1:8000/orders
		fmt.Fprintf(w, "Nothing to see here\n")
	case "POST":
		//curl  -v -i -X  POST -F "file=@README.md" http://0.0.0.0:8000/orders
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "Encountered an error %+v", err)
			return
		}
		f, h, err := r.FormFile("file")

		if err != nil {
			fmt.Fprintf(w, "Could not decode body data : %+v\n", err)
			return
		}
		defer f.Close()
		fName := "/tmp/" + h.Filename
		out, err := os.Create(fName)
		if err != nil {
			fmt.Fprintf(w, "Could not create file : %+v\n", err)
			return
		}
		defer out.Close()
		io.Copy(out, f)
		fmt.Fprintf(w, "Upload complete")
	}
	//save info to database
}
