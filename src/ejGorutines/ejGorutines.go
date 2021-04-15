package ejGorutines

import (
	"fmt"
	"sync"
)

func say(text string) {
	fmt.Println(text)
}

func Main() {
	//conSyncWaitGroup()
	//noEjecutadaWorldPorTerminacionDeHilo()
	conChannels()
}

func noEjecutadaWorldPorTerminacionDeHilo() {
	fmt.Println("con noEjecutadaWorldPorTerminacionDeHilo")
	say("Hello")
	// Gorutine, ejecutará la linea de forma concurrente
	go say("World")
	/* Se puede agrega un tiempo de espera con
	time.Sleep(time.Second * 1)
	pero es ineficiente
	*/
}

func saySync(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func conSyncWaitGroup() {
	fmt.Println("con conSyncWaitGroup")
	var wg sync.WaitGroup
	fmt.Println("Hello")
	wg.Add(1) // Se le dice que se va a agregar una gorutine

	// Gorutine, ejecutará la linea de forma concurrente
	go saySync("World", &wg)

	// Se puede utilizar también una función anonima
	go func(text string) {
		fmt.Println("En funcion anonima ", text)
	}("adios")

	/* El grupo esperará a que se ejecute todas la gorutines
	del grupo para seguit ejecutandos despues de esta línea
	*/
	wg.Wait()
	fmt.Println("Hello 4")
	/* Se puede agrega un tiempo de espera con
	time.Sleep(time.Second * 1)
	pero es ineficiente
	*/
}

// No se define en el parámetro (channel) va a ser de entrada o de salida
func sayConChannelSinEspecificarDato(text string, c chan string) {
	// Se le dice que se va a ingresar un dato en ese canal
	c <- text
}

// Se define en el parámetro (channel) va a ser de entrada
func sayConChannelDeEntrada(text string, c chan<- string) {
	// Se le dice que se va a ingresar un dato en ese canal
	c <- text
}

// Se define en el parámetro (channel) va a ser de entrada
func sayConChannelDeSalida(text string, c <-chan string) {
	// Se le dice que se va a ingresar un dato en ese canal
	text = <-c
}

func sayConChannel(text string, c chan<- string) {
	// Se le dice que se va a ingresar un dato en ese canal
	c <- text
}

func conChannels() {
	/* Crea un channel, que procesa datos string y la cantidad de datos que va
	a procesar es 1, este último parámetro es opcional
	*/
	c := make(chan string, 1)

	fmt.Println("Hello")

	go sayConChannel("Bye", c)

	// Para que espere solo se debe obtener el dato del channel
	fmt.Println(<-c)
}
