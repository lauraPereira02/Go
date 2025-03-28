package main

import "fmt"

func main (){
	var numero1 int
	var numero2 int

	fmt.Print("digite um numero")
 	fmt.Scan(&numero1)
	 fmt.Print("digite outro numero")
 	fmt.Scan(&numero2)
 	
	fmt.Println("a soma é", numero1 + numero2)
	fmt.Println("a subtraçao é:", numero1 - numero2)
	fmt.Println("a divisao é ", numero1 / numero2)
	fmt.Println("a multiplicaçao é:", numero1 * numero2)
	fmt.Println("o resto da divisao:", numero1 % numero2)
	
}