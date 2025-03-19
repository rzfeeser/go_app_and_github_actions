// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"os"
	)

func main() {
	fmt.Println("Hello, Russell Zachary Feeser of https://rzfeeser.com and https://iris7.com")
        fmt.Println("Attempting to print the ENV secret from GitHub repo; value of MYTESTENV: ", os.Getenv("MYTESTENV")) 
}
