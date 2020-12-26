## I copy this code from website that I cannot remember. Sorry !!
package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    fmt.Println("Starting !!")
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":4000", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s :: %d !\n", r.URL.Path[1:],os.Getpid())
}
