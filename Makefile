# Define las imágenes de Docker para los servicios
SERVICE_A_IMAGE=service-a
SERVICE_B_IMAGE=service-b
MONGO_IMAGE=mongo:6.0
RABBITMQ_IMAGE=rabbitmq:management

# Comandos básicos de Docker
build:
	docker-compose build

up:
	docker-compose up -d

down:
	docker-compose down

logs:
	docker-compose logs -f

restart:
	docker-compose restart

clean:
	docker-compose down --volumes --remove-orphans

# Para iniciar el proyecto con todos los servicios
start: build up

# Para detener todo y limpiar recursos
stop: down clean

# Para reiniciar el proyecto (reconstruir y levantar servicios)
restart-all: clean start

# Ejecutar el servicio A y B con Docker Compose
docker-compose-exec-a:
	docker-compose exec service-a bash

docker-compose-exec-b:
	docker-compose exec service-b bash

# Ejecutar pruebas unitarias de Service A
test-a:
	cd service-a/application/user && go test ./ -v



