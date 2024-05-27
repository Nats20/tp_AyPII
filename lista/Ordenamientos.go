package lista

const (
	DURACION = iota
	PRIORIDAD
	CANT_SUBTAREAS
)

func (l *LinkedList) Ordenar(orderBy int) {
	position := 1
	for position <= (int(l.size) - 1) {
		if orderBy == DURACION {
			l.OrdenarPorDuracion()
			l.SetOrderBy(DURACION)
		} else if orderBy == PRIORIDAD {
			l.OrdenarPorPrioridad()
			l.SetOrderBy(PRIORIDAD)
		} else if orderBy == CANT_SUBTAREAS {
			l.OrdenarPorCantSubtareas()
			l.SetOrderBy(CANT_SUBTAREAS)
		}
		position++
	}
}

func (l *LinkedList) OrdenarPorPrioridad() {
	current := l.head
	for current != nil {
		if current.next != nil {
			if current.value.GetPrioridad() < current.next.value.GetPrioridad() {
				current.value, current.next.value = current.next.value, current.value
				return
			}
			current = current.next
		} else {
			return
		}
	}
}

func (l *LinkedList) OrdenarPorDuracion() {
	current := l.head
	for current != nil {
		if current.next != nil {
			if current.value.GetDuracion() > current.next.value.GetDuracion() {
				current.value, current.next.value = current.next.value, current.value
				return
			}
			current = current.next
		} else {
			return
		}
	}
}

func (l *LinkedList) OrdenarPorCantSubtareas() {
	current := l.head
	for current != nil {
		if current.next != nil {
			if len(current.value.GetSubtareas()) < len(current.next.value.GetSubtareas()) {
				current.value, current.next.value = current.next.value, current.value
				return
			}
			current = current.next
		} else {
			return
		}
	}
}
