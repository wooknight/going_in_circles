package webhook

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/wooknight/going_in_circles/go/web-server/mytemplate"
)

type SecretTokenHandler struct {
	Next      http.Handler
	SecretJWT string
}

func (h *SecretTokenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("secretJWT") == h.SecretJWT {
		h.Next.ServeHTTP(w, r)
	} else {
		http.NotFound(w, r) //unauthenticated flow
	}
}

type UptimeHandler struct {
	Started time.Time
}

func NewUptimeHandler() *UptimeHandler {
	return &UptimeHandler{Started: time.Now()}
}

func (h *UptimeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Fatalf("Encountered an error : %v\n", err)
	}
	log.Printf("Request data %s", string(dump))
	switch r.Method {
	case "GET":
		// curl http://127.0.0.1:8000/orders
		mytemplate.PrintDB(w, fmt.Sprintf("Current uptime: %s", time.Since(h.Started)))

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
		safeDir := "/tmp/"
		absPath, err := filepath.Abs(filepath.Join(safeDir, h.Filename))
		if err != nil || !strings.HasPrefix(absPath, safeDir) {
			http.Error(w, "Invalid file name", http.StatusBadRequest)
			return
		}
		out, err := os.Create(absPath)
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

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	params := r.URL.Query()

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category is %s\nID is %s\nparams num is %s\nDescription is %s\n", vars["category"], vars["id"], params["num"], params["desc"])

}
