package arraySliceMapRange

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) bool {
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}
	var esPalindromo bool = strings.EqualFold(text, textReverse)
	if esPalindromo {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
	return esPalindromo
}

func Main() {
	fmt.Println("------------- ARRAY Y SLICES-------------")
	// Array no se puede modificar
	// Por defecto todos los indices toman el valor por decto según  el tipo de dato
	var array1 [4]int
	fmt.Println(array1)
	fmt.Println(array1, len(array1), cap(array1))

	// Array se puede modificar
	var slice1 = []int{0, 1, 2, 3, 4, 5, 6}

	fmt.Println(slice1, len(slice1), cap(slice1))

	//Métodos en slice

	fmt.Println(slice1[0], "Elemento del indice 0")
	fmt.Println(slice1[:3], "Elementos hasta el indice 3 inclusivo")
	fmt.Println(slice1[2:4], "Elementos del 2 al 4 exclusivo")
	fmt.Println(slice1[4:], "Elementos del 4 inclusivo en adelante")

	//Append, Agregar un nuevo/s elemento al final
	slice1 = append(slice1, 7, 8, 9)
	fmt.Println(slice1)

	//Append, Agregar un nuevo/s elementos de un slice ya existente
	newSlice := []int{2, 11, 12}
	slice1 = append(slice1, newSlice...)
	fmt.Println(slice1)

	fmt.Println("------------- RANGE-------------")

	for i, valor := range slice1 {
		fmt.Println(valor, i)
	}

	for _, valor := range slice1 {
		fmt.Println(valor)
	}

	for indice := range slice1 {
		fmt.Println(indice)
	}

	isPalindromo("ama")
	isPalindromo("casa")
	isPalindromo("holaloh")

	fmt.Println("------------- MAPS-------------")

	// make nos permite crear diccionaros, este recive el tipo de dato y la cantidad
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	// Recorrer map, este lo hace en desorden
	for i, v := range m {
		fmt.Printf("Llave %s valor %d \n", i, v)
	}
	/* El segundo retorno es un booleano que nos dice si existe o no el indice
	ya que por defectyo si no existe el indice devuelve el valor por defecto
	según el tipo de dato
	*/
	valor, existe := m["Jose"]
	fmt.Println("Valor de la llave Jose ", valor, existe)
	valor, existe = m["otro"]
	fmt.Println("Valor de la llave otro ", valor, existe)
}
