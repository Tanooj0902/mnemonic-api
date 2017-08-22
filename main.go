package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	InitLog(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	http.HandleFunc("/private_key", requestLogger(headerSetter(postRequestHandler(requestParser(privateKeyAction)))))

	Info.Println("Server Started and Listening on PORT: 8080 ")

	http.ListenAndServe(":8080", nil)
}
