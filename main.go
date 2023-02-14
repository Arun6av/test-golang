package main

import (
	"fmt"
	"net/http"
)

/*func main() {
	HandleFunc("/", func(w ResponseWriter, r *Request) {
		fmt.Fprintf(w, "Hello World")
	})

	err := ListenAndServe("localhost:8000/", nil)
	if err != nil {
		fmt.Println(err)
	}
}*/

func main() {
	http.HandleFunc("/", Server)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		return
	}

}
func Server(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "message: %s", r.URL.Path[1:])
}
