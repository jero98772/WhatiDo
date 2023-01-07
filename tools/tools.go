//whatIDo - by jero98772
package tools
import "fmt"
type choice struct{
	thechoice string
}
func Str2Graph(text string) map[string][]string{
	fmt.Println("tool")
	fmt.Println(text)
	graph:=make(map[string][]string)
	token:=""
	key:=""
	values:=make([]string,0)
	for _,i:=range text{
		i:=string(i)
		//fmt.Println(i)
		if i=="-"{
			values=make([]string,0)
			key=token
			token=""
		}else if i==","{
			values=append(values,token)
			token=""
		}else if i=="\n"{
			values=append(values,token)
			token=""
			graph[key]=values
			key=""
		}else{
			token=token+i
		}
		//fmt.Println(string(i))
	}
	return graph
}
func Whaticanmake(graph map[string][]string,objets []string) string {
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
	return item
}
func Str2arrstr(text string ,sep string)[]string{
	token:=""
	words:=make([]string,0)
	fmt.Println(text)
	for _,i:= range text{
		i:=string(i)
		fmt.Println(i,token)
		if i==sep{
			words=append(words,token)			
			token=""
		}else{
			token=token+i
		}
	}
	words=append(words,token)
	fmt.Println(words)
	return words
}
