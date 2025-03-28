package main

import "fmt"

func main (){
	var usuario string
	var senha string

	fmt.Print("digite o usuario")
 	fmt.Scan(&usuario)
	 fmt.Print("digite a senha")
 	fmt.Scan(&senha)
	
	if senha == "1234" && usuario == "admin" {
        fmt.Println("acesso permitido")
		} else {
			fmt.Println("acesso negado")
		}
	}

