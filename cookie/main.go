package main

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: "testcookiename", Value: "a,b,c", Path: "foo.com", MaxAge: 86400}
	http.SetCookie(w, &cookie)
	fmt.Fprintln(w, "hello world")
}

func main() {
	http.HandleFunc("/", IndexHandler)
	err := http.ListenAndServe(":19000", nil)
	if err != nil {
		fmt.Printf("ListenAndServe: ", err)
	}
}
