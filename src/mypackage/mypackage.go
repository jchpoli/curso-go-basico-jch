// Por estandar el archivo principal se llama como la carpeta del paquete
package mypackage

import "fmt"

// CarPublic con acceso público
type CarPublic struct {
	Brand string // Si inicia con mayúscula es que la propiedad es pública
	Year  int
}

// carPrivate con acceso privado
type carPrivate struct {
	brand string // Si inicia con minúscula es que la propiedad es privada
	year  int
}

// función publica
func PrintMessage(message string) {
	fmt.Println(message)
}

// función privada
func printMessage(message string) {
	fmt.Println(message)
}
