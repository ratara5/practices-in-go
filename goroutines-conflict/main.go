package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
	"github.com/ratara5/goroutines-conflict/data"
)

func main() {

	start := time.Now()
	wg := &sync.WaitGroup{}
	m := &sync.Mutex{} //Esta variable se usa con otros comandos asociados en puntos convenientes del código determinados por go run --race main.go, con el objetivo de solucionar conflictos como los que se presentan cuando por la ejecución simultanea de routines se intentan leer y escribir valores al mismo tiempo

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go readBook(i, wg, m)
	}


	wg.Wait() 
	duration := time.Since(start).Milliseconds() 
	fmt.Printf("%dms", duration)

	//Otro problema de las goroutines es que a veces una función puede estar intentando acceder a determinado valor, mientras otra puede estar intentando modificarlo (conflicto)
	//Correr go run --race main.go ===> Esto mostrará que existe un conflicto...
}

func readBook(id int, wg *sync.WaitGroup, m *sync.Mutex) {

	data.FinishBook(id, m) //...En un caso, las funciones findBook() y FinishBook() del paquete data están escribiendo y leyendo algo al mismo tiempo. Para solucionar este problema, se usa Mutex en la función FinishBook()...

	delay := rand.Intn(500)
	time.Sleep(time.Millisecond * time.Duration(delay))

	wg.Done()
}