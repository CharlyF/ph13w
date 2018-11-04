package main


import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
    r.HandleFunc("/uploadImage", func(w http.ResponseWriter, r *http.Request) {
        r.ParseMultipartForm(640000000)
        file, handler, err := r.FormFile("image")
        if err != nil {
            fmt.Println(err)
            return
        }

        defer file.Close()
        fmt.Fprintf(w, "%v", handler.Header)

		ctx := context.Background()
		client, err := storage.NewClient(ctx)
		if err != nil {
			log.Fatalf("Failed to create client: %v", err)
		}

		wc := client.Bucket("ph13w-images").Object(handler.Filename).NewWriter(ctx)
		resp := fmt.Sprintf("{'uploaded': '%s'}", handler.Filename)
		if _, err = io.Copy(wc, file); err != nil {
			fmt.Println(err)
			resp = fmt.Sprintf("{'error': '%s'}", err)
		}

		if err := wc.Close(); err != nil {
			fmt.Println(err)
			resp = fmt.Sprintf("{'error': '%s'}", err)
		}

        fmt.Fprintf(w, resp)
    })

    r.HandleFunc("/createImage", func(w http.ResponseWriter, r *http.Request) {
        //vars := mux.Vars(r)

        fmt.Fprintf(w, "{'status': 'ok', 'action': 'createImage'}");
    })

    r.HandleFunc("/listImages", func(w http.ResponseWriter, r *http.Request) {
        //vars := mux.Vars(r)

        fmt.Fprintf(w, "{'status': 'ok', 'action': 'listImages'}");
    })

    http.ListenAndServe(":8080", r)
}
