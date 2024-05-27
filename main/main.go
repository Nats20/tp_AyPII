package main

import (
	"TP--nosehizoconchatgpt/menu"
)

func main() {
	menu := Menu.CrearMenu()
	menu.EncenderMenu()
}

// se podria añadir una referencia a la tarea padre para indicar si
// lo tiene -> GRAFO

// opc 1: tener todas las tareas al mismo nivel
// opc 2: buscar recursivamente y tener en la lista solo el primer nivel
