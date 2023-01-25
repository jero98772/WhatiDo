package main
import "fmt"
type unit struct{
	amount int
	amountUnits string
}
type material struct{
	units unit 
	conexions []string

}
/*
func whaticanmake(graph map[material][]string,objets []material) string {
	ocurrences:=make(map[material]int)
	max:=0
	item:=""
	for _,keyi:=range objets{
		for keyj,_:= range graph{
			if keyj.name==keyi.name && keyi.amount >=keyj.amount{
				_,ok:=ocurrences[keyi]
				if ok{
					ocurrences[keyi]=ocurrences[keyi]+1
				}else{
					ocurrences[keyi]=1
				}
			}
		}
	}
	for i,_:=range ocurrences{
		if max<ocurrences[i]{
			max=ocurrences[i]			
			item=i.name
		}
	}
	//fmt.Println(ocurrences)
	return item
}*/
func whaticanmake(graph map[string]material,objets map[string]unit) string {
	ocurrences:=make(map[string]int)
	max:=0
	item:=""
	for key,_:=range objets{
		for _,j:= range graph[key].conexions{
			_,ok:=ocurrences[j]
			if objets[key].amount>=graph[key].units.amount{//in the future the units will be managed
				if ok{
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
	fmt.Println(ocurrences)
	return item
}
func main(){

	graph:=make(map[string]material)
	graph["tierra"]=material{unit{100,"m²"},[]string{"huerta","casa"}}
	graph["madera"]=material{unit{2000,"tablas"},[]string{"huerta","silla","casa"}}
	graph["ladrillo"]=material{unit{20000,"unidades"},[]string{"casa","silla"}}
//-------error dont work correctly
	//ladrillos:=material{"ladrillos",1000,"unidades",[]string{"casa","silla"}}
	objets:=make(map[string]unit)
	objets["tierra"]=unit{1740,"m²"}
	objets["madera"]=unit{965822,"tablas"}
	objets["ladrillo"]=unit{5521200,"unidades"}
	//objets:=[]string{"tierra","madera","ladrillo"}//objets i have
	fmt.Println(whaticanmake(graph,objets))
}
