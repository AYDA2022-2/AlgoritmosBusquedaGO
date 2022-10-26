package main

import "fmt"

func main(){
	
	listaDesordenada := []int {24,57,17,62,41,32,19}

	listaOrdenada := Burbuja(listaDesordenada[:])
	
	for i := 0; i < len(listaOrdenada); i++ {
		var aux int = listaOrdenada[i]
		fmt.Println(aux)
	}
}

func Burbuja(ListaDesordenada []int) []int {//https://www.makingcode.dev/2018/09/implementacion-del-metodo-de-ordenacion.html
	var auxiliar int
	for i := 0; i < len(ListaDesordenada); i++ {
	 	for j := 0; j < len(ListaDesordenada); j++ {
	  		if ListaDesordenada[i] > ListaDesordenada[j] {
	   			auxiliar = ListaDesordenada[i]
	   			ListaDesordenada[i] = ListaDesordenada[j]
	   			ListaDesordenada[j] = auxiliar
	  		}
		}
	}
	return ListaDesordenada
}
