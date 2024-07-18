package main

import "fmt"

/*
Implementa un árbol de segmentos en Go que admita dos operaciones: actualizar un valor en un índice específico del arreglo subyacente y calcular la suma de un rango de valores en el arreglo. Proporciona ejemplos de uso para demostrar la eficiencia de tu implementación.
*/

type SegmentTree struct {
	datos []int
	tree  []int
	n     int
}

func NuevaSegmentTree(datos []int) *SegmentTree {
	n := len(datos)
	tree := make([]int, 2*n) // crea un slice de enteros con una longitud y capacidad de 2*n
	st := &SegmentTree{datos, tree, n}
	st.build()
	return st
}

// build construye el árbol de segmentos
func (st *SegmentTree) build() {
	//Inicializamos los nodos
	for i := st.n; i < 2*st.n; i++ {
		st.tree[i] = st.datos[i-st.n]
	}

	//	construimos el arbol calculado los padres
	for i := st.n - 1; i > 0; i-- {
		st.tree[i] = st.tree[i*2] + st.tree[i*2+1]
	}
}

// Update actualiza un valor en un índice específico del arreglo
/*
:::Actualizar los padres
Después de actualizar el valor en la hoja, debemos actualizar todos los nodos padres para reflejar este cambio.
Esto se hace moviéndose hacia arriba en el árbol, dividiendo index por 2 en cada paso (index /= 2).
En cada paso, el nodo padre se actualiza sumando los valores de sus dos hijos:
go
Copy code
st.tree[index] = st.tree[index*2] + st.tree[index*2+1]
Este proceso continúa hasta que se alcanza la raíz del árbol (index > 1).
*/
func (st *SegmentTree) Update(index, value int) {
	index += st.n
	//Actualizar el Valor en el Árbol
	st.tree[index] = value

	// Actualizar los padres
	for index > 1 {
		index /= 2
		st.tree[index] = st.tree[index*2] + st.tree[index*2+1]
	}
}

// Query calcula la suma de un rango de valores en el arreglo
func (st *SegmentTree) Query(left, right int) int {
	left += st.n
	right += st.n
	sum := 0
	for left < right {
		if left%2 == 1 {
			sum += st.tree[left]
			left++
		}
		if right%2 == 1 {
			right--
			sum += st.tree[right]
		}
		left /= 2
		right /= 2
	}
	return sum
}

func main() {
	datos := []int{1, 3, 5, 7, 9, 11}

	st := NuevaSegmentTree(datos)

	fmt.Println("Suma de rango (1, 5):", st.Query(1, 5)) // 3 + 5 + 7 + 9 = 24

	st.Update(1, 10) // Cambia datos[1] de 3 a 10

	fmt.Println("Suma de rango (1, 5):", st.Query(1, 5)) // 10 + 5 + 7 + 9 = 31
}

/*
Construcción del Árbol

- El arreglo original es [1, 3, 5, 7, 9, 11].

- Se ha creado un slice tree con el doble del tamaño del arreglo original para acomodar tanto los nodos hoja como los nodos internos.

- Los nodos hoja se llenan con los valores del arreglo original en la segunda mitad del slice tree.

Índices en tree:      0  1  2  3  4  5  6  7  8  9 10 11
Valores iniciales: [ 0, 0, 0, 0, 0, 0, 1, 3, 5, 7, 9, 11]


--> Cálculo de los Nodos Internos
Vamos a llenar los nodos internos comenzando desde el índice n-1 hasta el índice 1. En cada paso, un nodo interno se calcula como la suma de sus dos nodos hijos. La fórmula general es:

formula: tree[i] = tree[2*i] + tree[2*i + 1]
Cálculo del Nodo en Índice 5:
	tree[5] = tree[10] + tree[11]
	tree[10] = 9, tree[11] = 11
	tree[5] = 9 + 11 = 20

	Índices en tree:      	0  1  2  3  4  5  6  7  8  9  10  11
	Valores actualizados: [ 0, 0, 0, 0, 0, 20, 1, 3, 5, 7, 9, 11]

Cálculo del Nodo en Índice 4:
	tree[4] = tree[8] + tree[9]
	tree[8] = 5, tree[9] = 7
	tree[4] = 5 + 7 = 12

	Índices en tree:      	0  1  2  3  4  5  6  7  8  9  10  11
	Valores actualizados: [ 0, 0, 0, 0, 12, 20, 1, 3, 5, 7, 9, 11]



Cálculo del Nodo en Índice 1:

tree[1] = tree[2] + tree[3]
tree[2] = 32, tree[3] = 4
tree[1] = 32 + 4 = 36

.... terminamos las iteraciones
Índices en tree:      	0  1   2   3   4   5  6  7  8  9  10  11
Valores actualizados: [ 0, 36, 32, 4, 12, 20, 1, 3, 5, 7, 9, 11]



Representación Gráfica del Árbol de Segmentos
               [1]36
            /          \
       [2]32            [3]4
     /      \           /     \
 [4]12      [5]20     [6]1     [7]3
  /  \       /   \
[8]5 [9]7  [10]9 [11]11

*/
