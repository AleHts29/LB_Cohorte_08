-- create table Empleados(
-- id_empleado INT PRIMARY KEY AUTO_INCREMENT,
-- nombre VARCHAR(100),
-- apellido VARCHAR(100),
-- fecha_nacimiento DATE,
-- direccion VARCHAR(255),
-- telefono VARCHAR(15)

-- );



-- INSERT INTO Empleados (nombre, apellido, fecha_nacimiento, direccion, telefono)
-- VALUES ('Juan', 'Perez', '1990-04-23', 'Av 123, BsAs', '123-456-45')


-- INSERT INTO Empleados (nombre, apellido, fecha_nacimiento, direccion, telefono) 
-- VALUES 
-- ('María', 'González', '1990-08-15', 'Avenida Siempre Viva 456', '555-5678'),
-- ('Carlos', 'Ramírez', '1975-01-30', 'Boulevard de los Sueños 789', '555-8765'),
-- ('Luisa', 'Fernández', '1982-11-12', 'Calle del Sol 101', '555-3456'),
-- ('Ana', 'Martínez', '1995-07-20', 'Avenida del Mar 202', '555-2345'),
-- ('Pedro', 'López', '1988-02-10', 'Calle Luna 303', '555-6789'),
-- ('Laura', 'Torres', '1992-06-18', 'Avenida Estrella 404', '555-7890'),
-- ('Miguel', 'Sánchez', '1979-11-25', 'Boulevard Amanecer 505', '555-8901'),
-- ('Sofía', 'Jiménez', '1987-09-05', 'Calle Rayo 606', '555-9012'),
-- ('Andrés', 'Hernández', '1983-03-14', 'Avenida Aurora 707', '555-0123');


-- SELECT * FROM Empleados;

-- SELECT id_empleado, nombre, telefono FROM Empleados;


-- CREATE TABLE Departamentos (
--     id_departamento INT PRIMARY KEY AUTO_INCREMENT,
--     nombre_departamento VARCHAR(100)
-- );


-- ALTER TABLE Empleados ADD id_departamento INT;

-- ALTER TABLE Empleados
-- ADD CONSTRAINT fk_departamento
-- FOREIGN KEY (id_departamento) REFERENCES Departamentos(id_departamento);

INSERT INTO Departamentos (nombre_departamento)
VALUES ('Recursos Humanos'), ('Finanzas'), ('Desarrollo'), ('Ventas'), ('Marketing');


