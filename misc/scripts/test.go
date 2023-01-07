package main
import ("fmt"
	"io/ioutil")
func main(){
	content, _ := ioutil.ReadFile("data/graph.txt")
	fmt.Println(string(content))
}