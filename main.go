package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type task struct {
	id        int
	nombre    string
	contenido string
}

var allTasks = make(map[int]task)

func main() {
	var tarea int = 1

	for {
		fmt.Printf("Ingrese datos para la tarea %d:\n", tarea)
		insertTask(tarea)
		tarea++
	}
}

func insertTask(id int) {
	var nombre, contenido string

	fmt.Print("Nombre: ")
	nombre = readLine()

	fmt.Print("Contenido: ")
	contenido = readLine()

	fmt.Print("¿Desea agregar otra tarea? (y/n): ")
	respuesta := readLine()

	if strings.ToLower(respuesta) == "y" {
		t := task{
			id:        id,
			nombre:    nombre,
			contenido: contenido,
		}
		allTasks[id] = t
	} else {
		// Imprimir todas las tareas después de la inserción cuando el usuario decide no agregar más tareas.
		fmt.Println("Todas las tareas:", allTasks)
		os.Exit(0) // Salir del programa después de imprimir las tareas.
	}
}

func readLine() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
