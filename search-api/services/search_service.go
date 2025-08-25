package services

import (
	"context"
	"fmt"      
	cursosDomain "proyecto_arqui_soft_2/search-api/domain"
	cursosDAO "proyecto_arqui_soft_2/search-api/dao" 
)

type Repository interface {
	Index(ctx context.Context, curso cursosDAO.Curso) (string, error)
	Update(ctx context.Context, curso cursosDAO.Curso) error
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, query string, limit int, offset int) ([]cursosDAO.Curso, error) 
}

type ExternalRepository interface {
	GetCursoByID(ctx context.Context, id string) (cursosDomain.CursoData, error)
}

type Service struct {
	repository Repository
	cursosAPI  ExternalRepository
}

func NewService(repository Repository, cursosAPI ExternalRepository) Service {
	return Service{
		repository: repository,
		cursosAPI:  cursosAPI,
	}
}

func (service Service) Search(ctx context.Context, query string, offset int, limit int) ([]cursosDomain.CursoData, error) {
	// Llama al método Search del repositorio
	cursosDAOList, err := service.repository.Search(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error searching cursos: %w", err)
	}

	// Convierte los cursos del DAO al dominio
	cursosDomainList := make([]cursosDomain.CursoData, 0)
	for _, curso := range cursosDAOList {
		cursosDomainList = append(cursosDomainList, cursosDomain.CursoData{
			CursoID:     curso.CursoID,
			Nombre:      curso.Nombre,
			Descripcion: curso.Descripcion,
			Categoria:   curso.Categoria,
			Capacidad:   curso.Capacidad,
			Imagen: 	curso.Imagen,
			Valoracion:   curso.Valoracion,
			Requisito:   curso.Requisito,
			Profesor:   curso.Profesor,
			Duracion: curso.Duracion,
		})
	}

	return cursosDomainList, nil
}

func (service Service) HandleCursoNew(cursoNew cursosDomain.CursoNew) {
	switch cursoNew.Operation {
	case "CREATE", "UPDATE":
		// Obtiene los detalles del curso desde el servicio externo
		cursoData, err := service.cursosAPI.GetCursoByID(context.Background(), cursoNew.CursoID)
		if err != nil {
			fmt.Printf("Error getting curso (%s) from API: %v\n", cursoNew.CursoID, err)
			return
		}

		// Convierte los datos del dominio al formato del DAO
		cursoDAO := cursosDAO.Curso{
			CursoID:     cursoData.CursoID,
			Nombre:      cursoData.Nombre,
			Descripcion: cursoData.Descripcion,
			Categoria:   cursoData.Categoria,
			Capacidad:   cursoData.Capacidad,
			Imagen: 	cursoData.Imagen,
			Valoracion:   cursoData.Valoracion,
			Requisito:   cursoData.Requisito,
			Profesor:   cursoData.Profesor,
			Duracion: cursoData.Duracion,
		}

		// Maneja la operación de Indexación
		if cursoNew.Operation == "CREATE" {
			if _, err := service.repository.Index(context.Background(), cursoDAO); err != nil {
				fmt.Printf("Error indexing curso (%s): %v\n", cursoNew.CursoID, err)
			} else {
				fmt.Println("Curso indexed successfully:", cursoNew.CursoID)
			}
		} else { 
			if err := service.repository.Update(context.Background(), cursoDAO); err != nil {
				fmt.Printf("Error updating curso (%s): %v\n", cursoNew.CursoID, err)
			} else {
				fmt.Println("Curso updated successfully:", cursoNew.CursoID)
			}
		}

	case "DELETE":
		
		if err := service.repository.Delete(context.Background(), cursoNew.CursoID); err != nil {
			fmt.Printf("Error deleting curso (%s): %v\n", cursoNew.CursoID, err)
		} else {
			fmt.Println("Curso deleted successfully:", cursoNew.CursoID)
		}

	default:
		fmt.Printf("Unknown operation: %s\n", cursoNew.Operation)
	}
}
