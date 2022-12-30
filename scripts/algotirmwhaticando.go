package main
import "fmt"
func printMat(mat [][]int, cols []int)(int,int){
	ocurrences:=make(map[int]int)
	var max int	
	max=0//we take the max on the function for down complexity
	var posmax=-1
	for i, _:= range mat{
		for j, _:= range cols{
			if mat[i][j]==1{
				_,ok:=ocurrences[i]
				if ok{
					ocurrences[i]=ocurrences[i]+1
				}else{
					ocurrences[i]=1
				} 
				//fmt.Print(mat[i][j])
				
			}		
		}
		//fmt.Println("\n...=",max)
		if max<ocurrences[i]{
			max=ocurrences[i]			
			posmax=i
		}
	}
	return max,posmax
}
func main(){
	mat :=[][]int{{0,0,0,0},{0,0,0,0},{1,1,0,0},{1,1,1,1}}//directional matrix as graph
	col :=[]int{0,1}//selected data
	colsname :=[]string{"tierra","madera","huerta","silla"}
	max,pos:=printMat(mat,col)	
	fmt.Println(max,pos)
	fmt.Println(colsname[pos])
}