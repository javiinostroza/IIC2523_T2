# Construir y ejecutar una imagen como contenedor

se debe crear un documento Dockerfile para indicar como se construira el contenedor,
en este deben estar los comandos para ejecutar el contenedor

con el siguiente comando se iniciara el contenedor: docker build -t getting-started .
Para correr el programa del contenedor se debe utilizar este comando: 

docker run -dp 3000:3000 nombre


# Compartir im ́agenes usando Docker Hub

Para compartir un contenedor se debe iniciar secion en dockerhub con el comando: docker login -u nombre_usuario

Para subir el contenedor se utiliza este comando:

docker tag nombre nombre_usuario/nombre

seguido este comando para pushear todo:

docker push nombre_usuario/nombre


# Desplegar aplicaciones Docker utilizando m ́ultiples contenedores con una base de datos

se debe iniciar el contenedor como anteriormente mencionado, seguido por este comando para usar mysql:

docker run -d \
     --network todo-app --network-alias mysql \
     -v todo-mysql-data:/var/lib/mysql \
     -e MYSQL_ROOT_PASSWORD=contraseña \
     -e MYSQL_DATABASE=todos \
     mysql:5.7


uno se conecta a mysql con este comando:

docker exec -it <mysql-container-id> mysql -u root -p

se cambian los privilegios de las bases de datos a usar:

 ALTER USER 'root' IDENTIFIED WITH mysql_native_password BY 'secret';

 flush privileges;

 se crean las variable de ambiente necesarias:

  docker run -dp 3000:3000 \
   -w /app -v "$(pwd):/app" \
   --network todo-app \
   -e MYSQL_HOST=mysql \
   -e MYSQL_USER=root \
   -e MYSQL_PASSWORD=secret \
   -e MYSQL_DB=todos \
   node:12-alpine \
   sh -c "yarn install && yarn run dev"

y los contenedores estan conectados.

# Ejecutar aplicaciones utilizando Docker Compose

Primero se debe tener instalado docker compose

se debe crear un archivo con toda la información necesaria para el prgrama:
docker-compose.yml

se define los servicios de la app y la base de datos:

 docker run -dp 3000:3000 \
  -w /app -v "$(pwd):/app" \
  --network todo-app \
  -e MYSQL_HOST=mysql \
  -e MYSQL_USER=root \
  -e MYSQL_PASSWORD=contraseña \
  -e MYSQL_DB=todos \
  node:12-alpine \
  sh -c "yarn install && yarn run dev"




 docker run -d \
  --network todo-app --network-alias mysql \
  -v todo-mysql-data:/var/lib/mysql \
  -e MYSQL_ROOT_PASSWORD=contraseña \
  -e MYSQL_DATABASE=todos \
  mysql:5.7

se arma el compose:

docker compose up -d