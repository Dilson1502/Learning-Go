// Accediendo al paquete main, el cual tiene el contexto de inicialización princiapal

package main

// Importando el paquete format
import (
	"fmt"
)

func main() {
	// Funciones
	// Declaramos la función fuera de la función main, luego se llama dentro de la función main
	fmt.Println(myFunction())
}

// creando funciones, se debe crear fuera de la función main
func myFunction() string {
	return "Mi función"
}
