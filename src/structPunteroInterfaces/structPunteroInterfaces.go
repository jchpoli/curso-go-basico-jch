package structPunteroInterfaces

import (
	"fmt"

	"../mypackage"
)

// Una manera extender funcionalidad de un struct como añadir funciones
func (elPc pc) ping() {
	fmt.Println(elPc.ram, "ping")
}

// Por puntero puede modificar el objeto
func (elPc *pc) duplicateRam() {
	elPc.ram = elPc.ram * 2
	fmt.Println(elPc.ram, "multiplicada")
}

type car struct {
	brand string
	year  int
}

type pc struct {
	ram  int
	disk int
}

// Se personaliza la impresión en string de un struct
func (elPC pc) String() string {
	return fmt.Sprintf("Tengo %d GB de RAM, %d GB de Disco .", elPC.ram, elPC.disk)
}

// Intefaces
type figuras2D interface {
	area() float64
}

type cuadrado struct {
	base float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

type rectangulo struct {
	base   float64
	altura float64
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func calcular(f figuras2D) {
	fmt.Println("Area:", f.area())
}

func Main() {
	fmt.Println("------------- STRUCT (Clases)-------------")
	// Forma 1 de instanciar
	myCar := car{brand: "Ford", year: 2021}
	fmt.Println(myCar)
	var otherCar car
	otherCar.brand = "Ferrari"
	// year toma por defecto su valor por defecto
	fmt.Println(otherCar)

	var myCar2 mypackage.CarPublic
	myCar2.Brand = "Ferrari"
	myCar2.Year = 2020
	fmt.Println(myCar2)

	/* //Arroja errores porque la clase es privada en el paquete
	var myCar3 mypackage.carPrivate
	myCar3.brand = "Ferrari"
	myCar3.year = 2020
	fmt.Println(myCar3)
	*/
	mypackage.PrintMessage("holas")

	fmt.Println("------------- PUNTERO, Stringers Y STRUCT -------------")
	t := 50
	u := &t // Se busca la posición de memoria donde esta t

	fmt.Println(t)
	fmt.Println(u)
	fmt.Println(*u) // con * se accede al valor que está en la dirección de memoria
	// Cambiamos el valor y se modifica las demás variables que apuntan a esa posición
	*u = 100

	fmt.Println(t)
	fmt.Println(*u)

	myPc := pc{ram: 16, disk: 250}
	fmt.Println(myPc)

	myPc.ping()
	myPc.duplicateRam()
	myPc.duplicateRam()

	fmt.Println(myPc)

	fmt.Println("------------- INTERFACES Y LISTAS DE INTERFACES 	-------------")

	myCuadrado := cuadrado{base: 2}
	myRectangulo := rectangulo{base: 3, altura: 4}

	// Si coincide todos los métodos no fallará
	calcular(myCuadrado)
	calcular(myRectangulo)

	// Lista de interfaces
	/* Permite crea un objeto con diferentes tipos de datos
	estilo lista
	*/
	myInterface := []interface{}{"Hola", 123, true, 4.2}
	// Los 3 puntos para imprimir cada uno de los elementos de forma individual
	fmt.Println(myInterface...)
}
