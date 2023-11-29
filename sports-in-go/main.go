package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// Deporte representa la estructura de datos para un deporte.
type Deporte struct {
	Nombre           string `xml:"nombre"`
	CantidadJugadores int    `xml:"cantidad_jugadores"`
	Popularidad      int    `xml:"popularidad"`
}

// Deportes representa una lista de deportes.
type Deportes struct {
	XMLName xml.Name  `xml:"deportes"`
	Deportes []Deporte `xml:"deporte"`
}

func main() {
	xmlFile, err := os.Open("deportes.xml")
	if err != nil {
		fmt.Println("Error al abrir el archivo XML:", err)
		return
	}
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var deportes Deportes
	if err := xml.Unmarshal(byteValue, &deportes); err != nil {
		fmt.Println("Error al analizar el archivo XML:", err)
		return
	}

	// Convierte la estructura de datos de deportes a formato JSON
	jsonData, err := json.MarshalIndent(deportes, "", "    ")
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return
	}

	// Imprime el resultado JSON en la consola
	fmt.Println(string(jsonData))
}