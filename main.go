package main
import "fmt"
import "os"
import S "strings"
func main() {
  fmt.Println("Hello, World!" + S.Join(os.Args[1:], " "))
}