package tarea

const (
	PENDIENTE = iota
	EN_PROGRESO
	INTERRUMPIDA
	COMPLETADA
)

func (t *Tarea) Realizando() {
	t.estado = EN_PROGRESO
}

func (t *Tarea) Posponiendo() {
	t.estado = INTERRUMPIDA
}

func (t *Tarea) Completando() {
	t.estado = COMPLETADA

	if len(t.GetSubtareas()) > 0 {
		for _, subTarea := range t.GetSubtareas() {
			subTarea.Completando()
		}
	}

}
