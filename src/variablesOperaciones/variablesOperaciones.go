package variablesOperaciones

import (
	"fmt"
)

func Main() {
	// Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415
	fmt.Println("-------------CONSTANTES-------------")
	fmt.Println("Hello, World")
	fmt.Println("1.Pi:", pi)
	fmt.Println("2.Pi:", pi2)

	fmt.Println("-------------VARIABLES-------------")
	// Declaración de variables
	// NO SE PUEDEN DECLARAR VARIABLES QUE NO SE UTILIZAN
	base := 12          // Crea la variable si no existe
	var altura int = 14 // Crea e inicializa
	var area int        // Crea variable

	fmt.Println(base, altura, area)

	fmt.Println("-------------ZERO VALUES-------------")
	// Zero values
	// GO por defecto les da un valor a las variables
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	fmt.Println("-------------OPERACIONES-------------")
	// Area de cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado:", areaCuadrado)

	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("Suma:", result)

	// Resta
	result = x - y
	fmt.Println("Resta:", result)

	// Multiplicación
	result = x * y
	fmt.Println("Multiplicación:", result)

	// División
	result = y / x
	fmt.Println("División:", result)

	// Modulo
	result = y % x
	fmt.Println("Modulo:", result)

	// Incremental
	x++
	fmt.Println("Incremental:", x)

	// Decremental
	x--
	fmt.Println("Decremental:", x)

	fmt.Println("-------------fmt-------------")

	helloMessage := "Hello"
	wordlMessage := "world"

	// Println
	fmt.Println(helloMessage, wordlMessage)

	// Printf
	nombre := "Programar"
	var horas int16 = 500
	fmt.Printf("%s tiene mas de %d cursos.\n", nombre, horas)
	fmt.Printf("%v tiene mas de %d cursos. con v si no se sabe el tipo de datos\n", nombre, horas)

	// Sprintf

	message := fmt.Sprintf("%s tiene mas de %d cursos.\n", nombre, horas)
	fmt.Println("Con el mensaje retornado con Sprintf:", message)

	// Tipo datos
	fmt.Printf("tipo de dato de 'message' es %T\n", message)
	fmt.Printf("tipo de dato de 'horas' es %T\n", horas)
}
