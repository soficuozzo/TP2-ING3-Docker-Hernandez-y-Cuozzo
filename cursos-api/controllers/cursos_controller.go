package controllers

import (
	"fmt"
	"net/http"
	"proyecto_arqui_soft_2/cursos-api/dao"
	"proyecto_arqui_soft_2/cursos-api/domain"

	"strconv"
	"strings"

	"proyecto_arqui_soft_2/cursos-api/services"

	"github.com/gin-gonic/gin"
)

// Definimos el controlador con el servicio
type CursoController struct {
	service services.CursoService
}

// Constructor para el controlador
func NewCursoController(service services.CursoService) CursoController {
	return CursoController{
		service: service,
	}
}

// Endpoint para inscribir un curso
func (controller CursoController) CrearInscripcion(c *gin.Context) {
	var inscripcion dao.Inscripcion
	if err := c.ShouldBindJSON(&inscripcion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Llamada al servicio para inscribir el curso
	err := controller.service.InscribirCurso(c.Request.Context(), inscripcion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Curso inscrito con éxito"})
}

// Endpoint para calcular la disponibilidad de múltiples cursos de manera concurrente
func (controller CursoController) CalcularDisponibilidadCursos(c *gin.Context) {
	var cursos []string
	if err := c.ShouldBindJSON(&cursos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Llamada al servicio para calcular la disponibilidad
	disponibilidad, err := controller.service.CalcularDisponibilidad(c.Request.Context(), cursos)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, disponibilidad)
}

func (controller CursoController) GetCursosByIds(c *gin.Context) {
	var requestBody struct {
		CursoIDs []string `json:"curso_ids"`
	}

	// Parsear el cuerpo de la solicitud
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if len(requestBody.CursoIDs) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "curso_ids cannot be empty"})
		return
	}

	// Obtener cursos desde el repositorio usando los IDs
	cursos, err := controller.service.GetCursosbyIds(c.Request.Context(), requestBody.CursoIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Devolver la lista de cursos
	c.JSON(http.StatusOK, gin.H{"cursos": cursos})

}

func (controller CursoController) GetInscripcionByUserId(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {

		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	cursoData, er := controller.service.GetInscripcionByUserId(c.Request.Context(), id)

	if er != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("error obteniendo curso: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"cursos": cursoData})
}

// Endpoint para obtener un curso por ID
func (controller CursoController) GetCursoByID(c *gin.Context) {
	cursoID := strings.TrimSpace(c.Param("id"))

	curso, err := controller.service.GetCursoByID(c.Request.Context(), cursoID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("error obteniendo curso: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, curso)
}

// Endpoint para obtener un curso por ID
func (controller CursoController) GetCursoByName(c *gin.Context) {
	cursoNombre := strings.TrimSpace(c.Param("name"))

	curso, err := controller.service.GetCursoByName(c.Request.Context(), cursoNombre)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("error obteniendo curso: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, curso)
}

// Endpoint para crear un curso
func (controller CursoController) Create(c *gin.Context) {
	var curso domain.CursoData
	if err := c.ShouldBindJSON(&curso); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("solicitud inválida: %s", err.Error()),
		})
		return
	}

	id, err := controller.service.Create(c.Request.Context(), curso)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error creando curso: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

// Endpoint para actualizar un curso
func (controller CursoController) Update(c *gin.Context) {
	id := strings.TrimSpace(c.Param("id"))

	var curso domain.CursoData
	if err := c.ShouldBindJSON(&curso); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("solicitud inválida: %s", err.Error()),
		})
		return
	}

	curso.CursoID = id

	if err := controller.service.Update(c.Request.Context(), curso); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error actualizando curso: %s", err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Curso actualizado",
	})
}

// Endpoint para eliminar un curso
func (controller CursoController) Delete(c *gin.Context) {
	id := strings.TrimSpace(c.Param("id"))
	if err := controller.service.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error eliminando curso: %s", err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Curso eliminado",
	})
}

// lo agregue para TODOS los cursos
func (controller CursoController) GetAllCursos(c *gin.Context) {

	cursoData, er := controller.service.GetAllCursos(c.Request.Context())

	if er != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error eliminando curso: %s", er.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, cursoData)

}
