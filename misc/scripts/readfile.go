package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strconv"
)
type unit struct{
	amount int
	amountUnits string
}
type material struct{
	units unit 
	conexions []string

}
func main() {

     content, err := ioutil.ReadFile("graph.txt")
     graph:=make(map[string]material)
     if err != nil {
          log.Fatal(err)
     }
     token:=""
     key:=""
     varunit:=""
     varamount:=""
     values:=make([]string,0)
    	for _,i:=range content{
			i:=string(i)
			if i=="{"{
				values=make([]string,0)
				//graph[token]=[]string{}
				key=token
				token=""
			}else if i==","{
				values=append(values,token)
				token=""
			}else if i=="\n"{
				//graph=append(graph[key],token)
				values=append(values,token)
				token=""
				amount,_:=strconv.Atoi(varamount)
				graph[key]=material{unit{amount,varunit},values}
				key=""
				varunit=""
				varamount=""
			}else if i=="."{
				varamount=token
				token=""
			}else if i=="}"{
				varunit=token
				token=""
			}else{
				token=token+i
			}
			fmt.Println(string(i))
    	}
    fmt.Println(graph)
}