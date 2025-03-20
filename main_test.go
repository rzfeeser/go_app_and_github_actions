/* test example | RZFeeser
   When run, go test will look at all Go files in the current directory for two keywords:
   
   Files matching the pattern, *_test.go
   Any function matching the pattern [Tt]est* Note: Best practice would indicate you should use a capital T */

package main

import (
    "testing"   // used for testing
    "regexp"    // regular expression library
)

func Test_site_responder(t *testing.T) {

    result := site_responder()
    
    if result != "https://iris7.com" {
        t.Fatalf("something is wrong with the function site_responder()")
    }
}
