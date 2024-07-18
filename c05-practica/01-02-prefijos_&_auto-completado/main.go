package main

import "fmt"

//Ejercicio 1: Árbol de búsqueda de prefijos
//
//Diseña e implementa una estructura de datos en Go para un árbol de búsqueda de prefijos (Trie) que almacene un conjunto de palabras. Proporciona funciones para insertar palabras en el árbol y buscar palabras por prefijo.

type Nodo struct {
	hijos        map[rune]*Nodo
	finDePalabra bool
}

type Trie struct {
	raiz *Nodo
}

func NuevaTrie() *Trie {
	return &Trie{
		raiz: &Nodo{hijos: make(map[rune]*Nodo)},
	}
}

// Insertar añade una palabla al trie
func (t *Trie) Insertar(palabra string) {
	nodoActual := t.raiz
	for _, ch := range palabra {
		if _, existe := nodoActual.hijos[ch]; !existe {
			nodoActual.hijos[ch] = &Nodo{hijos: make(map[rune]*Nodo)}
		}
		nodoActual = nodoActual.hijos[ch]
		//nodoActual.finDePalabra = false
	}
	nodoActual.finDePalabra = true
}

// Bucar - verifica si una palabla esta en el trie
func (t *Trie) Buscar(palabra string) bool {
	nodoActual := t.raiz
	for _, ch := range palabra {
		if _, existe := nodoActual.hijos[ch]; !existe {
			return false
		}
		nodoActual = nodoActual.hijos[ch]
	}
	return nodoActual.finDePalabra
}

// ComienzaCon verifica si hay alguna palabra en el trie que comienza con el prefijo dado
func (t *Trie) ComienzaCon(prefijo string) bool {
	nodoActual := t.raiz
	for _, ch := range prefijo {
		if _, existe := nodoActual.hijos[ch]; !existe {
			return false
		}
		nodoActual = nodoActual.hijos[ch]
	}
	return true
}

// PalabrasConPrefijo devuelve todas las palabras en el trie que comienzan con el prefijo dado
func (t *Trie) PalabrasConPrefijo(prefijo string) []string {

	var resultados []string
	nodoActual := t.raiz

	for _, ch := range prefijo {
		if _, existe := nodoActual.hijos[ch]; !existe {
			return resultados
		}
		nodoActual = nodoActual.hijos[ch]
	}

	//funcion aux recogerPalabras
	t.recogerPalabras(nodoActual, prefijo, &resultados)
	return resultados
}

// recogerPalabras es una función auxiliar para recolectar palabras desde un nodo dado
func (t *Trie) recogerPalabras(node *Nodo, prefijo string, resultados *[]string) {
	if node.finDePalabra {
		*resultados = append(*resultados, prefijo)
	}

	for ch, hijo := range node.hijos {
		t.recogerPalabras(hijo, prefijo+string(ch), resultados)
	}
}

func main() {
	//	creamos una instacia del arbol
	trie := NuevaTrie()

	//	datos a insertar
	palabras := []string{"hola", "mundo", "holaquetal", "holanda"}

	//	iteramos para poblar el arbol
	for _, palabra := range palabras {
		trie.Insertar(palabra)
	}

	//fmt.Println(trie.Buscar("hola"))    // true
	//fmt.Println(trie.Buscar("mundo"))   // true
	//fmt.Println(trie.Buscar("Hola"))    // false
	//fmt.Println(trie.Buscar("adios"))   // false
	//fmt.Println(trie.Buscar("holanda")) // true

	fmt.Println(trie.ComienzaCon("hol"))
	fmt.Println(trie.ComienzaCon("mua"))
	fmt.Println(trie.ComienzaCon("mu"))

	//	02 - Auto Completado, simulamos una busqueda de prefijos
	fmt.Println(trie.PalabrasConPrefijo("hol")) // ["hola", "holaquetal", "holanda" ]
	fmt.Println(trie.PalabrasConPrefijo("m"))
}
