package sentenciasFunciones

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a int, b int, c string) {
	fmt.Println(a, b, c)
}

func tripleArgumentConMismoTipoDato(a, b int, c string) {
	fmt.Println(a, b, c)
}

func multiplyByTow(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a * 4, a * 2
}

func Main() {

	fmt.Println("-------------FUNCIONES-------------")

	normalFunction("Hola mundo")

	tripleArgument(1, 2, "Hola")
	tripleArgumentConMismoTipoDato(1, 2, "Hola")

	value := multiplyByTow(4)
	fmt.Println("Valor:", value)

	value1, value2 := doubleReturn(5)
	fmt.Printf("Valor uno %d, valor dos %d \n", value1, value2)

	// Descartar un valor que retorna se coloca "_"
	value1, _ = doubleReturn(5)
	fmt.Printf("Valor uno nada mas %d \n", value1)

	fmt.Println("-------------CICLOS-------------")

	// For condicional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("")

	// For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	// For forever
	/*
		counterForever := 0

		for {
			fmt.Println(counterForever)
			counterForever++
		}
	*/

	fmt.Println("-------------IF-------------")

	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// With and
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es verdad")
	}

	// With or
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("Es verdad")
	}

	fmt.Println("-------------SWITCH-------------")
	modulo := 45 % 2
	switch modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("No es par")
	}

	// Con asignación
	switch modulo = 22 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("No es par")
	}

	// Sin condición
	value = 200
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No condicion")
	}

	fmt.Println("------------- defer, break y continue-------------")

	// Defer, antes de morir la funcion actual ejecuta la linea de código
	defer fmt.Println("Hola por el defer")
	fmt.Println("Mundo")
}
