package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	cursosDomain "proyecto_arqui_soft_2/search-api/domain"
)

type HTTPConfig struct {
	Host string
	Port string
}

type HTTP struct {
	baseURL func(CursoID string) string
}

func NewHTTP(config HTTPConfig) HTTP {
	return HTTP{
		baseURL: func(CursoID string) string {
			return fmt.Sprintf("http://%s:%s/cursos/%s", config.Host, config.Port, CursoID)
		},
	}
}

func (repository HTTP) GetCursoByID(ctx context.Context, id string) (cursosDomain.CursoData, error) {
	url := repository.baseURL(id)
	fmt.Println("URL: ", url)
	resp, err := http.Get(url)
	if err != nil {
		return cursosDomain.CursoData{}, fmt.Errorf("Error fetching curso (%s): %w\n", id, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return cursosDomain.CursoData{}, fmt.Errorf("Failed to fetch cursos (%s): received status code %d\n", id, resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return cursosDomain.CursoData{}, fmt.Errorf("Error reading response body for cursos (%s): %w\n", id, err)
	}

	// Unmarshal the cursos details into the cursos struct
	var cursos cursosDomain.CursoData
	if err := json.Unmarshal(body, &cursos); err != nil {
		return cursosDomain.CursoData{}, fmt.Errorf("Error unmarshaling cursos data (%s): %w\n", id, err)
	}

	return cursos, nil
}