package main

import (
	"log"
	"proyecto_arqui_soft_2/search-api/clients/queues"
	controllers "proyecto_arqui_soft_2/search-api/controllers"
	repositories "proyecto_arqui_soft_2/search-api/repositories"
	services "proyecto_arqui_soft_2/search-api/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Solr
	solrRepo := repositories.NewSolr(repositories.SolrConfig{
		Host:       "solr",  // Solr host
		Port:       "8983",  // Solr port
		Collection: "curso", // Collection name
	})

	// Rabbit
	eventsQueue := queues.NewRabbit(queues.RabbitConfig{
		Host:      "rabbitmq",
		Port:      "5672",
		Username:  "user",
		Password:  "password",
		QueueName: "cursos-news",
	})

	// Cursos API
	cursosAPI := repositories.NewHTTP(repositories.HTTPConfig{
		Host: "nginx",
		Port: "8082",
	})

	// Services
	service := services.NewService(solrRepo, cursosAPI)

	// Controllers
	controller := controllers.NewController(service)

	// Launch rabbit consumer
	if err := eventsQueue.StartConsumer(service); err != nil {
		log.Fatalf("Error running consumer: %v", err)
	}

	// Create router
	router := gin.Default()

	// Configurar CORS (Permitir todos los orígenes)
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization")
	router.Use(cors.New(config))

	router.GET("/search", controller.Search)

	if err := router.Run(":8083"); err != nil {
		log.Fatalf("Error running application: %v", err)
	}
}
