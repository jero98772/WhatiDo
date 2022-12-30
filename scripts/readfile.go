package main

import (
    "fmt"
    "io/ioutil"
    "log"
)

func main() {

     content, err := ioutil.ReadFile("graph.txt")
     graph:=make(map[string][]string)
     if err != nil {
          log.Fatal(err)
     }
     token:=""
     key:=""
     values:=make([]string,0)
    	for _,i:=range content{
		i:=string(i)
		if i=="-"{
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
			graph[key]=values
			key=""
		}else{
		token=token+i
		}
		fmt.Println(string(i))
    
	}
    fmt.Println(graph)
}