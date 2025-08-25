package service

import (
	"fmt"
	"os/exec"
	"strings"
)

type ContainerService struct{}

func NewContainerService() *ContainerService {
	return &ContainerService{}
}

type ContainerStatus struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

// Obtener el estado de un contenedor por nombre
func (s *ContainerService) GetContainerStatus(containerName string) string {
	cmd := exec.Command("docker", "inspect", "--format", "{{.State.Status}}", containerName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error obteniendo estado del contenedor %s: %v\n", containerName, err)
		return "unknown"
	}
	return strings.TrimSpace(string(output))
}

// Manejar la acción de un contenedor (start/stop)
func (s *ContainerService) ManageContainer(containerName, action string) error {
	if action != "start" && action != "stop" {
		return fmt.Errorf("acción inválida: %s", action)
	}

	cmd := exec.Command("docker", action, containerName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error al ejecutar %s en el contenedor %s: %v\n%s", action, containerName, err, string(output))
	}

	fmt.Printf("Contenedor %s %s correctamente\n", containerName, action)
	return nil
}

// Obtener el estado de múltiples contenedores
func (s *ContainerService) GetContainersStatus(containerNames []string) []ContainerStatus {
	var statuses []ContainerStatus
	for _, name := range containerNames {
		statuses = append(statuses, ContainerStatus{
			Name:   name,
			Status: s.GetContainerStatus(name),
		})
	}
	return statuses
}
