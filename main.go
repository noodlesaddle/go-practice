package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("http get error.")
		os.Exit(1)
	}

	io.Copy((os.Stdout), resp.Body)

}
