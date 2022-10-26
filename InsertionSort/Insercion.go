package main

import "fmt"

func main(){
	
	listaDesordenada := []int {24,57,17,62,41,32,19}

	listaOrdenada := Insercion(listaDesordenada[:])
	
	for i := 0; i < len(listaOrdenada); i++ {
		var aux int = listaOrdenada[i]
		fmt.Println(aux)
	}
}

func Insercion(ListaDesordenada []int) []int { //https://www.makingcode.dev/2018/09/implementacion-del-metodo-de-ordenacion_20.html
	var auxiliar int
	for i := 1; i < len(ListaDesordenada); i++ {
	 	auxiliar = ListaDesordenada[i]
	 	for j := i - 1; j >= 0 && ListaDesordenada[j] > auxiliar; j-- {
	  		ListaDesordenada[j+1] = ListaDesordenada[j]
	  		ListaDesordenada[j] = auxiliar
	 	}
	}
	return ListaDesordenada
   }

