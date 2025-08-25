package main

import (
	"context"
	"log"
	"proyecto_arqui_soft_2/cursos-api/clients"
	"proyecto_arqui_soft_2/cursos-api/controllers"
	"proyecto_arqui_soft_2/cursos-api/repositories"
	"proyecto_arqui_soft_2/cursos-api/services"

	"os"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {

	instanceID := os.Getenv("INSTANCE_ID")
	log.Printf("Starting instance: %s", instanceID)

	// Configuración del repositorio principal
	mainRepository := repositories.NewMongo(repositories.MongoConfig{
		Host:       "mongo",
		Port:       "27017",
		Username:   "root",
		Password:   "root",
		Database:   "cursos_db",
		Collection: "cursos",
	})

	InscripcionRepository := repositories.NewMySQL(
		repositories.MySQLConfig{
			Host:     "mysql",
			Port:     "3306",
			Database: "users-api",
			Username: "root",
			Password: "root1234",
		},
	)

	eventsQueue := clients.NewRabbit(clients.RabbitConfig{
		Host:      "rabbitmq",
		Port:      "5672",
		Username:  "user",
		Password:  "password",
		QueueName: "cursos-news",
	})

	// Prueba de conexión a MongoDB
	ctx := context.Background()
	if err := mainRepository.TestConnection(ctx); err != nil {
		log.Fatalf("Conexión a MongoDB fallida: %v", err)
	} else {
		log.Println("¡Conexión a MongoDB exitosa!")
	}

	// Crear servicio de cursos pasando el repositorio principal, el de inscripciones y la cola de eventos
	cursoService := services.NewCursoService(mainRepository, InscripcionRepository, eventsQueue)

	// Creación del controlador de cursos
	cursoController := controllers.NewCursoController(cursoService)

	// Configuración de rutas
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization")
	router.Use(cors.New(config))

	router.GET("/cursos/:id", cursoController.GetCursoByID)
	router.GET("/cursos/nombre/:name", cursoController.GetCursoByName)
	router.POST("/cursos/create", cursoController.Create)
	router.PUT("/cursos/update/:id", cursoController.Update)
	router.DELETE("/cursos/delete/:id", cursoController.Delete)
	router.POST("/inscripcion", cursoController.CrearInscripcion)
	router.GET("/usuario/miscursos/:id", cursoController.GetInscripcionByUserId)

	//lo agregue para TODOS los cursos
	router.GET("/cursos/todos", cursoController.GetAllCursos)

	// Inicia el servidor en el puerto 8081
	if err := router.Run(":8082"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
