package lista

import (
	"TP--nosehizoconchatgpt/tarea"
	"fmt"
	"strings"
)

type node struct {
	value *tarea.Tarea
	next  *node
}

func newNode(value *tarea.Tarea) *node {
	return &node{value: value, next: nil}
}

type LinkedList struct {
	head    *node
	tail    *node
	orderBy int
	size    int64
}

func NewLinkedList() *LinkedList {
	return &LinkedList{head: nil, tail: nil, orderBy: DURACION, size: 0}
}

func (l *LinkedList) Append(value *tarea.Tarea) {
	newNode := newNode(value)
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		l.size++
		return
	}
	l.tail.next = newNode
	l.tail = newNode
	l.size++
}

func (l *LinkedList) EliminarTarea(posicion int64) {
	if posicion < 0 || posicion >= l.size {
		return
	}
	// Eliminar el primer nodo
	if posicion == 0 {
		l.head = l.head.next
		l.size--
		return
	}

	current := l.head
	for current != nil && posicion > 1 {
		current = current.next
		posicion--
	}

	current.next = current.next.next
	if current.next == nil {
		l.tail = current
	}
	l.size--
}

func (l *LinkedList) GetOrderBy() int {
	return l.orderBy
}

func (l *LinkedList) SetOrderBy(orderBy int) {
	l.orderBy = orderBy
}

func (l *LinkedList) Size() int64 {
	return l.size
}

func (l *LinkedList) AgregarSubtarea(t *tarea.Tarea, nuevaSubtarea *tarea.Tarea) {
	t.AgregarSubtarea(nuevaSubtarea)
	//t.SetDuracion(t.GetDuracion()) // TODO revisar esto
}

func (l *LinkedList) BuscarTareasPorPalabraClave(palabraClave string) {

	if l.head == nil {
		fmt.Println("La lista está vacía")
		return
	}
	current := l.head
	position := 0
	for current != nil {
		if strings.Contains(current.value.GetNombre(), palabraClave) {
			current.value.MostrarTarea()
			fmt.Println()

		}
		current = current.next
		position++
	}
}

func (l *LinkedList) BuscarTareasPorTag(tag string) {
	indicesEncontrados := l.obtenerPosicionPorTag(tag)

	if len(indicesEncontrados) == 0 {
		return
	}

	for _, v := range indicesEncontrados {
		t := l.SeleccionarTarea(int64(v))
		t.MostrarTarea()
		fmt.Println()

	}

}

func (l *LinkedList) ListarTareas() {
	if l.head == nil {
		fmt.Println("La lista está vacía")
		return
	}
	fmt.Println("Tareas disponibles:")

	position := 0
	current := l.head
	for current != nil {
		fmt.Print("\n-------------------")
		fmt.Print("\n [ ", position, " ] ")
		current.value.MostrarTarea()
		position++
		current = current.next
	}
}

func (l *LinkedList) SeleccionarTarea(position int64) *tarea.Tarea {
	if position < 0 || position >= l.size {
		return nil
	}

	current := l.head
	for current != nil && position > 0 {
		current = current.next
		position--
	}

	return current.value
}

func (l *LinkedList) ModificarNombre(t *tarea.Tarea, nombreNuevo string) {
	t.SetNombre(nombreNuevo)
}
func (l *LinkedList) ModificarDuracion(t *tarea.Tarea, duracionNueva float64) {
	t.SetDuracion(duracionNueva)
}
func (l *LinkedList) ModificarEtiquetas(t *tarea.Tarea, etiquetas []string) {
	t.SetEtiquetas(etiquetas)
}
func (l *LinkedList) SubirPrioridad(t *tarea.Tarea) {
	t.PrioridadAlta()
}
func (l *LinkedList) BajarPrioridad(t *tarea.Tarea) {
	t.PrioridadBaja()
}

func (l *LinkedList) obtenerPosicionPorNombre(nombre string) int {
	if l.head == nil {
		return -1
	}
	current := l.head
	position := 0
	for current != nil {
		if current.value.GetNombre() == nombre {
			return position
		}
		current = current.next
		position++
	}
	return -1
}

func (l *LinkedList) obtenerPosicionPorTag(tagBuscado string) []int {
	var indicesEncontrados []int

	if l.head == nil {
		fmt.Println("La lista de tareas está vacía")
	} else {
		current := l.head
		position := 0
		for current != nil {
			tags := current.value.GetEtiquetas()
			for _, tag := range tags {
				if strings.Contains(tag, tagBuscado) {
					indicesEncontrados = append(indicesEncontrados, position)
				}
			}

			current = current.next
			position++
		}
	}

	return indicesEncontrados
}

func (l *LinkedList) ObtenerTareaActual() *tarea.Tarea {

	if l.head == nil {
		return nil
	}

	current := l.head
	for current != nil {
		if current.value.GetEstado() == tarea.EN_PROGRESO {
			return current.value
		} else {
			current = current.next
		}

	}

	return nil

}

func (l *LinkedList) obtenerTareaRecomendada(duracion float64) *tarea.Tarea {
	if l.head == nil {
		fmt.Println("La lista de tareas está vacía.")
	}

	current := l.head
	for current != nil {
		if current.value.GetDuracion() == duracion {
			return current.value
		} else if current.value.GetDuracion() < duracion {
			if current.value.GetDuracion() > current.next.value.GetDuracion() {
				return current.value
			} else {
				return current.next.value
			}
		} else if current.value.GetDuracion() > duracion {
			fmt.Println("No existe ninguna tarea recomendada para el tiempo ingresado")
		}
		current = current.next
	}

	return nil
}

func (l *LinkedList) IniciarPrimerTareaInterrumpida() {
	if l.head == nil {
		fmt.Println("La lista de tareas está vacía.")
	}

	current := l.head
	for current != nil {
		if current.value.GetEstado() == tarea.INTERRUMPIDA {
			fmt.Println("Comenzando tarea:" + current.value.GetNombre())
			current.value.Realizando()
			return
		}
		current = current.next
	}

	fmt.Println("No hay tareas interrumpidas para comenzar")
}

func (l *LinkedList) InterrumpirTareas() {
	if l.head == nil {
		fmt.Println("La lista de tareas está vacía.")
	}

	current := l.head
	for current != nil {
		if current.value.GetEstado() == tarea.EN_PROGRESO {
			fmt.Println("Interrumpiendo tarea:" + current.value.GetNombre())
			current.value.Posponiendo()
		}
		current = current.next
	}
}

func (l *LinkedList) ObtenerColaDeTareas(tiempo float64) []tarea.Tarea {
	fmt.Println("Tareas que pueden ser realizadas en este tiempo", tiempo,":")
	var cola []tarea.Tarea

	tiempoDisponible := tiempo
	current := l.head
	for current != nil {
		if tiempoDisponible >= current.value.GetDuracion() {
			cola = append(cola, *current.value)
			tiempo = tiempoDisponible - current.value.GetDuracion()
			current.value.MostrarTarea()
		}
		current = current.next
	}

	return cola

}
