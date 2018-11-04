package main


import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
    r.HandleFunc("/uploadImage", func(w http.ResponseWriter, r *http.Request) {
        r.ParseMultipartForm(640000000)
        whoUploaded := r.FormValue("who");
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

		filenameNoUnderscores := strings.Replace(handler.Filename, "-", "_", -1)
		objectName := fmt.Sprintf("%s-%s-%s", whoUploaded, filenameNoUnderscores, int32(time.Now().Unix()))

		wc := client.Bucket("ph13w-images").Object(objectName).NewWriter(ctx)
		resp := fmt.Sprintf("{'uploaded': '%s'}", objectName)
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
