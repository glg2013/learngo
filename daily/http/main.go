package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "say hello~~")
}

func main() {
	http.HandleFunc("/hi", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("不出意外的出意外了~~ ", err)
		return
	}
}
