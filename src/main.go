package main

import (
	"net/http"

	"github.com/jchpoli/go-curso-jch/src/ejChannels"
	"github.com/labstack/echo/v4"
)

//Alias "ruta a paquete"
//"./arraySliceMapRange"
//"./sentenciasFunciones"
//"./structPunteroInterfaces"
//"./variablesOperaciones"

//"./ejChannels"
//"./ejGorutines"

func main() {
	//variablesOperaciones.Main()
	//sentenciasFunciones.Main()
	//arraySliceMapRange.Main()
	//structPunteroInterfaces.Main()
	//ejGorutines.Main()
	ejChannels.Main()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hola mundo")
	})
	e.Logger.Fatal(e.Start(":1323"))

}
