package main


import (
"fmt"
"net/http"

"github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/uploadImage", func(w http.ResponseWriter, r *http.Request) {
        r.ParseMultipartForm(32 << 20)
        file, handler, err := r.FormFile("image")
        if err != nil {
            fmt.Println(err)
            return
        }
        defer file.Close()
        fmt.Fprintf(w, "%v", handler.Header)
        f, err := os.OpenFile("./"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
        if err != nil {
            fmt.Println(err)
            return
        }
        defer f.Close()
        io.Copy(f, file)

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
