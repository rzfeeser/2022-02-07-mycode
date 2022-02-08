/* RZFeeser | Alta3 Research
   Introduction to working with methods.

   The basic pattern of a method is as follows:

   func(receiver_name Type) method_name(parameter_list)(return_type){
      // Code
   }
   

*/

package main

import "fmt"

// define the type astro
type astro struct {
    name     string
    age      int
    mission  string
    isneeded bool
}

// define a method called "honorific" that does not expect any additional values to be passed
func (a astro) honorific() string {
        return "Space Hero " + a.name
}

// define a method called "customhonorific" that expects a string to be passed to it
func (a astro) customhonorific(honor string) string {
	return honor + a.name
}

// define a method called "situp". The * is a "dereference" and is saying, "we don't want to
// copy the VALUES that were just passed, we want to actually point to the memory cell location
// of the instance passed". This is required as we want to MUTATE the data
func (a *astro) suitup() {
        a.isneeded = true
}

func main() {

    // create several structures
    ast1 := astro{"Megan McArthur", 35, "ISS", false}
    ast2 := astro{"Barry Wilmore", 41, "Boeing Crew Flight Test (CFT)", true}
    ast3 := astro{"Raja Chari", 39, "SpaceX Crew-3", true}

    // display the structures created
    fmt.Println(ast1)
    fmt.Println(ast2)
    fmt.Println(ast3)

    // call our honorific function
    fmt.Println(ast1.honorific())
    // call the customhonorific function
    fmt.Println(ast1.customhonorific("Awesome Space Hero "))
    
    // this says, "point to the memory address of ast1"
    // the reasons for doing this would be to avoid making a NEW copy of the data
    // or because we are interested in MODIFYING the received structure (ast1) via the pointer
    ast1_pointer := &ast1
    fmt.Println(ast1_pointer.customhonorific("Awesome Space Hero ")) // methods move with pointers

    // try to MUTATE our struct
    fmt.Println("False to True")
    fmt.Println(ast1)  // the ast1 struct has "isneeded: false"
    ast1.suitup()      // modifies the ast1 struct to "isneeded: true"
    fmt.Println(ast1)  // proof that the ast1 struct was modified


    // slice named people made up of astro
    p := []astro{ast1, ast2, ast3}

    // display the slice
    fmt.Println(p)

    // select data from the struct
    fmt.Println(p[2].mission)

}

