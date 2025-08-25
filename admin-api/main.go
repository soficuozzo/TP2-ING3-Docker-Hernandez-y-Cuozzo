package main

import (
	"fmt"
	"proyecto_arqui_soft_2/admin-api/controller"
	"proyecto_arqui_soft_2/admin-api/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Crear el servicio y controlador de contenedores
	containerService := service.NewContainerService()
	containerController := controller.NewContainerController(containerService)

	// Crear una nueva instancia del router
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization")
	router.Use(cors.New(config))

	// URL mappings
	router.GET("/containers/:name", containerController.GetContainerStatus)
	router.POST("/containers/:action/:name", containerController.ManageContainer)
	router.GET("/containers", containerController.GetContainersStatus)

	// Ejecutar el servidor
	fmt.Println("Servidor corriendo en el puerto 8086")
	router.Run(":8086")
}
