Trabajo Práctico 02 – Introducción a Docker
1. Elegir y preparar tu aplicación

Elegimos una aplicación web desarrollada en Go, llamada cursos-api, que forma parte de un sistema de cursos online. Esta API permite manejar operaciones de cursos y usuarios.

Creamos un repositorio en GitHub llamado TP2-ING3-Docker-Hernandez-y-Cuozzo, lo clonamos con git clone y copiamos todos los archivos de la aplicación dentro de la carpeta del repositorio.

Configuramos el entorno Docker instalando Docker Desktop. Verificamos la instalación con:

docker version

2. Construir una imagen personalizada

Creamos un Dockerfile para la aplicación.

Elegimos como base la imagen golang:1.23-alpine.

Etiquetamos la imagen con nuestro usuario de Docker Hub y un nombre significativo.

Dockerfile:
FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .
RUN go build -o app ./main.go

EXPOSE 8082

CMD ["./app"]

Explicación línea por línea:

FROM golang:1.23-alpine: usamos la versión oficial de Go 1.23 sobre Alpine Linux, lo que da una imagen más liviana y segura.

WORKDIR /app: define /app como directorio de trabajo donde se ejecutarán los comandos siguientes.

COPY go.mod go.sum ./: copiamos primero los archivos de dependencias para que Docker aproveche el cache si no cambian.

RUN go mod tidy: descarga y organiza las dependencias del proyecto.

COPY . .: copia todo el código fuente al contenedor.

RUN go build -o app ./main.go: compila la aplicación Go y genera el binario app.

EXPOSE 8082: documenta el puerto en el que la aplicación escuchará.

CMD ["./app"]: indica que al correr el contenedor se ejecute el binario compilado.

Justificación de la imagen base:

Se eligió golang:1.23-alpine porque:

Incluye Go preinstalado en su última versión estable.

Está basada en Alpine, una distribución ligera, reduciendo el tamaño final de la imagen.

Mantiene un buen balance entre facilidad de uso (trae Go y herramientas listas) y eficiencia.

Etiquetado de la imagen

Construimos la imagen con:

docker build -t uccsoficuozzo/cursos-api:v1.0 .


3. Publicar la imagen en Docker Hub
Para publicar la imagen en Docker Hub, primero realizamos el login con:
docker login
Luego subimos la imagen ya construida y etiquetada con nuestro usuario:

docker push soficuzzo/cursos-api:v1.0

De esta forma la imagen queda disponible en el repositorio de Docker Hub y puede ser utilizada en cualquier máquina que tenga Docker instalado, con el comando:

docker pull soficuozzo/cursos-api:v1.0

Estrategia de versionado

Para el versionado de imágenes seguimos la convención SemVer (MAJOR.MINOR.PATCH):
MAJOR: cambios incompatibles o de arquitectura.
MINOR: nuevas funcionalidades que mantienen compatibilidad.
PATCH: correcciones menores o parches de seguridad.


4. Integrar una base de datos en contenedor

Bases elegidas: MySQL 8.0 (relacional) y MongoDB 7 (documental).
Motivación del proyecto: la API cursos-api necesita almacenar datos estructurados y transaccionales (usuarios, inscripciones, pagos) junto con contenido flexible (módulos/temas, metadatos variables). Por eso combinamos:

MySQL para integridad referencial, transacciones ACID y consultas relacionales.

MongoDB para esquema flexible y agregaciones sobre contenidos que cambian con frecuencia.

Volúmenes persistentes:

mysql-data:/var/lib/mysql y mongo-data:/data/db aseguran persistencia de datos más allá del ciclo de vida de los contenedores.

Conexión de la aplicación:

La app lee variables de entorno inyectadas por docker-compose.yml.

MySQL (DSN): user:pass@tcp(mysql:3306)/cursos?parseTime=true

MongoDB (URI): mongodb://root:root1234@mongo:27017/?authSource=admin

Se configuró depends_on con healthchecks para que la aplicación arranque solo cuando las bases estén healthy, reduciendo fallos de arranque por conexiones tempranas.

Justificación técnica:

MySQL 8.0: madurez, soporte a utf8mb4, buen rendimiento OLTP y driver Go estable.

MongoDB 7: esquema flexible orientado a documentos, pipelines de agregación potentes y facilidad para modelar estructuras de curso dinámicas.

Buenas prácticas (dev/QA): credenciales en variables de entorno (en PROD, usar secretos), principio de menor privilegio para usuarios de DB, backups de volúmenes y observabilidad con healthchecks.

¡Acá tenés el bloque **listo para pegar** en tu `decisiones.md` para el **Paso 5** (QA/PROD con la misma imagen). Incluye explicación, variables, comandos y una mini sección de troubleshooting. 👇

---

## 5. Configurar QA y PROD con la misma imagen

**Estrategia general.**
Se reutiliza la **misma imagen** de la aplicación (`uccsimon/tp2-app:v1.0`) y se ejecutan **dos contenedores** en paralelo: uno para **QA** y otro para **PROD**. La **única diferencia** entre entornos son las **variables de entorno** y los **puertos expuestos**; el binario y dependencias son idénticos. Esto reduce el “drift” entre entornos y hace que lo probado en QA sea representativo de PROD.

**Variables de entorno (principales) y propósito.**

* `APP_ENV`: define el entorno (`qa` / `prod`) para habilitar comportamientos específicos.
* `LOG_LEVEL`: `debug` (QA) vs `error` (PROD) para controlar la verbosidad de logs.
* **MySQL**: `DB_HOST`, `DB_PORT`, `DB_NAME`, `DB_USER`, `DB_PASSWORD`.

  * **Aislamiento de datos** usando nombres de base distintos: `cursos_qa` (QA) y `cursos_prod` (PROD).
* **MongoDB**: `MONGO_HOST`, `MONGO_PORT`, `MONGO_USER`, `MONGO_PASSWORD`, `MONGO_DB`.

  * **Aislamiento de datos** con `cursosdb_qa` (QA) y `cursosdb_prod` (PROD).
* `PORT`: puerto interno de la app (8082).

  * Hacia el host se publican puertos **distintos** para convivir: QA en `8082`, PROD en `8089`.

**Implementación en `docker-compose.yml` (resumen).**
Se definen dos servicios de app con la **misma imagen**:

```yaml
tp2-app-qa:
  image: uccsimon/tp2-app:v1.0
  environment:
    APP_ENV: qa
    LOG_LEVEL: debug
    PORT: "8082"
    DB_HOST: mysql
    DB_PORT: "3306"
    DB_NAME: cursos_qa
    DB_USER: cursos_user
    DB_PASSWORD: cursos_pass
    MONGO_HOST: mongo
    MONGO_PORT: "27017"
    MONGO_USER: root
    MONGO_PASSWORD: root1234
    MONGO_DB: cursosdb_qa
  ports: ["8082:8082"]

tp2-app-prod:
  image: uccsimon/tp2-app:v1.0
  environment:
    APP_ENV: prod
    LOG_LEVEL: error
    PORT: "8082"
    DB_HOST: mysql
    DB_PORT: "3306"
    DB_NAME: cursos_prod
    DB_USER: cursos_user
    DB_PASSWORD: cursos_pass
    MONGO_HOST: mongo
    MONGO_PORT: "27017"
    MONGO_USER: root
    MONGO_PASSWORD: root1234
    MONGO_DB: cursosdb_prod
  ports: ["8089:8082"]
```

Ambos servicios **dependen** de `mysql` y `mongo` (con `depends_on` y *healthchecks*) para garantizar que las bases estén listas antes de iniciar la app.

**Cómo se aplican las variables en la app.**
Al comenzar, la aplicación lee las variables (ej. `os.Getenv` en Go) para:

* Construir las cadenas de conexión:

  * MySQL (QA): `cursos_user:cursos_pass@tcp(mysql:3306)/cursos_qa?parseTime=true`
  * Mongo (PROD): `mongodb://root:root1234@mongo:27017/?authSource=admin` usando `cursosdb_prod`
* Ajustar el **nivel de logs** y cualquier feature flag por entorno.
* Escuchar en el puerto indicado (`PORT=8082` dentro del contenedor).

**Comandos ejecutados (QA y PROD a la vez).**

```bash
# Levantar todo (DBs + QA + PROD)
docker compose up -d

# Ver estado
docker compose ps

# Logs por servicio
docker compose logs -f tp2-app-qa
docker compose logs -f tp2-app-prod
```

**URLs de prueba:**

* QA: `http://localhost:8082`
* PROD: `http://localhost:8089`

**Justificación.**

* **Misma imagen** ⇒ elimina discrepancias entre entornos; lo probado en QA es lo que llega a PROD.
* **Config por variables** ⇒ práctica 12-factor; no se recompila ni reconstruye la imagen al cambiar de entorno.
* **Aislamiento de datos** ⇒ bases diferenciadas (`_qa` vs `_prod`), evitando mezclar información.
* **Puertos distintos** ⇒ permiten correr QA y PROD al mismo tiempo en una misma máquina para comparaciones y validaciones rápidas.

**Troubleshooting rápido.**

* Error de validación tipo `additional properties '+version' not allowed`:
  Asegurarse de que la primera línea del compose sea exactamente `version: "3.8"` (sin `+` ni espacios extra).
* Variables en `environment`: usar formato **mapa** (`CLAVE: valor`) o **lista sin espacios** (`- CLAVE=valor`). Evitar `CLAVE = valor` (con espacios), ya que no las toma.

---
