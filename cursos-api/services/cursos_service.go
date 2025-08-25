package services

import (
	"context"
	"fmt"
	"proyecto_arqui_soft_2/cursos-api/dao"
	"proyecto_arqui_soft_2/cursos-api/domain"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
)

// Estructura para recibir solicitudes de inscripción

type DisponibilidadCurso struct {
	CursoID        string `json:"curso_id"`
	Disponibilidad int64  `json:"disponibilidad"`
}

// Interfaz del repositorio
type Repository interface {
	GetCursoByID(ctx context.Context, id string) (dao.Curso, error)
	Create(ctx context.Context, curso dao.Curso) (string, error)
	Update(ctx context.Context, id string, updateData bson.M) error
	Delete(ctx context.Context, id string) error
	InscribirCurso(ctx context.Context, inscripcion dao.Inscripcion) error

	GetInscripcionByUserId(ctx context.Context, userid int64) ([]dao.Inscripcion, error)
	GetAllCursos(ctx context.Context) ([]domain.CursoData, error)

	GetCursosByIds(tx context.Context, id []string) ([]domain.CursoData, error)
	GetCursoByName(ctx context.Context, name string) (dao.Curso, error)
}

// Definición de CursoService con repositorios y otros clientes
type CursoService struct {
	mainRepository        Repository
	inscripcionRepository Repository
	eventsQueue           Queue
}
type Queue interface {
	Publish(CursoNew domain.CursoNew) error
}

// Constructor para CursoService
func NewCursoService(mainRepository, inscripcionRepository Repository, eventsQueue Queue) CursoService {
	return CursoService{
		mainRepository:        mainRepository,
		inscripcionRepository: inscripcionRepository,
		eventsQueue:           eventsQueue,
	}
}

// Endpoint para inscribir un curso
func (service CursoService) InscribirCurso(ctx context.Context, inscripcion dao.Inscripcion) error {
	// Obtener curso
	curso, err := service.mainRepository.GetCursoByID(ctx, inscripcion.CursoID)
	if err != nil {
		return fmt.Errorf("error al obtener el curso: %w", err)
	}

	// Validar disponibilidad
	if curso.Capacidad <= 0 {
		return fmt.Errorf("no hay cupos disponibles para el curso")
	}

	service.inscripcionRepository.InscribirCurso(ctx, inscripcion)

	// Reducir capacidad y actualizar el curso
	curso.Capacidad--
	err = service.mainRepository.Update(ctx, curso.CursoID, bson.M{"capacidad": curso.Capacidad})
	if err != nil {
		return fmt.Errorf("error al actualizar el curso: %w", err)
	}

	return nil
}

func (service CursoService) GetCursosbyIds(ctx context.Context, id []string) ([]dao.Curso, error) {

	cursos, err := service.mainRepository.GetCursosByIds(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo curso: %w", err)
	}

	results := make([]dao.Curso, 0)
	for _, curso := range cursos {
		results = append(results, dao.Curso{

			CursoID:     curso.CursoID,
			Nombre:      curso.Nombre,
			Descripcion: curso.Descripcion,
			Categoria:   curso.Categoria,
			Capacidad:   curso.Capacidad,
			Requisito:   curso.Requisito,
			Duracion:    curso.Duracion,
			Imagen:      curso.Imagen,
			Valoracion:  curso.Valoracion,
			Profesor:    curso.Profesor,
		})
	}

	return results, nil
}

func (service CursoService) GetInscripcionByUserId(ctx context.Context, userid int64) ([]domain.CursoData, error) {

	inscripciones, err := service.inscripcionRepository.GetInscripcionByUserId(ctx, userid)

	if err != nil {
		return nil, fmt.Errorf("error al obtener el curso: %w", err)
	}

	// Imprimir las inscripciones para asegurarse de que contienen los datos esperados
	for _, inscripcion := range inscripciones {
		fmt.Println("Inscripcion completa:", inscripcion)           // Imprime la inscripción completa
		fmt.Println("CursoID en inscripcion:", inscripcion.CursoID) // Verifica que CursoID esté presente
	}

	var cursoids []string
	for _, inscripcion := range inscripciones {
		cursoids = append(cursoids, inscripcion.CursoID)
	}

	cursos, err := service.mainRepository.GetCursosByIds(ctx, cursoids)

	if err != nil {
		return nil, fmt.Errorf("error obteniendo cursos por IDs: %w", err)
	}

	return cursos, nil

}

// Endpoint para calcular la disponibilidad de múltiples cursos de manera concurrente
func (service CursoService) CalcularDisponibilidad(ctx context.Context, cursosIDs []string) ([]DisponibilidadCurso, error) {
	var wg sync.WaitGroup
	disponibilidad := make([]DisponibilidadCurso, len(cursosIDs))
	var mu sync.Mutex

	// Iterar sobre los IDs y lanzar una goroutine para cada curso
	for i, cursoID := range cursosIDs {
		wg.Add(1)
		go func(i int, cursoID string) {
			defer wg.Done()

			// Obtener curso
			curso, err := service.mainRepository.GetCursoByID(ctx, cursoID)
			if err != nil {
				// Si hay error, continuar sin bloquear el programa
				return
			}

			// Guardar disponibilidad en la estructura de resultados
			mu.Lock()
			disponibilidad[i] = DisponibilidadCurso{
				CursoID:        curso.CursoID,
				Disponibilidad: curso.Capacidad,
			}
			mu.Unlock()
		}(i, cursoID)
	}

	// Esperar a que terminen todas las goroutines
	wg.Wait()

	return disponibilidad, nil
}

func (service CursoService) GetCursoByID(ctx context.Context, id string) (domain.CursoData, error) {
	cursoDAO, err := service.mainRepository.GetCursoByID(ctx, id)
	if err != nil {
		return domain.CursoData{}, fmt.Errorf("error obteniendo curso: %w", err)
	}

	return domain.CursoData{
		CursoID:     cursoDAO.CursoID,
		Nombre:      cursoDAO.Nombre,
		Descripcion: cursoDAO.Descripcion,
		Categoria:   cursoDAO.Categoria,
		Capacidad:   cursoDAO.Capacidad,
		Requisito:   cursoDAO.Requisito,
		Duracion:    cursoDAO.Duracion,
		Imagen:      cursoDAO.Imagen,
		Valoracion:  cursoDAO.Valoracion,
		Profesor:    cursoDAO.Profesor,
	}, nil
}

func (service CursoService) GetCursoByName(ctx context.Context, name string) (domain.CursoData, error) {
	cursoDAO, err := service.mainRepository.GetCursoByName(ctx, name)

	if err != nil {
		return domain.CursoData{}, fmt.Errorf("error obteniendo curso: %w", err)
	}

	return domain.CursoData{
		CursoID:     cursoDAO.CursoID,
		Nombre:      cursoDAO.Nombre,
		Descripcion: cursoDAO.Descripcion,
		Categoria:   cursoDAO.Categoria,
		Capacidad:   cursoDAO.Capacidad,
		Requisito:   cursoDAO.Requisito,
		Duracion:    cursoDAO.Duracion,
		Imagen:      cursoDAO.Imagen,
		Valoracion:  cursoDAO.Valoracion,
		Profesor:    cursoDAO.Profesor,
	}, nil
}

func (service CursoService) Create(ctx context.Context, curso domain.CursoData) (string, error) {
	cursoDAO := dao.Curso{
		Nombre:      curso.Nombre,
		Descripcion: curso.Descripcion,
		Categoria:   curso.Categoria,
		Capacidad:   curso.Capacidad,
		Requisito:   curso.Requisito,
		Duracion:    curso.Duracion,
		Imagen:      curso.Imagen,
		Valoracion:  curso.Valoracion,
		Profesor:    curso.Profesor,
	}

	id, err := service.mainRepository.Create(ctx, cursoDAO)
	if err != nil {
		return "", fmt.Errorf("error creando curso: %w", err)
	}
	if err := service.eventsQueue.Publish(domain.CursoNew{
		Operation: "CREATE",
		CursoID:   id,
	}); err != nil {
		return "", fmt.Errorf("error publishing curso new: %w", err)
	}

	return id, nil
}

func (service CursoService) Update(ctx context.Context, curso domain.CursoData) error {
	updateData := bson.M{}
	if curso.Nombre != "" {
		updateData["nombre"] = curso.Nombre
	}
	if curso.Descripcion != "" {
		updateData["descripcion"] = curso.Descripcion
	}
	if curso.Categoria != "" {
		updateData["categoria_id"] = curso.Categoria
	}
	if curso.Capacidad != 0 {
		updateData["capacidad"] = curso.Capacidad
	}
	if curso.Profesor != "" {
		updateData["profesor"] = curso.Profesor
	}

	if curso.Requisito != "" {
		updateData["requisito"] = curso.Requisito
	}

	if curso.Duracion != 0 {
		updateData["duracion"] = curso.Duracion
	}

	if curso.Imagen != "" {
		updateData["imagen"] = curso.Imagen
	}

	if curso.Valoracion != 0 {
		updateData["valoracion"] = curso.Valoracion
	}

	if curso.Profesor != "" {
		updateData["profesor"] = curso.Profesor
	}

	if curso.Requisito != "" {
		updateData["requisito"] = curso.Requisito
	}

	if curso.Duracion != 0 {
		updateData["duracion"] = curso.Duracion
	}

	if curso.Imagen != "" {
		updateData["imagen"] = curso.Imagen
	}

	if curso.Valoracion != 0 {
		updateData["valoracion"] = curso.Valoracion
	}

	if len(updateData) == 0 {
		return fmt.Errorf("no hay campos para actualizar")
	}

	err := service.mainRepository.Update(ctx, curso.CursoID, updateData) // Se queda igual
	if err != nil {
		return fmt.Errorf("error actualizando curso: %w", err)
	}
	if err := service.eventsQueue.Publish(domain.CursoNew{
		Operation: "UPDATE",
		CursoID:   curso.CursoID,
	}); err != nil {
		return fmt.Errorf("error publishing curso update: %w", err)
	}

	return nil
}

func (service CursoService) Delete(ctx context.Context, id string) error {
	err := service.mainRepository.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("error eliminando curso: %w", err)
	}
	if err := service.eventsQueue.Publish(domain.CursoNew{
		Operation: "DELETE",
		CursoID:   id,
	}); err != nil {
		return fmt.Errorf("error publishing curso delete: %w", err)
	}
	return nil
}

// lo agregue para TODOS los cursos, para que vaya mostrando cursos a mis cursos mediante lo vas agregando
func (service CursoService) GetAllCursos(ctx context.Context) ([]domain.CursoData, error) {
	cursosDAO, err := service.mainRepository.GetAllCursos(ctx)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo cursos: %w", err)
	}

	var cursos []domain.CursoData
	for _, cursoDAO := range cursosDAO {
		cursos = append(cursos, domain.CursoData{
			CursoID:     cursoDAO.CursoID,
			Nombre:      cursoDAO.Nombre,
			Descripcion: cursoDAO.Descripcion,
			Categoria:   cursoDAO.Categoria,
			Capacidad:   cursoDAO.Capacidad,
			Requisito:   cursoDAO.Requisito,
			Duracion:    cursoDAO.Duracion,
			Imagen:      cursoDAO.Imagen,
			Valoracion:  cursoDAO.Valoracion,
			Profesor:    cursoDAO.Profesor,
		})
	}

	return cursos, nil
}
