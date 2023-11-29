/*
CONCEPTOS BÁSICOS de Go
https://www.youtube.com/watch?v=AGiayASyp2Q
*/
package main

import (
	"fmt"
	"reflect"
	"container/list"
	"github.com/ratara5/basics/function_copy_reference"
)

func main() {
	// Comentario de una linea
	/*
		Comentario de 
		varias lineas
	*/
	fmt.Println("Hola, Go!")

	// Variables

	var myString string = "Esta es una cadena de texto" // Las comillas simples no funcionan para string
	fmt.Println(myString)

	myString = "Aquí cambia el valor de la cadena de texto"
	fmt.Println(myString)

	// myString = 6 // No es posible cambiar el tipo de un dato

	var myString2 = "Go intenta inferir el tipo de un dato"
	fmt.Println(myString2)

	var myInt int = 7 // Por defecto es un int de 32 bits
	fmt.Println(myInt)
	fmt.Println(myInt + 3)

	fmt.Printf("%s %d", myString, myInt)

	fmt.Println("")
	fmt.Println(reflect.TypeOf(myInt)) // Saber el tipo de un dato

	var myFloat float64 = 6.5
	fmt.Println(myFloat)

	fmt.Println(reflect.TypeOf(myFloat))

	fmt.Println(myFloat + float64(myInt)) // El IDE transforma el int a float64

	var myBool bool = true
	myBool = false
	fmt.Println(myBool)

	myString3 := "Este es un string declarado e inicializado con un operador especial"
	fmt.Println(myString3)

	// Constantes
	const myConst = "Esta es una constante" // El programa va a compilar aunque se declare la constante y no se use

	// Control de flujo
	if myInt == 10 && myString == "Hola" {
		fmt.Println("El valor es 10")
	} else if myInt == 7 || myString == "Adios" {
		fmt.Println("El valor es 7")
	} else {
		fmt.Println("El valor no es 10")
	}

	// Array
	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3 // Si no se asigna, por defecto el valor es 0
	fmt.Println(myArray)

	// Mapa
	myMap := make(map[string]int) //[clave]valor
	myMap["Ray"] = 34
	myMap["Yei"] = 26
	myMap["Alv"] = 32
	fmt.Println(myMap)
	fmt.Println(myMap["Ray"])

	myMap2 :=map[string]int{"Ray": 34, "Yei": 26, "Alv": 32} // Esta es otra manera
	fmt.Println(myMap2)

	// Lista: OJO, la concepción de lista en GO no es la típica,
	// es más bien la de puntero, y también funciona como una pila
	myList := list.New()
	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)
	fmt.Println(myList.Back().Value)

	// Bucles
	for index := 0; index < len(myArray); index++ {
		fmt.Println(myArray[index])
	}

	for index, value := range myArray{
		fmt.Println(index, value)
	}

	for key, value := range myMap{
		fmt.Println(key, value)
	}

	// Función
	fmt.Println(myFunction())

	// Estructura (Alternativa para concebir clases, enumerados...)
	type MyStruct struct {
		name string
		age int
	}

	myStruct := MyStruct{"Ray", 34}
	fmt.Println(myStruct)

	//Punteros
		vehiculo1 := "rojo"
		fmt.Println("El vehículo1 es", vehiculo1, &vehiculo1)
	
		vehiculo2 := vehiculo1 //COPIA
		fmt.Println("El vehículo2 es", vehiculo2, &vehiculo2)
	
		vehiculo3 := &vehiculo1 //REFERENCIA
		vehiculo1 = "gris"
		fmt.Println("El vehículo3 es", *vehiculo3, vehiculo3)
	
		//Para ver la dirección de una variable, usar <&> 
		//ASI &var1 /Dirección de var1/
		
		//Para asignar una variable a la dirección de otra variable, usar <&> 
		//ASI var2 := &var1 /var2 es exactamente igual a var1. Cuidado que &var2 != &var1/. ASIGNAR REFERENCIA
		
		//Para llamar el valor de una variable asignada según el procedimiento anterior, usar <*>
		//ASI *var2 /Valor de var2, que SIEMPRE es igual al valor de var1/. LLAMAR REFERENCIA

		//Funciones por copia y referencia
		function_copy_reference.FunctionsCopyReference()
	
}

func myFunction() string {
	return "Mi función"
}
