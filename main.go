package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello from home!!")
	})
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil)) // 프로그램이 Exit(1) : {error 로 인해 종료되는 경우} 로그를 출력해줌

}
