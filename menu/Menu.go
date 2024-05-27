package Menu

import (
	"TP--nosehizoconchatgpt/lista"
	"TP--nosehizoconchatgpt/tarea"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Menu struct {
	lista *lista.LinkedList
}

func CrearMenu() *Menu {
	list := lista.NewLinkedList()
	return &Menu{lista: list}
}

func (m *Menu) EncenderMenu() {
	fmt.Println("\n   MENU  ")
	m.lista.ListarTareas()

	if m.lista.ObtenerTareaActual()==nil{
		fmt.Println("______________No hay una tarea en progreso______________")
	}else{
		fmt.Println("______________Tarea Actual______________")
		fmt.Print("> ")
		m.lista.ObtenerTareaActual().MostrarTarea()
	}
	
	fmt.Print(
		`	
[ a ] Editar tarea
[ b ] Nueva tarea
[ c ] Tarea actual
[ d ] Preparar cola de tareas
[ e ] Buscar tarea 
[ f ] Reordenar lista de tareas
[ g ] Eliminar una tarea de la lista
[ x ] Cerrar Menu
`)

	eleccion := leerString()
	switch eleccion {
	case "a":
		if m.lista.Size() == 0 {
			CallClear()
			fmt.Println("No hay tareas para editar")
			m.EncenderMenu()
		} else {
			CallClear()
			m.abrirEdicionDeTareas()
			m.EncenderMenu()
		}
	case "b": // listo
		CallClear()
		m.nuevaTarea()
	case "c": // listo
		CallClear()
		m.tareaActual()
	case "d": // listo
		if m.lista.Size() == 0 {
			CallClear()
			fmt.Println("No hay tareas para sugerir")
			m.EncenderMenu()
		} else {
			CallClear()
			m.prepararColaDeTareas()
			m.EncenderMenu()
		}
	case "e": // listo
		if m.lista.Size() == 0 {
			CallClear()
			fmt.Println("NO HAY TAREAS EN LA LISTA")
			m.EncenderMenu()
		} else {
			CallClear()
			m.buscarTareas()
			m.EncenderMenu()
		}
	case "f": // listo
		CallClear()
		m.reordenarLista()
		m.EncenderMenu()
	case "g": // listo
		CallClear()
		if m.lista.Size() == 0 {
			CallClear()
			fmt.Println("NO HAY TAREAS PARA ELIMINAR")
			m.EncenderMenu()
		} else {
			CallClear()
			m.eliminarUnaTarea()
			m.EncenderMenu()
		}
	case "x":
		CallClear()
		return
	default:
		CallClear()
		m.EncenderMenu()
	}
}

func (m *Menu) buscarTareas() {
	var tareasEncontradas []tarea.Tarea
	var err error
	var opt string
	repeat := true

	for repeat {
		fmt.Println("\nElija el método de consulta")
		fmt.Println("Para buscar por palabra clave, ingrese 1.")
		fmt.Println("Para buscar por tag, ingrese 2.")
		fmt.Println("Para salir ingrese 0.")
		opt = leerString()
		fmt.Println("Opcion ingresada", opt)
		switch opt {
		case "1":
			fmt.Println("Ingresar palabra clave de la tarea")
			nombreBusqueda := leerString()
			m.lista.BuscarTareasPorPalabraClave(nombreBusqueda)
		case "2":
			fmt.Println("Ingresar tag de la tarea")
			tagBusqueda := leerString()
			m.lista.BuscarTareasPorTag(tagBusqueda)
		case "0":
			repeat = false
			CallClear()
			fmt.Println("Saliendo de búsqueda\n")
		default:
			fmt.Println("Opción incorrecta\n")
		}
	}

	if opt == "1" || opt == "2" {
		if err == nil {
			fmt.Println("\nTareas encontradas:")
			for _, t := range tareasEncontradas {
				t.MostrarTarea()
			}
		} else {
			fmt.Println(err)
		}
	}
}

func (m *Menu) eliminarUnaTarea() {
	repeat := true
	for repeat {
		CallClear()
		m.lista.ListarTareas()
		fmt.Println("\nIngresar posicion de la tarea. Para salir, ingrese 'x'.")
		posicion := leerString()

		if posicion == "x" {
			repeat = false
			CallClear()
			fmt.Println("Saliendo")
		} else {
			p, _ := strconv.ParseInt(posicion, 0, 0)
			t := m.lista.SeleccionarTarea(p)

			if t == nil {
				fmt.Println("la tarea NO existe. Seleccione una posición válida")
			} else {
				m.lista.EliminarTarea(p)
			}
		}
	}
}

func (m *Menu) abrirEdicionDeTareas() {
	repeat := true
	for repeat {
		m.lista.ListarTareas()
		fmt.Println("\nIngresar posicion de la tarea. Para salir, ingrese 'x'.")
		posicion := leerString()

		if posicion == "x" {
			repeat = false
			CallClear()
			fmt.Println("Saliendo de la edición de tarea")
		} else {
			p, _ := strconv.ParseInt(posicion, 0, 0)
			t := m.lista.SeleccionarTarea(p)

			if t == nil {
				fmt.Println("la tarea NO existe. Seleccione una posición válida")
			} else {
				m.editarTarea(t)
				//fmt.Println(t)
				//t.SetDuracion(t.GetDuracion()) // TODO revisar esto
				//fmt.Println(t)

			}
		}
	}
}

func (m *Menu) editarTarea(t *tarea.Tarea) {
	CallClear()
	fmt.Println("tarea a editar:")
	t.MostrarTarea()

	fmt.Println("\nElegir campo a editar")
	fmt.Print(
		`
[ a ] editar nombre
[ b ] editar duracion
[ c ] subir prioridad
[ d ] bajar prioridad
[ e ] modificar etiquetas
[ f ] comenzar tarea
[ g ] nueva subtarea
[ h ] editar subtareas
[ i ] eliminar una subtarea
[ x ] volver
				`)

	opcion := leerString()
	switch opcion {
	case "a": // listo
		fmt.Println("Ingresar nuevo nombre")
		nombreNuevo := leerString()
		m.lista.ModificarNombre(t, nombreNuevo)
		CallClear()
	case "b": // listo
		fmt.Println("Ingresar nueva duracion")
		duracionNueva := leerString()
		duracion, _ := strconv.ParseFloat(duracionNueva, 8)
		m.lista.ModificarDuracion(t, duracion)
		CallClear()
	case "c": // listo
		m.lista.SubirPrioridad(t)
		CallClear()
	case "d": // listo
		m.lista.BajarPrioridad(t)
		CallClear()
	case "e": // listo
		fmt.Println("Ingresar nuevas tags separadas por comas")
		tags := leerString()
		nuevasTags := strings.Split(tags, ",")
		m.lista.ModificarEtiquetas(t, nuevasTags)
		CallClear()
	case "f": // listo
		m.lista.InterrumpirTareas()
		t.Realizando()
		CallClear()
	case "g":
		m.lista.AgregarSubtarea(t, m.nuevaSubtarea())
		CallClear()
	case "h":
		subtareas := t.GetSubtareas()
		if len(subtareas) == 0 {
			fmt.Println("No hay subtareas para editar!")
		} else {
			for i, st := range subtareas {
				fmt.Println("Subtarea Nro ", i)
				st.MostrarTarea()
			}

			repeat := true
			for repeat {
				fmt.Println("\nIngresar número de subtarea a editar")
				posicion := leerString()
				p, _ := strconv.ParseInt(posicion, 0, 0)

				if p < 0 || int(p) >= len(subtareas) {
					fmt.Println("Posición inválida!")
				} else {
					repeat = false
					m.editarTarea(subtareas[p])
				}
			}
		}
		CallClear()
	case "i":
		subtareas := t.GetSubtareas()
		if len(subtareas) == 0 {
			fmt.Println("No hay subtareas para eliminar!")
		} else {
			for i, st := range subtareas {
				fmt.Println("Subtarea Nro ", i)
				st.MostrarTarea()
			}

			repeat := true
			for repeat {
				fmt.Println("\nIngresar número de subtarea a eliminar:")
				posicion := leerString()
				p, _ := strconv.ParseInt(posicion, 0, 0)

				if p < 0 || int(p) >= len(subtareas) {
					fmt.Println("Posición inválida!")
				} else {
					repeat = false
					nuevasSubtareas := append(subtareas[:p], subtareas[(p+1):]...)
					t.SetSubtareas(nuevasSubtareas)
				}
			}
		}
		CallClear()
	case "x":
		CallClear()
	default:
		CallClear()
	}

	t.SetDuracion(t.GetDuracion()) // TODO revisar esto
	m.lista.Ordenar(m.lista.GetOrderBy())
}

func (m *Menu) nuevaSubtarea() *tarea.Tarea {
	fmt.Print("Ingresar Nombre: ")
	name := leerString()

	fmt.Print("Ingresar Duracion (mayor o igual a 0): ")
	inputTime := leerString()
	parsedTime, _ := strconv.ParseFloat(inputTime, 8)

	fmt.Println(`Ingresar tags separadas por comas: `)
	tags := leerString()
	slicedTags := strings.Split(tags, ",")

	t := tarea.CrearTarea(name, parsedTime, slicedTags)
	CallClear()
	return t
}

// 1. Alta
func (m *Menu) nuevaTarea() {
	fmt.Print("Ingresar Nombre: ")
	name := leerString()

	fmt.Print("Ingresar Duracion (mayor o igual a 0): ")
	inputTime := leerString()
	parsedTime, _ := strconv.ParseFloat(inputTime, 8)

	fmt.Println(`Ingresar tags separadas por comas: `)
	tags := leerString()
	slicedTags := strings.Split(tags, ",")

	m.lista.Append(tarea.CrearTarea(name, parsedTime, slicedTags))
	m.lista.Ordenar(m.lista.GetOrderBy())

	CallClear()
	m.EncenderMenu()
}

func (m *Menu) tareaActual() {
	CallClear()
	tareaActual := m.lista.ObtenerTareaActual()

	if tareaActual == nil {
		CallClear()
		fmt.Println("---------------No hay ninguna tarea en progreso!---------------")
		m.EncenderMenu()
	} else {
		fmt.Println("Tarea actual:")
		tareaActual.MostrarTarea()

		fmt.Print(`
Tarea actual:
[ 1 ] completada
[ 2 ] abandonar
[ x ] volver a menu
`)
		reader := bufio.NewReader(os.Stdin)
		entrada, _ := reader.ReadString('\n')
		eleccion := strings.TrimRight(entrada, "\r\n")

		switch eleccion {
		case "1":
			tareaActual.Completando()
			m.lista.IniciarPrimerTareaInterrumpida()
			CallClear()
			m.tareaActual()
		case "2": // listo
			tareaActual.Posponiendo()
			m.lista.IniciarPrimerTareaInterrumpida()
			CallClear()
			m.tareaActual()
		case "x":
			CallClear()
			m.EncenderMenu()
		}
	}

}

// 4. Consulta
func (m *Menu) prepararColaDeTareas() {
	fmt.Print(`
	ingresar tiempo disponible para ofrecer un listado de tareas:
	`)
	reader := bufio.NewReader(os.Stdin)
	entrada, _ := reader.ReadString('\n')
	eleccion := strings.TrimRight(entrada, "\r\n")
	parsedEleccion, _ := strconv.ParseFloat(eleccion, 8)
	m.lista.ObtenerColaDeTareas(parsedEleccion)
}

func (m *Menu) reordenarLista() {
	fmt.Print(
		`
	Reordenar tareas por:
	[ 1 ] Prioridad
	[ 2 ] Duracion
	[ 3 ] Cantidad de subtareas
	`)
	reader := bufio.NewReader(os.Stdin)
	entrada, _ := reader.ReadString('\n')
	eleccion := strings.TrimRight(entrada, "\r\n")

	switch eleccion {
	case "1": // terminar
		m.lista.Ordenar(lista.PRIORIDAD)
	case "2": // listo
		m.lista.Ordenar(lista.DURACION)
	case "3": // listo
		m.lista.Ordenar(lista.CANT_SUBTAREAS)
	default:
		fmt.Print("ingresar opción valida")
		m.reordenarLista()
	}
	CallClear()
	m.EncenderMenu()
}

func leerString() string {
	reader := bufio.NewReader(os.Stdin)
	entrada, _ := reader.ReadString('\n')
	dato := strings.TrimRight(entrada, "\r\n")
	return dato
}
