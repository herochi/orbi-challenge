Este proyecto es una prueba técnica que consiste en un sistema con dos servicios:

- **Service A**: Un servicio REST desarrollado en Go para gestionar usuarios, basado en una arquitectura limpia.
- **Service B**: Un servicio gRPC en Go que maneja la notificación de los usuarios y se comunica con el servicio A a través de gRPC y RabbitMQ, implementado con una arquitectura hexagonal.

## Requisitos previos

Para levantar el proyecto, necesitas tener instalados los siguientes servicios:

- Docker y Docker Compose
- Go (si quieres modificar el código fuente directamente)

## Levantar el proyecto

Para levantar todo el sistema con sus contenedores de Docker, simplemente ejecuta el siguiente comando:

```bash
make start
```
Esto construirá y levantará todos los servicios necesarios (service-a, service-b, MongoDB y RabbitMQ) en sus respectivos contenedores Docker.


## Detener el proyecto
```bash
make stop
```
## Documentación de Endpoints de Service A

Los endpoints de **Service A** están documentados utilizando **Swagger** y se encuentran en el archivo `swagger.yaml`, ubicado en la carpeta `documents` del proyecto. Puedes usar **Swagger UI** o cualquier otra herramienta compatible con el formato Swagger para explorar los endpoints de **Service A**.

### Endpoints disponibles en Service A

- **POST** `/orbi-api/v1/a/users/`:  
  Crea un nuevo usuario. Este endpoint no solo crea un usuario en **Service A**, sino que también realiza una petición gRPC a **Service B** enviando el ID del usuario creado.

- **GET** `/orbi-api/v1/a/users/{id}`:  
  Obtiene los detalles de un usuario a partir de su ID.

- **PATCH** `/orbi-api/v1/a/users/{id}`:  
  Actualiza la información de un usuario. Al realizar la actualización, **Service A** envía una notificación a **Service B** a través de RabbitMQ.

**Puerto de Service A:**  
Service A corre en el puerto **8080**.

## Contratos de Comunicación gRPC de Service B

Los contratos de comunicación gRPC de **Service B** se encuentran en el archivo `user.proto`, ubicado en la carpeta `user` dentro del proyecto de **Service B**. Puedes revisar este archivo para ver las definiciones de las peticiones y respuestas gRPC.

**Puerto de Service B:**  
Service B corre en el puerto **50051**.

## Flujo de Trabajo entre los Servicios

### Creación de usuario en Service A:
1. Cuando se realiza un **POST** en el endpoint `/users/` de **Service A**, se crea un nuevo usuario.
2. Después de crear el usuario, **Service A** hace una petición gRPC a **Service B**, enviando el ID del usuario recién creado.
3. **Service B** recibe el ID del usuario, realiza una consulta **GET** a **Service A** para obtener los datos del usuario.
4. Con la información del usuario, **Service B** envía un mensaje a RabbitMQ en la cola `service-b-notifies`, simulando el envío de un correo al usuario.

### Actualización de usuario en Service A:
1. Cuando se realiza un **PATCH** en el endpoint `/users/{id}` de **Service A**, se actualiza la información del usuario.
2. Al actualizar los datos, **Service A** envía un mensaje a RabbitMQ en la cola `service-a-notifies`, notificando a **Service B** de la actualización.
3. **Service B** consume el mensaje de la cola `service-a-notifies`, realiza un log con los datos obtenidos y procesa la notificación.

## Arquitecturas Utilizadas

- **Service A**: Se optó por una **arquitectura limpia (Clean Architecture)** para separar las responsabilidades de los diferentes componentes y garantizar que el código sea modular, escalable y fácilmente mantenible. Esta arquitectura también facilita las pruebas unitarias y la extensión futura del servicio.

- **Service B**: Se implementó una **arquitectura hexagonal (Ports and Adapters)**, lo que permite que el servicio sea independiente de las tecnologías externas (como RabbitMQ y gRPC) y facilita la integración con diferentes sistemas a través de puertos bien definidos.

## Pruebas Unitarias en Service A

Service A incluye pruebas unitarias para el servicio de gestión de usuarios. Estas pruebas están implementadas para garantizar el correcto funcionamiento de la creación, obtención y actualización de usuarios. Puedes ejecutar las pruebas con el siguiente comando:

```bash
make test-a
```

## Nota sobre rabbitMQ

Sí se desea observar las colas en la interfaz gráfica de rabbit se pueden dirigir al path http://localhost:15672/ el usuario y contreña es: `guest`



## Tecnologías Utilizadas

- **Go**: Lenguaje de programación utilizado para desarrollar ambos servicios.
- **Docker**: Para contenerizar los servicios y facilitar su despliegue.
- **RabbitMQ**: Utilizado para la comunicación entre los servicios mediante colas de mensajes.
- **gRPC**: Para la comunicación entre **Service A** y **Service B**.
- **MongoDB**: Base de datos utilizada por **Service A**.
