package dao

// ContainerStatus representa el estado de un contenedor Docker
type ContainerStatus struct {
	Name   string `json:"container"`
	Status string `json:"status"`
}

// ContainerActionResponse representa la respuesta de una acción (start/stop)
type ContainerActionResponse struct {
	Container string `json:"container"`
	Action    string `json:"action"`
	Message   string `json:"message"`
}

// MultipleContainersStatus representa la respuesta para múltiples contenedores
type MultipleContainersStatus struct {
	Statuses map[string]string `json:"statuses"`
}
