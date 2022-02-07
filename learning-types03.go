/* Alta3 Research | Author: RZFeeser
  Types - transformations   */

package main

import (
    "fmt"
    "math"
)

func main() {
    
    // x and y are integers
    var x, y int = 3, 4
    
                            // transform x and y
    var fun float64 = math.Sqrt(float64(x*x + y*y))
    
    // tranform into unsigned integer
    var z uint = uint(fun)
    
    // display f
    fmt.Println(fun)

    // display f to 1 decimal
    fmt.Printf("%.1f\n", fun)

    // display f to 2 decimals
    fmt.Printf("%.2f\n", fun)

    // display x, y and z
    fmt.Println(x, y, z)
}

