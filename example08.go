/* Alta3 Research | RZFeeser
   Playing around with custom errors */

package main

import (
	"errors" // we can deep dive on this later
	"fmt"
)

func marvel(movie string) (string, error) {
	if movie == "Squid-man" || movie == "Fish-man" {
		return movie, errors.New("Ocean superhero fauna are not suitable movie names")
	}
	return movie + " is green lit!", nil
}

func main(){

	// create a var for spider-man
	movie := "Spider-man"

	movieResults, err := marvel(movie)
        fmt.Println(movie)

	if err != nil {
		panic(err)
	}

	fmt.Println(movieResults)
}
