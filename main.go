package main

import "fmt"

func main() {
	age := 45
	fmt.Println(age<= 50)
   fmt.Println(age >= 50)
   fmt.Println(age == 50)
   fmt.Println(age!= 50)

   if age < 50 {
      fmt.Println("menor que 30anos")
   } else if age < 40 {
      fmt.Println("menor que 40anos")
   } else {
      fmt.Println("não é menor que 40anos")
   }

   names := []string{"Isa", "madu", "maju", "leo", "lala", "lili" }

   for index, value := range names {
      if index == 1 {
         fmt.Println("continue apos a posiçao", index, "e valor", value)
         continue
      }
      if index > 2 {
         fmt.Println("sair apos" ,index)
      }
      fmt.Println("valor: ", value)
   }
}