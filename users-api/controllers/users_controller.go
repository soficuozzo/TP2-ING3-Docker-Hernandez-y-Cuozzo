package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	domain "proyecto_arqui_soft_2/users-api/domain"
	"strconv"
)

type Service interface {
	GetUsuariobyEmail(email string) (domain.UsuarioData, error)
	GetUsuariobyID(id int64) (domain.UsuarioData, error)
	CrearUsuario(newusuario domain.UsuarioData) (domain.UsuarioData, error)
	Login(email string, password string) (string, error)
}

type Controller struct {
	service Service
}

func NewController(service Service) Controller {
	return Controller{
		service: service,
	}
}

func (controller Controller) GetUsuariobyEmail(c *gin.Context) {

	email := c.Param("email")
	var usuarioData domain.UsuarioData

	usuarioData, err := controller.service.GetUsuariobyEmail(email)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("user not found: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, usuarioData)

}

func (controller Controller) GetUsuariobyID(c *gin.Context) {

	id := c.Param("id")
	usuarioID, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": fmt.Sprintf("invalid request: %s", err.Error()),
		})
		return
	}

	usuarioData, er := controller.service.GetUsuariobyID(usuarioID)

	if er != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("user not found: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, usuarioData)

}

func (controller Controller) CrearUsuario(c *gin.Context) {

	var newusuario domain.UsuarioData

	err := c.BindJSON(&newusuario)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": fmt.Sprintf("invalid request: %s", err.Error()),
		})
		return
	}

	newusuario, er := controller.service.CrearUsuario(newusuario)

	if er != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("user not found: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusCreated, newusuario)

}

func (controller Controller) Login(c *gin.Context) {
	var request domain.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.Resultado{
			Mensaje: fmt.Sprintf("Request invalido: %s", err.Error()),
		})
		return
	}

	token, err := controller.service.Login(request.Email, request.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.Resultado{
			Mensaje: fmt.Sprintf("Login no autorizado: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, domain.LoginResponse{
		Token: token,
	})
}
