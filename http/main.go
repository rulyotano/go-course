package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}

func main () {
	response, err := http.Get("https://google.com")

	if err != nil {
		fmt.Println("There was an error making the request:", err)
		os.Exit(1)
	}

	if response.StatusCode != http.StatusOK {
		fmt.Println("The request was not successful. Status code:", response.StatusCode)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, response.Body)
	// buffer := make([]byte, 256)
	// content := []byte{}
	// nreaded, _ := response.Body.Read(buffer)

	// for ; nreaded > 0; {
	// 	content = append(content, buffer[:nreaded]...)
	//   nreaded, _ = response.Body.Read(buffer)
	// }
	
	// fmt.Println(string(content))
}