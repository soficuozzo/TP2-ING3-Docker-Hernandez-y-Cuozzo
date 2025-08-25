package controller

import (
	"net/http"

	"proyecto_arqui_soft_2/admin-api/service"

	"github.com/gin-gonic/gin"
)

type ContainerController struct {
	containerService *service.ContainerService
}

func NewContainerController(service *service.ContainerService) *ContainerController {
	return &ContainerController{containerService: service}
}

// Obtener estado de un contenedor
func (cc *ContainerController) GetContainerStatus(c *gin.Context) {
	containerName := c.Param("name")
	status := cc.containerService.GetContainerStatus(containerName)
	c.JSON(http.StatusOK, gin.H{"container": containerName, "status": status})
}

// Manejar la acción de un contenedor (start/stop)
func (cc *ContainerController) ManageContainer(c *gin.Context) {
	containerName := c.Param("name")
	action := c.Param("action")

	err := cc.containerService.ManageContainer(containerName, action)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"container": containerName,
		"action":    action,
		"message":   "Acción ejecutada correctamente",
	})
}

// Obtener estado de múltiples contenedores
func (cc *ContainerController) GetContainersStatus(c *gin.Context) {

	// los controladores que va a poder ver para parar / iniciar

	containerNames := []string{
		"cursos-api-container", "cursos-api-container-1", "search-api-container", "users-api-container",
	}
	statuses := cc.containerService.GetContainersStatus(containerNames)
	c.JSON(http.StatusOK, statuses)
}
