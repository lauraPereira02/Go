package main

import "fmt"

 func main(){
 var ages = [4]int{16,20,24,28}
 nomes:= [4]string{"mario","luigi","superman","pool"}
 fmt.Println(ages)
 fmt.Println(nomes) 
 nomes[3] = "clara"
 fmt.Println(nomes)
// Slice
var score = []int{100,200,300,400}
fmt.Println(score)
score[1] = 2
fmt.Println(score, len(score), cap(score))
rengeOne := score [1:3]
fmt.Println(rengeOne)
rangetwo := score[2:]
fmt.Println(rangetwo)
rangethree := score [:3]
fmt.Println(rangethree)

var superherois = []string{"homem aranha","homem de ferro","thor"}
fmt.Println(superherois)
superherois = append(superherois, "bem dez","huck")
fmt.Println(superherois, len(superherois), cap(superherois))
}
