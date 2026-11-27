package main

import "fmt"

const appName = "Curso Go 2.0" //app = interno, App = exportable
const maxUsers = 1000

func constans() {
	var name string = "Enrique"
	last_name := "Garcia"

	fmt.Printf("Nombre: %s %s %T\n", name, last_name, last_name)
	fmt.Println(appName, maxUsers)
}
