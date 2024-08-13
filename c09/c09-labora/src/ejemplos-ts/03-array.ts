// 1. Declaración de un array de números
let numeros: number[] = [1, 2, 3, 4, 5]
console.log(numeros);

// 2. Declaración de un array de strings
let frutas: string[] = ['manzana', 'banana', 'naranja'];
console.log(frutas);

// 3. Declaración de un array utilizando la sintaxis genérica
let nombres: Array<string> = ['Juan', 'Pedro', 'Maria'];
console.log(nombres);

// Mixto
let mezcla: (number | string | boolean)[] = [1, 'dos', true, 3];
console.log(mezcla);


// 6. Agregar un elemento al final del array
frutas.push('pera');  // ['manzana', 'banana', 'naranja', 'pera']


// 7. Eliminar el último elemento del array
frutas.pop();  // ['manzana', 'banana', 'naranja']


// 8. Iteracion
for (let fruta of frutas) {
    console.log(fruta);
}


// map: para transformar elementos
let numerosDoblados: number[] = numeros.map(num => num * 2);  // [2, 4, 6, 8, 10]
console.log(numerosDoblados);

// filter: para filtrar elementos
let numerosPares: number[] = numeros.filter(num => num % 2 === 0);  // [2, 4]
console.log(numerosPares);

// reduce: para reducir los elementos a un solo valor
let sumaTotal: number = numeros.reduce((total, num) => total + num, 0);  // 15
console.log(sumaTotal);



/*=============================================
=                   Section  02               =
=============================================*/
interface IPet {
    name: string;
    age: string;
    species: string;
    isFemale: boolean;
    vaccines: string[]; // un objecto tambien puede tener atributos de tipo "array" u "objecto"😎 
}

const petHomero: IPet = {
    name: 'Homero',
    age: '15 años',
    species: 'Perro',
    isFemale: false,
    vaccines: ['Rabia']
}


const petKitty: IPet = {
    name: 'Kitty',
    age: '5 años',
    species: 'Gato',
    isFemale: true,
    vaccines: ['Rabia', 'Moquillo']
}

// Copntruir un array de obj
const pets: IPet[] = [petHomero, petKitty]
console.log(pets);


const petFilter = pets.filter((pet) => {
    return pet.isFemale
})
console.log(petFilter);


export { }
