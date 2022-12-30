//whatIDo - by jero98772
package main
import ("fmt"
	"net/http"
	"html/template"
	"jerotools/tools"
	"log"
	)
type Answer struct {
	Text string
	isDonde bool
}
func Index(rw http.ResponseWriter,r *http.Request){
	template, _ := template.ParseFiles("templates/index.html")
	template.Execute(rw,nil)
}
func Playground(rw http.ResponseWriter,r *http.Request){
	err := r.ParseForm()
	data:=""
	state:=false
	if err != nil {
		log.Fatal(err)
	}
	if len(r.Form)<=0{
		data=""
		state=false
	}else{
		graph:=tools.Str2Graph(r.Form["allmaterials"][0]+"\n")
		mymaterials:=tools.Str2arrstr(r.Form["mymaterials"][0],",")
		data=tools.Whaticanmake(graph,mymaterials)
		state=true
	}
	choice:=Answer{Text:data,isDonde:state}
	fmt.Println(choice)
	template, _ := template.ParseFiles("templates/playground.html")
	template.Execute(rw,choice)
}
func main(){
	tools.PrintHello()
	http.HandleFunc("/",Index)
	http.HandleFunc("/playground",Playground)
	//http.HandleFunc("/playonline",playonline)
	port :="9600"
	fmt.Println("run at localhost:"+port)
	http.ListenAndServe("0.0.0.0:"+port,nil)

}