/* RZFeeser
   strings.SplitN() */

package main

import (
	"fmt"
	"strings"
)

func main() {

	// String x a is comma serparated string
	// The separater used is ","
	// This will split the string into 6 parts
	x := strings.SplitN("a,b,c,d,e,f", ",",6)
	fmt.Println(x)

	// I want 5 substrings (leave 1 comma at the end)
	y := strings.SplitN("a,b,c,d,e,f", ",",5)
        fmt.Println(y)

	// I want 4 substrings (leave 2 commas at the end)
	z := strings.SplitN("a,b,c,d,e,f", ",", 0)
        fmt.Println(z)
}
