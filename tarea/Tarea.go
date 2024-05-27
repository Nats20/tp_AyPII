package tarea

import (
	"fmt"
)

type Tarea struct {
	nombre    string
	duracion  float64
	subtareas []*Tarea
	prioridad int
	estado    int64
	tags      []string
}

func CrearTarea(nombre string, duracion float64, tag []string) *Tarea {
	return &Tarea{nombre: nombre, duracion: modulo(duracion), prioridad: SIN_PRIORIDAD, estado: PENDIENTE, tags: tag}
}

func (t *Tarea) MostrarTarea() {
	fmt.Println("Nombre:", t.nombre, " (", t.duracion, "hs ), Prioridad: "+t.mostrarPrioridad()+", Estado: "+t.mostrarEstado()+", Tags:"+ t.mostrarTags(),"\n")
	if len(t.subtareas) > 0 {
		fmt.Println("\nSubtareas de la tarea " + t.nombre + ":")
		for _, v := range t.subtareas {
			v.MostrarTarea()
		}
	}
}

// setters

func (t *Tarea) AgregarSubtarea(nuevaTarea *Tarea) {
	t.subtareas = append(t.subtareas, nuevaTarea)
}

func (t *Tarea) EliminarSubtarea(indice int) {

}

func (t *Tarea) SetNombre(nuevoNombre string) {
	t.nombre = nuevoNombre
}

func (t *Tarea) SetDuracion(nuevaDuracion float64) {
	if len(t.GetSubtareas()) > 0 {
		acum := float64(0)
		for _, subTarea := range t.GetSubtareas() {
			acum += subTarea.duracion
		}
		t.duracion = modulo(nuevaDuracion) + acum

	} else {
		t.duracion = modulo(nuevaDuracion)
	}
}

func (t *Tarea) SetEstado(nuevoEstado int64) {
	t.estado = nuevoEstado
}

func (t *Tarea) SetEtiquetas(newTags []string) {
	t.tags = newTags
}

func (t *Tarea) SetSubtareas(subtareas []*Tarea) {
	t.subtareas = subtareas
}

func (t *Tarea) PrioridadAlta() {
	t.prioridad = ALTA
}
func (t *Tarea) PrioridadBaja() {
	t.prioridad = BAJA
}

// getters

func (t *Tarea) GetSubtareas() []*Tarea {
	return t.subtareas
}

func (t *Tarea) GetNombre() string {
	return t.nombre
}

func (t *Tarea) GetPrioridad() int {
	return t.prioridad
}

func (t *Tarea) GetEstado() int64 {
	return t.estado
}

func (t *Tarea) GetEtiquetas() []string {
	return t.tags
}
func (t *Tarea) GetDuracion() float64 {
	return t.duracion
}

func modulo(v float64) float64 {
	if v < 0 {
		return v / -1
	} else {
		return v
	}
}

func (t Tarea) mostrarEstado() string {
	switch t.estado {
	case PENDIENTE:
		return "Pendiente"
	case EN_PROGRESO:
		return "En progreso"
	case INTERRUMPIDA:
		return "Interrumpida"
	case COMPLETADA:
		return "Completada"
	default:
		return "Pendiente"
	}
}

func (t Tarea) mostrarPrioridad() string {
	switch t.prioridad {
	case SIN_PRIORIDAD:
		return "Sin prioridad"
	case BAJA:
		return "Baja"
	case ALTA:
		return "Alta"
	default:
		return "Sin prioridad"
	}
}

func (t *Tarea)mostrarTags()string{
	r:=""
	for i := 0; i < len(t.tags); i++ {
		r+=" #"
		r+=t.tags[i]
	}
	return r
}