package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync"
)

func main() {

	start := time.Now()
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go showGoroutine(i, wg) //Sin el go, las funciones se ejecutan en orden de i. Con el go, las funciones se ejecutan en desorden (en el orden que encuentran disponible)
	}

	//time.Sleep(time.Minute) //Sin esta instrucción, el programa termina inmediatamente. Este tiempo de 1 minuto fue agregado de forma arbitraria, ya que en realidad no sabemos cuánto tiempo tardan todas las funciones en terminar...

	wg.Wait() //...Para solucionar este problema, Go nos ofrece un paquete llamado WaitGroup. Esta instrucción espera a que las wg asociadas a este grupo terminen de ejecutarse.

	duration := time.Since(start).Milliseconds() //sin usar go, este valor es de 2400ms, con go este valor es de 450ms. Entonces la palabra go, usa la concurrencia para ejecutar las funciones prácticamente que a la vez
	fmt.Printf("%dms", duration)

	//Otro problema de las goroutines es que a veces una función puede estar intentando acceder a determinado valor, mientras otra puede estar intentando modificarlo (conflicto)
}

func showGoroutine(id int, wg *sync.WaitGroup) {
	delay := rand.Intn(500)
	fmt.Printf("Goroutine #%d\n with %dms\n", id, delay)
	time.Sleep(time.Millisecond * time.Duration(delay))
	wg.Done()
}