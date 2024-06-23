package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	//fmt.Println(resp)
	//readBytes(resp)
	//copyReaderToWriterInterface(resp)

	readWithCustomWriter(resp)
}

func readWithCustomWriter(resp *http.Response) {
	logWriter := logWriter{}
	io.Copy(logWriter, resp.Body)
}

/**
* Working with the read function from Reader interface
 */
func readBytes(resp *http.Response) {
	bs := make([]byte, 99999) // Give me an empty byte slice that has 9999 byte length
	// Read the response body into the byte slice
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}

/**
* io.Copy is used to print the entire response body to the console. This is useful for debugging or displaying the response content.
* The io.Copy function takes two arguments: a Writer and a Reader. The Writer is the destination, and the Reader is the source.
 */
func copyReaderToWriterInterface(resp *http.Response) {
	io.Copy(os.Stdout, resp.Body)
}

func (logWriter) Write(byteSlice []byte) (int, error) {
	fmt.Println(string(byteSlice))
	fmt.Println("Wrote bytes of length ", len(byteSlice))
	return len(byteSlice), nil
}
