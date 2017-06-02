package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type metadata struct {
	Size int64 `json:"size"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/get-file-size", getFileSize)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":"+port, nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./static/index.html")
}

func getFileSize(res http.ResponseWriter, req *http.Request) {
	file, _, err := req.FormFile("selected-file")
	check(err) //Need to check why these errors are coming up.
	defer file.Close()

	var buff bytes.Buffer
	fileSize, err := buff.ReadFrom(file)
	check(err) //Need to check why these errors are coming up.

	var data metadata

	data.Size = fileSize

	js, err := json.Marshal(data)
	check(err) //Need to check why these errors are coming up.
	res.Header().Set("Content-Type", "application/json")
	res.Write(js)
}

func check(err error) {
	fmt.Printf("Error: %v", err)
}
