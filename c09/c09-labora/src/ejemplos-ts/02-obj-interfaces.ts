// Interfaces
interface IPersona {
    name: string;
    lastName: string;
    age: number;
}

// objeto
const persona: IPersona = {
    name: "Juan", lastName: "Perez", age: 23
}
console.log(persona);

// hacemos un extent
interface IStudent extends IPersona {
    carrera: string
    universidad: string
}

const student: IStudent = {
    name: "Juan",
    lastName: "Perez",
    age: 23,
    carrera: "Ing Informatica",
    universidad: "UBA"
}
console.log(student);

export { }



