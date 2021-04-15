package ejChannels

import (
	"fmt"
)

func Main() {
	c := make(chan string, 3)
	c <- "Mensaje1"
	c <- "Mensaje2"

	fmt.Println("Hay ", len(c), " datos y la cantidad máxima de almacenamiento es ", cap(c))

	// Close, cierra el canal para que no se puedan agregar mas datos
	//close(c)
	/*
		close(c)
		c <- "Mensaje3" //Fallará
	*/

	// Recorre los elementos de un canal, es necesario que este esté cerrado
	close(c)
	for message := range c {
		fmt.Println(message)
	}

	// Select

	email1 := make(chan string)
	email2 := make(chan string)

	go message("mensaje 1", email1)
	go message("mensaje 2", email2)
	// aqui no se sabe cual se va a ejecutar primero

	/* Con select capturamos los canales y escogemos con los case
	el canal deseado para realizar la operaciob que se requiera por cada canal
	*/
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2", m2)
		}
	}

}

func message(text string, c chan string) {
	c <- text
}
