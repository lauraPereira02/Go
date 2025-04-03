package main

import "fmt"

 func main(){
 var ages = [4]int{16,20,24,28}
 nomes:= [4]string{"mario","luigi","superman","pool"}
 fmt.Println(ages)
 fmt.Println(nomes) 
 nomes[3] = "clara"
 fmt.Println(nomes)

 }