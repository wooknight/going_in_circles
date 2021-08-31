package rest

import "net/http"

func GetRecords(w http.ResponseWriter, r *http.Request) {
	[]records 
}

func AddRecord(w http.ResponseWriter, r *http.Request) {
	[]records 
}

func DeleteRecord(w http.ResponseWriter, r *http.Request) {
}
func UpdateRecord(w http.ResponseWriter, r *http.Request) {

}


func InitRest(){
	rest = new (http.Handler)
	rest.HandleFunc()
	http.ListenAndServeTLS()
}