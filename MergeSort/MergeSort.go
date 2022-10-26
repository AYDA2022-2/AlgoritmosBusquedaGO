package main

import "fmt"

//METODO DE ORDENACION: MERGESORT
func main() {
	var ListaDesordenada = []int{24,57,17,62,41,32,19}//https://www.makingcode.dev/2018/09/implementacion-del-metodo-de-ordenacion_91.html
	ListaOrdenada := mergeSort(ListaDesordenada, 0, len(ListaDesordenada)-1)
	fmt.Println(ListaDesordenada)
}
func mergeSort(ListaDesordenada []int, iMin int, iMax int) []int {
	// Caso Base
	if iMin >= iMax {
		return ListaDesordenada
	}
	// Cortamos para aplicar mergeSort recursivamente
	k := (iMin + iMax) / 2
	mergeSort(ListaDesordenada, iMin, k)
	mergeSort(ListaDesordenada, k+1, iMax)
	// Utilizamos un arreglo temporal
	l := iMax - iMin + 1
	var temp = []int{}
	for i := 0; i < l; i++ {
		temp[i] = ListaDesordenada[iMin+i]
	}
	// Mezclamos
	i1 := 0
	i2 := k - iMin + 1
	for i := 0; i < l; i++ {
		if i2 <= iMax-iMin {
			if i1 <= k-iMin {
				if temp[i1] > temp[i2] {
					ListaDesordenada[i+iMin] = temp[i2+1]
				} else {
					ListaDesordenada[i+iMin] = temp[i1+1]
				}
			} else {
				ListaDesordenada[i+iMin] = temp[i2+1]
			}
		} else {
			ListaDesordenada[i+iMin] = temp[i1+1]
		}
	}
	return ListaDesordenada
}
