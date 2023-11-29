package data

import (
		"fmt"
		"sync"
)

type Book struct {
	ID       int
	Title    string
	Finished bool
}

var books = []*Book{
	{1, "Dune", false},
	{2, "El Perfume", false},
	{3, "The World of Ice and Fire", false},
	{4, "Teoría de la noche", false},
	{5, "Blanca Olmedo", false},
	{6, "El principito", false},
	{7, "100 años de soledad", false},
	{8, "El alquimista", false},
	{9, "El libro del cementerio", false},
	{10, "Maze runner", false},
}

func findBook(id int, m *sync.Mutex) (int, *Book) { //Ver comentarios en la función FinishBook. Si se usara m *sync.RWMutex...
	index := -1
	var book *Book

	m.Lock() //...Esta instrucción cambia a m.RLock() que es específica para las lecturas de valores (no confundir con readBook)
	for i, b := range books {
		if b.ID == id {
			index = i
			book = b
		}
	}
	m.Unlock() //...y esta instrucción cambia a m.RUnlock() que es específica para las lecturas de valores

	return index, book
}

func FinishBook(id int, m *sync.Mutex) { //antes sin m *sync.Mutex había conflicto pues se escribía y se leía un valor al mismo tiempo, ahora con el m *sync.Mutex y los respectivos Lock() y Unlock() Se bloquea el valor a modificar y ya no hay conflicto. //Si acá se pusiera m *Sync.RWMutex no habría cambios en Lock() y Unlock(), cosa que sí sucede en findBook. También habría que cambiar a RWMutex la variable m y las llamadas en las funciones superiores a esta
	i, book := findBook(id, m)
	if i < 0 {
		return
	}

	m.Lock() //Para solucionar conflicto de lectura y escritura simultanea: Se bloquea antes de escribir (?)
	book.Finished = true
	books[i] = book
	m.Unlock() //Para solucionar conflicto de lectura y escritura simultanea: Se descloquea después de escribir (?)

	fmt.Printf("Fisnished Book: %s\n", book.Title)

}
