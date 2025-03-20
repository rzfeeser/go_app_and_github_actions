
package main

import (
	"fmt"
	"os"
	)

func site_responder() string {
	return "https://iris7.com"

func main() {
	fmt.Println("Hello, Russell Zachary Feeser of https://rzfeeser.com and", site_responder())
        fmt.Println("Attempting to print the ENV secret from GitHub repo; value of MYTESTENV: ", os.Getenv("MYTESTENV")) 
}
