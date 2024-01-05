package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Person struct {
	ID      string
	Name    string
	Country string
}

func main() {

	// Open a file as reader mode
	// file es *os.File, e implementa io.Reader (por eso se puede usar como reader)
	file, err := os.Open("customers.csv")

	if err != nil {
		panic("The indicated file was not found or is damaged")
	}

	defer file.Close()

	reader := csv.NewReader(file)

	// skip header
	reader.Read()

	listaPersonas := []Person{}
	for {
		// record is a slice of strings (el metodo implementa la lectura y el manejo de buffer interno)
		record, err := reader.Read()
		if err != nil {
			break
		}
		fmt.Println(record)
		listaPersonas = append(listaPersonas, Person{
			ID:      record[0],
			Name:    record[1],
			Country: record[2],
		})
	}

	fmt.Println(listaPersonas)

}
