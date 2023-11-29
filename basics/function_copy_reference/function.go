package function_copy_reference

import "fmt"

func equivalenciaEnPies(altura float32) float32 { //PARÁMETROS POR COPIA
		altura = altura * 3.28
		return altura
	}

func alturaPorEdad(h *float32) float32 { //PARÁMETROS POR REFERENCIA. LLAMAR POR REFERENCIA
    *h = *h - 0.10 //Acá se llama y se asigna por referencia
    return *h
}


func FunctionsCopyReference() {
    
	var altura float32 = 1.70
	fmt.Println("La altura es:", altura, "mts")
	fmt.Println("La altura es:", equivalenciaEnPies(altura), " pies") //ARGUMENTOS POR COPIA
	fmt.Println("La nueva altura es:", altura, "mts") //El valor original de la variable altura NO cambia

	fmt.Println("\n")

	var h float32 = 1.70
	fmt.Println("La altura es:", h, "mts")
	fmt.Println("Al envejecer:", alturaPorEdad(&h), "mts") //ARGUMENTOS POR REFERENCIA. ASIGNAR POR REFERENCIA
	fmt.Println("Despues de envejecer:", h, "mts") //El valor original de la variable h SI cambia

}
