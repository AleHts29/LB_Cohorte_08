package main

import (
	"errors"
	"fmt"
)

type Pila struct {
	elementos []interface{}
}

// Metodo push: inserta elementos en el tope de la pila
func (p *Pila) Push(e interface{}) {
	p.elementos = append(p.elementos, e)
}

// Metodo push: inserta elementos en el tope de la pila
func (p *Pila) Size() int {
	return len(p.elementos)
}

func (p *Pila) IsEmphy() bool {
	return len(p.elementos) == 0
}

// Metodo Top: retorna el elemeto del tope
func (p *Pila) Top() (interface{}, error) {
	//	Validamos si la pila esta vacia
	if p.IsEmphy() {
		return nil, errors.New("error: La pila esta vacia")
	}

	return p.elementos[len(p.elementos)-1], nil
}

func (p *Pila) Pop() (interface{}, error) {
	//	Validamos si la pila esta vacia
	if p.IsEmphy() {
		return nil, errors.New("error: La pila esta vacia")
	}

	elemento := p.elementos[len(p.elementos)-1]

	p.elementos = p.elementos[:len(p.elementos)-1]

	return elemento, nil
}

func main() {

	//	Creamos la pila
	pila := &Pila{}

	//	Agregamos elementos a la pila
	//pila.Push(10)
	//pila.Push(20)
	//pila.Push(50) // <---

	fmt.Println("Tamaño de la pila: ", pila.Size())

	topElement, err := pila.Top()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("Element del tope:", topElement)

	//	Usando POP
	popElement, _ := pila.Pop()
	fmt.Println("Elemento eliminado:", popElement)

	fmt.Println("Tamaño de la pila: ", pila.Size())

	pila.Push(40)

	fmt.Println("Tamaño de la pila: ", pila.Size())

	fmt.Println("La pila esta vacia?: ", pila.IsEmphy())
}
