
interface ICustomer {
    name: string;
    lastName: string;
    email: string;
    address: {
        description: string;
        city: string;
        country: string;
    };
}


const customer: ICustomer = {
    name: 'Melani',
    lastName: 'Figueroa Santos',
    email: 'mfigueroa@gmail.com',
    address: { description: 'Av. Juan Lecaros 128', city: 'Lima', country: 'Perú' }
};


const { name, address: { country } } = customer;
console.log(name);
console.log(country);


/*=============================================
=                   ARRAYs                   =
=============================================*/
const pokemons: string[] = ['Pikachu', 'Charizard', "tres"];
console.log('Pokemon 2: ', pokemons[1]);

const [, , poke3 = 'No existe el pokemon'] = pokemons
console.log(poke3);

export { }
