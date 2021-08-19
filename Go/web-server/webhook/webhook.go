package webhook

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func processData(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Nothing to see here\n")
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "Encountered an error %+v", err)
			return
		}
		fmt.Fprintf(w, "Post from site : %+v\n", r.PostForm)
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
