package services_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	dao "proyecto_arqui_soft_2/users-api/dao"
	domain "proyecto_arqui_soft_2/users-api/domain"
	repositories "proyecto_arqui_soft_2/users-api/repositories"
	servicio "proyecto_arqui_soft_2/users-api/services"
)

var (
	// Creamos mocks
	mainRepo      = repositories.NewMock()
	cacheRepo     = repositories.NewMock()
	memcachedRepo = repositories.NewMock()

	usersService = servicio.NewService(mainRepo, cacheRepo, memcachedRepo)
)

func TestService(t *testing.T) {

	// funciona
	t.Run("GetUsuariobyID - Success from Cache", func(t *testing.T) {
		hashedPassword := servicio.GenerateHash("password123")
		mockUser := dao.Usuario{UsuarioID: 1, Nombre: "nombre1", Apellido: "apellido1", Email: "email1", Tipo: "estudiante", Passwordhash: hashedPassword}

		// nombre de la función + el parametro q necesita
		cacheRepo.On("GetUsuariobyID", int64(1)).Return(mockUser, nil).Once()

		result, err := usersService.GetUsuariobyID(1)

		assert.NoError(t, err)
		assert.Equal(t, "email1", result.Email)

		mainRepo.AssertExpectations(t)
		cacheRepo.AssertExpectations(t)
		memcachedRepo.AssertExpectations(t)
	})

	// funciona
	t.Run("GetUsuariobyID - Not Found in Cache, Found in Memcached", func(t *testing.T) {
		hashedPassword := servicio.GenerateHash("password123")
		mockUser := dao.Usuario{UsuarioID: 1, Nombre: "nombre1", Apellido: "apellido1", Email: "email1", Tipo: "estudiante", Passwordhash: hashedPassword}

		cacheRepo.On("GetUsuariobyID", int64(1)).Return(dao.Usuario{}, errors.New("not found")).Once()
		memcachedRepo.On("GetUsuariobyID", int64(1)).Return(mockUser, nil).Once()

		cacheRepo.On("Actualizar", mockUser).Return(nil).Once()

		result, err := usersService.GetUsuariobyID(1)

		assert.NoError(t, err)
		assert.Equal(t, "email1", result.Email)

		mainRepo.AssertExpectations(t)
		cacheRepo.AssertExpectations(t)
		memcachedRepo.AssertExpectations(t)
	})

	// funciona
	t.Run("GetUsuariobyID - Not Found in Cache or Memcached, Found in Main Repo", func(t *testing.T) {
		hashedPassword := servicio.GenerateHash("password123")
		mockUser := dao.Usuario{UsuarioID: 1, Nombre: "nombre1", Apellido: "apellido1", Email: "email1", Tipo: "estudiante", Passwordhash: hashedPassword}

		cacheRepo.On("GetUsuariobyID", int64(1)).Return(dao.Usuario{}, errors.New("not found")).Once()
		memcachedRepo.On("GetUsuariobyID", int64(1)).Return(dao.Usuario{}, errors.New("not found")).Once()
		mainRepo.On("GetUsuariobyID", int64(1)).Return(mockUser, nil).Once()

		cacheRepo.On("Actualizar", mockUser).Return(nil).Once()
		memcachedRepo.On("Actualizar", mockUser).Return(nil).Once()

		result, err := usersService.GetUsuariobyID(1)

		assert.NoError(t, err)
		assert.Equal(t, "email1", result.Email)

		mainRepo.AssertExpectations(t)
		cacheRepo.AssertExpectations(t)
		memcachedRepo.AssertExpectations(t)
	})

	// funciona
	t.Run("GetUsuariobyID - Error in Main Repo", func(t *testing.T) {
		cacheRepo.On("GetUsuariobyID", int64(1)).Return(dao.Usuario{}, errors.New("not found")).Once()
		memcachedRepo.On("GetUsuariobyID", int64(1)).Return(dao.Usuario{}, errors.New("not found")).Once()
		mainRepo.On("GetUsuariobyID", int64(1)).Return(dao.Usuario{}, errors.New("db error")).Once()

		result, err := usersService.GetUsuariobyID(1)

		assert.Error(t, err)
		assert.Equal(t, "error getting user by ID: db error", err.Error())
		assert.Equal(t, domain.UsuarioData{}, result)

		mainRepo.AssertExpectations(t)
		cacheRepo.AssertExpectations(t)
		memcachedRepo.AssertExpectations(t)
	})

	// funciona
	t.Run("CrearUsuario - Success", func(t *testing.T) {
		hashedPassword := servicio.GenerateHash("password123")
		newUser := dao.Usuario{Nombre: "nombre1", Apellido: "apellido1", Email: "email1", Tipo: "estudiante", Passwordhash: hashedPassword}

		mainRepo.On("CrearUsuario", newUser).Return((newUser), nil).Once()

		newUser.UsuarioID = 0

		cacheRepo.On("CrearUsuario", newUser).Return((newUser), nil).Once()
		memcachedRepo.On("CrearUsuario", newUser).Return((newUser), nil).Once()

		usuario, err := usersService.CrearUsuario(domain.UsuarioData{Nombre: "nombre1", Apellido: "apellido1", Email: "email1", Tipo: "estudiante", Passwordhash: "password123"})

		resultUser := dao.Usuario{
			UsuarioID:    usuario.UsuarioID,
			Nombre:       usuario.Nombre,
			Apellido:     usuario.Apellido,
			Email:        usuario.Email,
			Tipo:         usuario.Tipo,
			Passwordhash: usuario.Passwordhash,
		}

		assert.NoError(t, err)
		assert.Equal(t, (newUser), resultUser)

		mainRepo.AssertExpectations(t)
		cacheRepo.AssertExpectations(t)
		memcachedRepo.AssertExpectations(t)
	})

	// funciona
	t.Run("CrearUsuario - Error", func(t *testing.T) {

		hashedPassword := servicio.GenerateHash("password123")
		newUser := dao.Usuario{Nombre: "nombre1", Apellido: "apellido1", Email: "email1", Tipo: "estudiante", Passwordhash: hashedPassword}

		mainRepo.On("CrearUsuario", newUser).Return(newUser, errors.New("db error")).Once()

		usuario, err := usersService.CrearUsuario(domain.UsuarioData{Nombre: "nombre1", Apellido: "apellido1", Email: "email1", Tipo: "estudiante", Passwordhash: "password123"})

		resultUser := dao.Usuario{
			UsuarioID:    usuario.UsuarioID,
			Nombre:       usuario.Nombre,
			Apellido:     usuario.Apellido,
			Email:        usuario.Email,
			Tipo:         usuario.Tipo,
			Passwordhash: usuario.Passwordhash,
		}

		assert.Error(t, err)
		assert.Equal(t, (newUser), resultUser)
		assert.Equal(t, "error creating user: db error", err.Error())

		mainRepo.AssertExpectations(t)
		cacheRepo.AssertExpectations(t)
		memcachedRepo.AssertExpectations(t)
	})

	//funciona
	t.Run("Login - Success", func(t *testing.T) {
		email := "email1"
		password := "password"
		hashedPassword := servicio.GenerateHash(password)

		mockUser := dao.Usuario{UsuarioID: 1, Email: email, Passwordhash: hashedPassword}
		cacheRepo.On("GetUsuariobyEmail", email).Return(mockUser, nil).Once()

		response, err := usersService.Login(email, password)

		assert.NoError(t, err)
		assert.Equal(t, int64(1), mockUser.UsuarioID)
		assert.NotEmpty(t, response)

		mainRepo.AssertExpectations(t)
		cacheRepo.AssertExpectations(t)
		memcachedRepo.AssertExpectations(t)
	})

	// funciona
	t.Run("Login - Invalid Credentials", func(t *testing.T) {
		email := "email1"
		password := "passwordmala"
		hashedPassword := servicio.GenerateHash("password")

		mockUser := dao.Usuario{UsuarioID: 1, Email: email, Passwordhash: hashedPassword}
		cacheRepo.On("GetUsuariobyEmail", email).Return(mockUser, nil).Once()

		response, err := usersService.Login(email, password)

		assert.Error(t, err)
		assert.Equal(t, "Contraseña incorrecta.", err.Error())
		assert.Equal(t, "", response)

		mainRepo.AssertExpectations(t)
		cacheRepo.AssertExpectations(t)
		memcachedRepo.AssertExpectations(t)
	})

	// funciona
	t.Run("Login - User Not Found", func(t *testing.T) {
		email := "email1"
		password := "password"

		cacheRepo.On("GetUsuariobyEmail", email).Return(dao.Usuario{}, errors.New("not found")).Once()
		memcachedRepo.On("GetUsuariobyEmail", email).Return(dao.Usuario{}, errors.New("not found")).Once()
		mainRepo.On("GetUsuariobyEmail", email).Return(dao.Usuario{}, errors.New("not found")).Once()

		response, err := usersService.Login(email, password)

		assert.Error(t, err)
		assert.Equal(t, "Hubo un error al buscar el usuario en la Base de Datos.", err.Error())
		assert.Equal(t, "", response)

		mainRepo.AssertExpectations(t)
		cacheRepo.AssertExpectations(t)
		memcachedRepo.AssertExpectations(t)
	})

	// ver
	// /Users/lucianahueda/Documents/Final-Arqui/proyecto_arqui_soft_2/users-api/services/users_service_test.go:232:
	// Error Trace:	/Users/lucianahueda/Documents/Final-Arqui/proyecto_arqui_soft_2/users-api/services/users_service_test.go:232
	// Error:      	An error is expected but got nil.
	// Test:       	TestService/Login_-_Token_Generation_Error
	t.Run("Login - Token Generation Error", func(t *testing.T) {
		email := "email1"
		password := "password"
		hashedPassword := servicio.GenerateHash(password)

		mockUser := dao.Usuario{UsuarioID: 1, Email: email, Passwordhash: hashedPassword}
		cacheRepo.On("GetUsuariobyEmail", email).Return(mockUser, nil).Once()

		response, err := usersService.Login(email, password)

		assert.Error(t, err)
		assert.Equal(t, "error generating token: token error", err.Error())
		assert.Equal(t, "", response)

		mainRepo.AssertExpectations(t)
		cacheRepo.AssertExpectations(t)
		memcachedRepo.AssertExpectations(t)
	})

}
