import { IAction } from './utils/actions.interfaces'

export class Persona {

    name: string;
    public lastName: string;

    private _age: number;

    protected _fullName: string;


    constructor(name: string, lastName: string, age: number) {
        this.name = name;
        this.lastName = lastName;
        this._age = age;
        this._fullName = `${this.name} ${this.lastName}`;
    }


    // Meotodos
    walk(): void {
        console.log(`${this.name} esta caminando y tiene ${this._age} años`);
    }

    talk(): void {
        console.log(`${this.name} esta hablando: "Naci en el año ${this._yearBirth()}" `);
    }

    // Funcion/Metodo privado
    private _yearBirth() {
        const currentDate = new Date();
        const yearBirth = currentDate.getFullYear() - this._age;
        return yearBirth;
    }
}


const juan = new Persona('Juan', 'Celiz Vega', 31)
console.log('Instancia (Person) "JUAN": ', juan);
juan.walk();
juan.talk();




/*=============================================
=                   Section 02                =
=============================================*/
export class Astronauta extends Persona implements IAction {
    numberMastersDegrees: number;

    constructor(name: string, lastName: string, age: number, numberMastersDegrees: number) {
        // super es una palabra reserva que indica que se tiene que implementar el constructor de la clase extendida
        super(name, lastName, age);
        this.numberMastersDegrees = numberMastersDegrees;
    }

    // Implementacion de la interface
    pilotShip(): void {
        console.log(`${this.name} Tiene permiso para pilotear la nave🚀`);
    }

    canBeAstronaut() {
        return this.numberMastersDegrees > 0;
    }

    getFullName(): string {
        return this._fullName;
    }
}
console.log('*****************************************');
console.log('****DATOS DEL POSTULANTE A ASTRONAUTA****');
console.log('*****************************************');

const astronauta = new Astronauta('Jhon', 'Connor', 45, 0);
console.log('Datos del Astronauta', astronauta);
console.log('Nombres y Apellidos del Astronauta', astronauta.getFullName());
console.log('Puede ser astronauta?', astronauta.canBeAstronaut());

astronauta.pilotShip();

