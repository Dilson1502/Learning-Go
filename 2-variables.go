// Accediendo al paquete main, el cual tiene el contexto de inicialización princiapal

package main

// Importando el paquete format
import (
	"container/list"
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hola mundo. Estoy usando Go!")
	// Variables
	// var name type
	var myString string = "Esto es una cadena de texto"
	fmt.Println(myString)

	// Las variables pueden cambiar su valor
	myString = "Aqui cambio el valor de la cadena de texto"
	fmt.Println(myString)

	// podemos declarar la variable y asignarle un valor en la misma linea
	var myString2 string = "Esto es otra forma de declarar variables"
	fmt.Println(myString2)

	// No se puede cambiar el tipo de dato de una variable, el código a continuación retornará un error
	// myString = 6 → Error

	// exites varios tipos de tipos de datos numéricos según sus bits, Int8.Int32,Int64, etc. Saber el numero de bits del tipo de dato y limitarlo hará más eficiente el código
	var myInt = 7
	myInt = myInt + 4
	// imprimiendo la variable luego de la operación
	fmt.Println(myInt)
	// operando en tiempo de ejecución sin afectar la variable myInt
	fmt.Println(myInt - 1)
	// revisemos que no se afecta el valor original de la variable myInt por la operación en ejecución anterior
	fmt.Println(myInt)
	// Formateando una cadena de texto para que pueda imprimir un string + int
	fmt.Printf("%s %d", myString, myInt)
	// usando el paquete reflect para saber el tipo de dato de una variable
	fmt.Println(reflect.TypeOf(myInt))
	// creando una variable de tipo punto flotante
	var myFloat float32 = 6.5
	// transformando a float un float
	fmt.Println(myInt + int(myFloat))
	// creando un booleano
	var MyBool bool = false
	MyBool = true
	fmt.Println(MyBool)
	// creando variables asignadas automáticamente
	myString3 := "Esto es una cadena de texto"
	fmt.Println(myString3)
	// creando constante
	const myConst = "Esto es una constante"
	fmt.Println(myConst)
	// Control de flujo
	myInt = 10
	myString = "Hola"

	if myInt == 10 && myString == "Hola" {
		fmt.Println("El valor es 10")
	} else if myInt == 11 || myString == "Hola" {
		fmt.Println("El valor es 11")
	} else {
		fmt.Println("El valor no es 10")
	}
	// creando arrays con enteros
	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	fmt.Println(myArray)
	// creando arrays con punto flotante
	var myArray2 [3]float64
	myArray2[0] = 1.5
	myArray2[1] = 2.5
	myArray2[2] = 3.5
	// Al declarar la variable se indica que es un arreglo de 3 posiciones, al tratar de ingresar un cuarto valor retornará un error, dado que los índice del array van de 0 a 3 →[0:3]
	// myArray2[3]= 4.5
	// Lo mismo sucedería al intentar imprimir la 4ta posición
	// fmt.Println(myArray2[3])
	fmt.Println(myArray2)
	// creando mapas (clave:valor), similar a un diccionario
	myMap := make(map[string]int)
	myMap["Dilson"] = 24
	myMap["Orlando"] = 61
	fmt.Println(myMap)
	// creando listas
	myList := list.New()
	myList.PushBack(1)
	myList.PushFront(2)
	fmt.Println(myList.Back().Value)
	// creando bucles
	for index := 0; index < len(myArray); index++ {
		fmt.Println(myArray[index])
	}
	// iterando un mapa (diccionario)
	for key, value := range myMap {
		fmt.Println(key, value)
	}
	// Funciones
	// Declaramos la función fuera de la función main, luego se llama dentro de la función main
	fmt.Println(myFunction())

	// Estructura, es una forma de crear objetos, no confundir con las clases
	type MyStruct struct {
		name string
		age  int
	}
	myStruct := MyStruct{"Dilson", 24}
	fmt.Println(myStruct.age)
}

// creando funciones, se debe crear fuera de la función main
func myFunction() string {
	return "Mi función"
}
