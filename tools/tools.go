//whatIDo - by jero98772
package tools
import ("fmt"
		"strconv"
		"strings"
	)
type choice struct{
	thechoice string
}
type unit struct{
	amount int
	amountUnits string
}
type material struct{
	units unit 
	conexions []string

}
func Str2GraphPond(text string) map[string]material{
	//fmt.Println("tool")
	//fmt.Println(text)
	text=strings.ReplaceAll(text, " ", "")
	graph:=make(map[string]material)
	token:=""
	key:=""
	varunit:=""
	varamount:=""
		fmt.Println(graph)
	values:=make([]string,0)
	for _,i:=range text{
		i:=string(i)
		//fmt.Println(i)
		if i=="{"{
			values=make([]string,0)
			key=token
			token=""
		}else if i==","{
			values=append(values,token)
			token=""
		}else if i=="\n"{
			values=append(values,token)
			token=""
			//print names and more
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
	}
	fmt.Println(graph)
	return graph
}
func Str2Graph(text string) map[string][]string{
	//fmt.Println("tool")
	//fmt.Println(text)
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
func Formatmymaterials(text string) map[string]unit{
	text=strings.ReplaceAll(text, " ", "")
	mymaterials:=make(map[string]unit)
	token:=""
	key:=""
	varunit:=""
	varamount:=""
	for _,vi:=range text{
		vi:=string(vi)
		if "{"==vi{
			key=token
			token=""
		}else if "."==vi{
			varamount=token
			token=""
		}else if "}"==vi{
			varunit=token
			token=""
		}else if ","==vi{
			amount,_:=strconv.Atoi(varamount)
			mymaterials[key]=unit{amount,varunit}
			token=""
			key=""
			varunit=""
			varamount=""
		}else if "\n"==vi{
			amount,_:=strconv.Atoi(varamount)
			mymaterials[key]=unit{amount,varunit}
			token=""
			key=""
			varunit=""
			varamount=""
		}else{
			token=token+vi
		}
	}
	return mymaterials
}
func Whaticanmakepon(graph map[string]material,objets map[string]unit) string {
	ocurrences:=make(map[string]int)
	max:=0
	item:=""
	fmt.Println("graph")
	fmt.Println(graph)
	fmt.Println("objets")
	fmt.Println(objets)
	for key,_:=range objets{
		fmt.Println(key,graph[key])
		for _,j:= range graph[key].conexions{
			_,ok:=ocurrences[j]
			if objets[key].amount>=graph[key].units.amount{//in the future the units will be managed
				if ok{
					//fmt.Println(ocurrences)
					ocurrences[j]=ocurrences[j]+1
				}else{
					ocurrences[j]=1
				}
			}
		}
	}
	for i,_:=range ocurrences{
		if max<ocurrences[i]{
			max=ocurrences[i]			
			item=i
		}
	}
	fmt.Println("ocurrences")
	fmt.Println(ocurrences)
	return item
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
	fmt.Println(ocurrences)
	return item
}
func Str2arrstr(text string ,sep string)[]string{
	token:=""
	words:=make([]string,0)
	for _,i:= range text{
		i:=string(i)
		if i==sep{
			words=append(words,token)			
			token=""
		}else{
			token=token+i
		}
	}
	words=append(words,token)
	return words
}
