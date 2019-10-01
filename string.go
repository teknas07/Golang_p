package main

import "fmt"

func main() {
   var greeting =  "Your ready to Go"
   
   fmt.Printf("normal string: ")
   fmt.Printf("%s", greeting)
   fmt.Printf("\n")
   fmt.Printf("hex bytes: ")
   
   for i := 0; i < len(greeting); i++ {
       fmt.Printf("%x ", greeting[i])
   }
   fmt.Printf("\n")
  
}