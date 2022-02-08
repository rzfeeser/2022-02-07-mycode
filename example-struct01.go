package main

import "fmt"

type position struct {
	lat string
	lon string
}

func main() {
	issloc := position{"39.8218", "-88.6526"}  // create and instance of iss lat / lon

	fmt.Println(issloc)

}
