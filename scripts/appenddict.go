package main
import "fmt"
func main(){
	graph:=make(map[string][]string)
	graph["1"]=[]string{""}
	graph["2"]=[]string{"hola"}
	fmt.Println(graph["1"])
	graph=append(graph["1"],"holas")
	//fmt.Println(graph)
}