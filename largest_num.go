package main
 
import "fmt"
 
func main() {
    fmt.Println("Enter 3 numbers :")
    var first,second,third int
    fmt.Scanln(&first)
    fmt.Scanln(&second)
    fmt.Scanln(&third)

    if (first>=second && first>=third) {
        fmt.Println(first, "is Largest among three numbers.") 
    }
    if (second>=first && second>=third){
        fmt.Println(second, "is Largest among three numbers.")
    }
    if (third>=first && third>=second) {
        fmt.Println(third, "is Largest among three numbers")
    }
}