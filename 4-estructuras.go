// Accediendo al paquete main, el cual tiene el contexto de inicialización princiapal

package main

// Importando el paquete format
import (
	"fmt"
)

func main() {
	// Estructura, es una forma de crear objetos, no confundir con las clases
	type MyStruct struct {
		name string
		age  int
	}
	myStruct := MyStruct{"Dilson", 24}
	fmt.Println(myStruct.age)
}
