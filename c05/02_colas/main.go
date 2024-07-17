package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	elemnets []interface{}
}

func (q *Queue) IsEmphy() bool {
	return len(q.elemnets) == 0
}

// Enqueue: Inserta el elemento en el rabo de la cola
func (q *Queue) Enqueue(e interface{}) {
	q.elemnets = append(q.elemnets, e)
}

// Dequeue: Elimina el elemento del frente de la cola y lo retorna
func (q *Queue) Dequeue() (interface{}, error) {
	if q.IsEmphy() {
		return nil, errors.New("la cola está vacía")
	}

	element := q.elemnets[0]
	q.elemnets = q.elemnets[1:]

	return element, nil
}

// Front: Retorna el elemento del frente de la cola
func (q *Queue) Front() (interface{}, error) {
	if q.IsEmphy() {
		return nil, errors.New("la cola está vacía")
	}
	return q.elemnets[0], nil
}

// Size: Retorna la cantidad de elementos de la cola
func (q *Queue) Size() int {
	return len(q.elemnets)
}

func main() {

	queue := Queue{}

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	fmt.Println("Tamaño de la cola:", queue.Size())

	front, _ := queue.Front()
	fmt.Println("Elemento del frente:", front)

	dequeued, _ := queue.Dequeue()
	fmt.Println("Elemento desencolado:", dequeued)

	fmt.Println("Tamaño de la cola después de desencolar:", queue.Size())

	isEmpty := queue.IsEmphy()
	fmt.Println("¿La cola está vacía?", isEmpty)

}
