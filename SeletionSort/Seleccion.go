package main

import "fmt"

func main(){
	
	listaDesordenada := []int {24,57,17,62,41,32,19}

	listaOrdenada := Seleccion(listaDesordenada[:])
	
	for i := 0; i < len(listaOrdenada); i++ {
		var aux int = listaOrdenada[i]
		fmt.Println(aux)
	}
}

func Seleccion(arreglo []int) []int { //https://awebytes.wordpress.com/2020/04/12/golang-algoritmos-ordenacion-implementados-bases/
	for i := 0; i < len(arreglo); i++ {
		minimo_encontrado, posicion_minimo := arreglo[i], i

		valor_original := arreglo[i]
		// encontrar minimo en parte desordenada
		for j := i + 1; j < len(arreglo); j++ {
			valor_comparacion := arreglo[j]
			if valor_comparacion < minimo_encontrado {
				minimo_encontrado, posicion_minimo = valor_comparacion, j
			}
		}

		if minimo_encontrado != valor_original {
			// intercambio posiciones con primer desordenado
			arreglo[i], arreglo[posicion_minimo] = minimo_encontrado, valor_original
		}
	}

	return arreglo
}
