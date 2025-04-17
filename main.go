package main

import (
   "fmt"
)

func dividir (dividendo int, divisor int)(int, string){
   if divisor == 0 {
   return 0,"erro na divisao por zero"
}
   return dividendo / divisor, "Sem erro"
} 
func operacaoBasica(a int, b int,)(int, int, int){
   soma := a + b
   multiplicacao := a * b
   subtracao := a - b
   return soma, multiplicacao, subtracao
}
func main(){
   resultado, erro := dividir(10,2)

   if erro != "Sem erro"{
fmt.Printf(erro)
   } else{
      fmt.Println("o resultado da divisão é:",resultado, erro)
   }
   soma, mult, sub := operacaoBasica(10,2)
   fmt.Println(soma)
   fmt.Println(mult)
   fmt.Println(sub)  
}