function showMessage() {
    console.log("Hola Coders");
}
showMessage()


const saludo = (name: string): void => {
    console.log(`Hola ${name}`);
}
saludo("Juan")


const sum = (a: number, b: number): number => {
    return a + b;
}
console.log('Suma 01:', sum(2, 3));


// Esto es una mala practica❌
const showStudentData = (
    name: string,
    lastName: string,
    age: number,
    career: string,
    university: string,
    cycle: number
) => {
    console.log('****Los datos del Estudiante son los siguientes:***');
    console.log('Nombre: ', name);
    console.log('Apellidos: ', lastName);
    console.log('Edad: ', age);
    console.log('Carrera: ', career);
    console.log('Universidad: ', university);
    console.log('Ciclo actual: ', cycle);
}

showStudentData('Ernesto', 'Torres Ugarte', 25, 'Bellas Artes', 'Universidad Catolica del Peru', 5);


// ------ Forma correcta

interface IPersona {
    name: string;
    lastName: string;
    age: number;
}

interface IStudent extends IPersona {
    career: string;
    university: string;
    cycle?: number; // con el operador '?' indicamos que un atributo del objecto es opcional
}


const showStudentData2 = (data: IStudent) => {
    console.log('****Los datos del Estudiante son los siguientes:***');
    console.log('Nombre: ', data.name);
    console.log('Apellidos: ', data.lastName);
    console.log('Edad: ', data.age);
    console.log('Carrera: ', data.career);
    console.log('Universidad: ', data.university);
    console.log('Ciclo actual: ', data.cycle);
}


const IStudent: IStudent = {
    name: 'Percy',
    lastName: 'Garcia Mori',
    age: 24,
    career: 'Ingienería Informática',
    university: 'Universidad Mayor de San Marcos',
    cycle: 5
}

showStudentData2(IStudent)
