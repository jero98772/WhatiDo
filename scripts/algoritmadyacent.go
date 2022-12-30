package main
import "fmt"
func whaticanmake(graph map[string][]string,objets []string) string {
	ocurrences:=make(map[string]int)
	max:=0
	item:=""
	for _,i:=range objets{
		for _,j:= range graph[i]{
			_,ok:=ocurrences[j]
			if ok{
				ocurrences[j]=ocurrences[j]+1
			}else{
				ocurrences[j]=1
			}
		}
	}
	for i,_:=range ocurrences{
		if max<ocurrences[i]{
			max=ocurrences[i]			
			item=i
		}
	}
	//fmt.Println(ocurrences)
	return item
}
func main(){
	graph:=make(map[string][]string)
	graph["tierra"]=[]string{"huerta","casa"}
	graph["madera"]=[]string{"huerta","silla","casa"}
	graph["huerta"]=[]string{}
	graph["ladrillos"]=[]string{"casa","silla"}
	objets:=[]string{"tierra","madera","ladrillos"}//objets i have
	fmt.Println(whaticanmake(graph,objets))
}