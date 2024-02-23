package main

import (
	"fmt"
	"os"
)

func switchcase(str string) string {
	 tr :=[]rune(str)
	for i:=0;i<len(tr);i++{
		if tr[i]>='a' && tr[i]<='z'{
			tr[i]-=32
		}else if tr[i]>='A' && tr[i]<='Z'{
			tr[i]+=32
		}
	}
	return string(tr)
}

func main(){
	arr :=os.Args[1]
	fmt.Println(switchcase(arr))
}