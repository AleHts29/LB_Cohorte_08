### AWS

- 01 _Crear la instacia EC2 desde AWS
- 02 _Generar la clave .pem
- 03 _Conectarce mediante aws cli

---
### Actualizacion de instacia 

- sudo apt update

### Desplegar app
```bash
#Opcion-01 Instalar golang en la EC2 y clonar repo
sudo apt install golang-go


#Opcion-02 Generar binario
GOOS=linux GOARCH=amd64 go build -o tu_aplicacion
```


### subir binario
```bash
scp -i "aws/labora_c08_aws.pem" "Cohorte_08/c07-servidores_http/server-http" ubuntu@ec2-18-222-58-105.us-east-2.compute.amazonaws.com:/home/ubuntu/
```

### dar permisos al binario despues de subirlo
```bash
chmod +x tu_aplicacion
```