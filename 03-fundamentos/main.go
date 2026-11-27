package main

import "fmt"

const isv = 0.15

func main() {
	var clientName string
	var clientEmail string
	var subTotal float64

	fmt.Println("Ticket por consola")
	fmt.Println("Nombre del cliente: ")
	fmt.Scan(&clientName)
	fmt.Println("Email del cliente: ")
	fmt.Scan(&clientEmail)
	fmt.Println("Subtotal de productos: ")
	fmt.Scan(&subTotal)

	fmt.Println("*****************")
	fmt.Printf("Subtotal: %.2f\n", subTotal)
	fmt.Printf("Impuestos: %.2f\n", subTotal*isv)
	fmt.Printf("total: %.2f\n", subTotal*(1+isv))
	fmt.Println("*****************")

}
