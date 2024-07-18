
### Esplicacion
Para entender gráficamente la estructura de un trie (árbol de búsqueda de prefijos) con las palabras "hola", "mundo", "holaquetal" y "holanda", podemos dibujar el árbol de la siguiente manera:
```text

       (raíz)
       /  |  \
      h   m   n
     /    |    \
    o     u     a
   /      |      \
  l       n       n
 / \      |        \
a   q     d         d
|    \    |         |
(✓)  u    o         a
      \   |         |
      e   (✓)       (✓)
       \
        t
         \
          a
           \
            l
             \
              (✓)

```

Notas:
- Cada rama desde la raíz hasta una hoja (marcada con un ✓) representa una palabra completa.
- Los nodos intermedios representan prefijos comunes a varias palabras.


# Uso de rune
En Go, se utiliza rune para trabajar con caracteres individuales de una cadena de texto en lugar de utilizar el tipo byte o string. Esto se debe a varias razones importantes relacionadas con la representación de los caracteres y el soporte para caracteres Unicode.

## Razones para Utilizar rune:
    1. Soporte para Unicode:

rune es un alias para int32 y se utiliza para representar un punto de código Unicode. Esto significa que puede manejar cualquier carácter Unicode, no solo los caracteres ASCII.
Una cadena en Go (string) es una secuencia de bytes, y cada byte es de tipo uint8. Sin embargo, los caracteres Unicode pueden ocupar más de un byte. Usar rune permite tratar cada carácter Unicode completo como una sola entidad, sin preocuparse por la longitud variable de los caracteres en una cadena.


    2. Manipulación de Caracteres:

Cuando trabajas con cadenas de texto, especialmente en idiomas que no son el inglés, es común encontrar caracteres que no se pueden representar con un solo byte (por ejemplo, caracteres acentuados o ideogramas chinos).
rune facilita la manipulación de estos caracteres porque cada rune representa un carácter completo, independientemente de cuántos bytes ocupe en la cadena original.


Ejemplo de Diferencia entre byte y rune:
Considera la cadena "hola" y la cadena "你好" (que significa "hola" en chino).

#### Usando byte:
```go
    palabra := "hola"
    for i := 0; i < len(palabra); i++ {
        fmt.Printf("%c ", palabra[i])  // Salida: h o l a 
    }
    
    palabraChino := "你好"
    for i := 0; i < len(palabraChino); i++ {
        fmt.Printf("%c ", palabraChino[i])  // Salida incorrecta debido a que cada carácter chino ocupa más de un byte
    }
```




#### Usando rune:
```go
    palabra := "hola"
    for _, ch := range palabra {
    fmt.Printf("%c ", ch)  // Salida: h o l a 
    }
    
    palabraChino := "你好"
    for _, ch := range palabraChino {
    fmt.Printf("%c ", ch)  // Salida correcta: 你 好 
    }

```


En el ejemplo de la cadena en chino, si usamos byte, obtenemos una salida incorrecta porque cada carácter chino ocupa más de un byte. Usar rune garantiza que tratemos cada carácter Unicode correctamente, proporcionando la salida correcta.