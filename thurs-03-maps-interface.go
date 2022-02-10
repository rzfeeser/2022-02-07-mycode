/* Alta3 Research | RZFeeser
   Maps - Using empty interfaces as a "wildcard" to create
   composite types */

package main

import "fmt"


func main() {

	
    futurama := make(map[string]interface{})

    futurama["name"] = "Turanga Leela"
    futurama["age"] = 30
    futurama["height"] = 182.5

    fmt.Printf("%+v\n", futurama)
}
