package main


import (
"fmt"
"net/http"

"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/uploadImage", func(w http.ResponseWriter, r *http.Request) {
		//vars := mux.Vars(r)

		fmt.Fprintf(w, "{'status': 'ok', 'action': 'uploadImage'}");
	})

	r.HandleFunc("/createImage", func(w http.ResponseWriter, r *http.Request) {
		//vars := mux.Vars(r)

		fmt.Fprintf(w, "{'status': 'ok', 'action': 'createImage'}");
	})

	r.HandleFunc("/listImages", func(w http.ResponseWriter, r *http.Request) {
		//vars := mux.Vars(r)

		fmt.Fprintf(w, "{'status': 'ok', 'action': 'listImages'}");
	})

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Errorf(err.Error());
	}
}
