package main

import "fmt"

const AppName = "Curso Go" //app = interno, App = exportable
const MaxUsers = 1000

func main() {
	var name string = "Enrique"
	last_name := "Garcia"

	fmt.Printf("Nombre: %s %s %T\n", name, last_name, last_name)
}
