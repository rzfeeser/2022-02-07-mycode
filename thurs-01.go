/* Alta3 Research | RZFeeser
   Interfaces - An interface type in Go is a bit like a definiton. It defines and describes teh exact methods taht some other type must have. 

   // Example taken from fmt.Stringer interface
   // for implimenting "toString" or functionality
   type Stringer interface {
    String() string
   }

Something "impliments the interface" when it has the exact signature String() string. */

package main

import (
    "fmt"
    "strconv"
    "log"    // https://pkg.go.dev/log - makes standard logging easier- writes to stderr
)

// Create a structure where we createa "Book" type
// later we will see that this satisfies the fmt.Stringer interface
type Book struct {
    Title  string
    Author string
}

// Create a function "String() string" with a receiver for Book
// we are now done, no additional code or switches to flip
func (b Book) String() string {
    return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
}

// This is not a structure, it is an integer
// even this can still satisfy the fmt.Stringer interface
type Count int

func (c Count) String() string {
    return strconv.Itoa(int(c))
}

// Declare a WriteLog() function which takes any object that satisfies
// the fmt.Stringer interface as a parameter.
func WriteLog(s fmt.Stringer) {
    log.Println(s.String())
}

func main() {
    // Initialize a Count object and pass it to WriteLog().
    book := Book{"1984", "George Orwell"}
    WriteLog(book)

    // Initialize a Count object and pass it to WriteLog().
    count := Count(5)
    WriteLog(count)
}
