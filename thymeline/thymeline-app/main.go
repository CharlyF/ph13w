package main

import (
	"cloud.google.com/go/storage"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/api/iterator"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type imgClient struct {
	cl *storage.Client
}
const (
	mainBucket = "ph13w-images"
)

type response struct {
	fileName string `json:"filename,omitempty"`
	url string		`json:"url,omitempty"`
	timestamp int64	`json:"ts,omitempty"`
}

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
		objectName := fmt.Sprintf("%s-%s-%d", whoUploaded, filenameNoUnderscores, int32(time.Now().Unix()))

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

    r.HandleFunc("/listImages", listImages)

    http.ListenAndServe(":8080", r)
}

func newClient(ctx context.Context) *imgClient {
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return &imgClient{
		client,
	}
}

func listImages(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	client := newClient(ctx)
	imgList := []response{}
	it := client.cl.Bucket(mainBucket).Objects(ctx, nil)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Fprintf(w, fmt.Sprintf(`{'status': 'error', log: '%s'}`, err.Error()))
			return
		}
		fmt.Fprintln(w, attrs.Name)
		img := response{
			fileName: attrs.Name,
			url: attrs.MediaLink,
			timestamp: attrs.Generation,
		}
		imgList = append(imgList, img)
	}
	jList, err := json.Marshal(imgList)
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf(`{'status': 'error', log: '%s'}`, err.Error()))
		return
	}
	fmt.Fprintf(w, string(jList))
	fmt.Fprintf(w, "{'status': 'ok', 'action': 'listImages'}");
}