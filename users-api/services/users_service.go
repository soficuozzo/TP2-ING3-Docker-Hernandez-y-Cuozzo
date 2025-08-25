package services

import (
	"crypto/md5"
	"errors"
	"fmt"
	dao "proyecto_arqui_soft_2/users-api/dao"
	domain "proyecto_arqui_soft_2/users-api/domain"

	//e "proyecto_arqui_soft_2/users-api/utils"
	"log"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Repository interface {
	GetUsuariobyEmail(email string) (dao.Usuario, error)
	GetUsuariobyID(id int64) (dao.Usuario, error)

	GenerarJWT(email string) (string, error)
	// actualizar cache y memcache
	Actualizar(usuario dao.Usuario) error
	CrearUsuario(newusuario dao.Usuario) (dao.Usuario, error)
}

type Service struct {
	mainRepository      Repository
	cacheRepository     Repository
	memcachedRepository Repository
}

func NewService(mainRepo, cacheRepo, memcachedRepo Repository) Service {
	return Service{
		mainRepository:      mainRepo,
		cacheRepository:     cacheRepo,
		memcachedRepository: memcachedRepo,
	}
}

func Usuario(user dao.Usuario) domain.UsuarioData {

	var us domain.UsuarioData

	us.Nombre = user.Nombre
	us.Apellido = user.Apellido
	us.Tipo = user.Tipo
	us.Email = user.Email
	us.Passwordhash = user.Passwordhash
	us.UsuarioID = user.UsuarioID

	return us
}

func GenerateHash(password string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(password)))
}

var jwtSecreto = []byte("llave_secreta")

func (service Service) GenerarJWT(email string) (string, error) {

	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtSecreto)
	if err != nil {
		// Log o manejo del error
		return "", fmt.Errorf("error generating token: %v", err)
	}

	return signedToken, nil
}

func (service Service) Gettoken(password string, usuario domain.UsuarioData) (string, error) {

	hash := GenerateHash(password)
	log.Println("contraseña hash:", hash)
	log.Println("hash:", usuario.Passwordhash)

	if hash != usuario.Passwordhash {
		return "", fmt.Errorf("Contraseña incorrecta.")
	}

	token, err := service.GenerarJWT(usuario.Email)

	if err != nil {
		return "", fmt.Errorf("Error al generar JWT token: %w", err)
	}
	return token, nil

}

func (service Service) Login(email string, password string) (string, error) {

	if strings.TrimSpace(email) == "" {
		return "", errors.New("Debe ingresar un email.")
	}

	if strings.TrimSpace(password) == "" {
		return "", errors.New("Debe ingresar una contraseña.")
	}

	var token string
	var error error

	usuarioo, err := service.cacheRepository.GetUsuariobyEmail(email)

	if err == nil {

		log.Println("Usuario encontrado en la caché.")

		result := Usuario(usuarioo)

		token, error = service.Gettoken(password, result)

		return token, error

	}

	usuarioo, err = service.memcachedRepository.GetUsuariobyEmail(email)

	if err == nil {
		log.Println("Usuario encontrado en la memcaché.")

		result := Usuario(usuarioo)

		token, error = service.Gettoken(password, result)

		fmt.Println("Hash generado:", GenerateHash(password))

		service.cacheRepository.Actualizar(usuarioo)

		return token, error

	}

	usuarioo, err = service.mainRepository.GetUsuariobyEmail(email)

	if err == nil {

		result := Usuario(usuarioo)

		token, error = service.Gettoken(password, result)

		// actualizar cache
		service.cacheRepository.Actualizar(usuarioo)

		// actualizar memcache
		service.memcachedRepository.Actualizar(usuarioo)

		return token, error

	} else {
		return "", fmt.Errorf("Hubo un error al buscar el usuario en la Base de Datos.")

	}

}

func (service Service) GetUsuariobyEmail(email string) (domain.UsuarioData, error) {

	// primero me fijo en la cache
	usuarioo, err := service.cacheRepository.GetUsuariobyEmail(email)

	if err == nil {

		result := Usuario(usuarioo)

		return result, nil

	}

	usuarioo, err = service.memcachedRepository.GetUsuariobyEmail(email)

	if err == nil {

		result := Usuario(usuarioo)

		service.cacheRepository.Actualizar(usuarioo)

		return result, nil

	}

	usuarioo, err = service.mainRepository.GetUsuariobyEmail(email)

	if err == nil {

		// actualizar cache
		service.cacheRepository.Actualizar(usuarioo)

		// actualizar memcache
		service.memcachedRepository.Actualizar(usuarioo)

		usuarioencontrado := Usuario(usuarioo)

		return usuarioencontrado, nil
	} else {
		return domain.UsuarioData{}, fmt.Errorf("error getting user by username: %w", err)
	}
}

func (service Service) CrearUsuario(newusuario domain.UsuarioData) (domain.UsuarioData, error) {

	var user dao.Usuario

	user.Nombre = newusuario.Nombre
	user.Apellido = newusuario.Apellido
	user.Email = newusuario.Email

	hash := GenerateHash(newusuario.Passwordhash)
	user.Passwordhash = hash
	newusuario.Passwordhash = user.Passwordhash

	user.Tipo = newusuario.Tipo

	user, err := service.mainRepository.CrearUsuario(user)

	if err != nil {
		return newusuario, fmt.Errorf("error creating user: %w", err)
	}

	newusuario.UsuarioID = user.UsuarioID

	service.memcachedRepository.CrearUsuario(user)
	service.cacheRepository.CrearUsuario(user)

	return newusuario, nil
}

func (service Service) GetUsuariobyID(id int64) (domain.UsuarioData, error) {

	// primero me fijo en la cache
	usuarioo, err := service.cacheRepository.GetUsuariobyID(id)

	if err == nil {
		log.Println("Datos obtenidos desde cache")

		result := Usuario(usuarioo)

		return result, nil

	}

	usuarioo, err = service.memcachedRepository.GetUsuariobyID(id)

	if err == nil {
		log.Println("Datos obtenidos desde Memcached")

		result := Usuario(usuarioo)

		service.cacheRepository.Actualizar(usuarioo)

		return result, nil

	}

	usuarioo, err = service.mainRepository.GetUsuariobyID(id)

	if err == nil {
		log.Println("Datos obtenidos desde base de datos")

		// actualizar cache
		service.cacheRepository.Actualizar(usuarioo)

		// actualizar memcache
		service.memcachedRepository.Actualizar(usuarioo)

		usuarioencontrado := Usuario(usuarioo)

		return usuarioencontrado, nil

	} else {
		return domain.UsuarioData{}, fmt.Errorf("error getting user by ID: %w", err)
	}

}

func (service Service) Actualizar(usuario domain.UsuarioData) error {

	var apellidoA, nombreA, passwordA string

	if usuario.Apellido != "" {
		apellidoA = usuario.Apellido
	} else {
		usuarioo, err := service.mainRepository.GetUsuariobyID(usuario.UsuarioID)
		if err != nil {
			return fmt.Errorf("error para devolver usuario: %w", err)
		}
		apellidoA = usuarioo.Apellido

	}

	if usuario.Nombre != "" {
		nombreA = usuario.Nombre
	} else {
		usuarioo, err := service.mainRepository.GetUsuariobyID(usuario.UsuarioID)
		if err != nil {
			return fmt.Errorf("error para devolver usuario: %w", err)
		}
		nombreA = usuarioo.Nombre

	}

	if usuario.Passwordhash != "" {
		passwordA = GenerateHash(usuario.Passwordhash)
	} else {
		usuarioo, err := service.mainRepository.GetUsuariobyID(usuario.UsuarioID)
		if err != nil {
			return fmt.Errorf("error para devolver usuario: %w", err)
		}
		passwordA = usuarioo.Passwordhash

	}

	usuarioActualizado := dao.Usuario{
		UsuarioID:    usuario.UsuarioID,
		Nombre:       nombreA,
		Apellido:     apellidoA,
		Email:        usuario.Email,
		Tipo:         usuario.Tipo,
		Passwordhash: passwordA,
	}

	err := service.mainRepository.Actualizar(usuarioActualizado)

	if err != nil {
		return fmt.Errorf("error actualizando usuario: %w", err)
	}

	err = service.cacheRepository.Actualizar(usuarioActualizado)

	if err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}

	err = service.memcachedRepository.Actualizar(usuarioActualizado)

	if err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}

	return nil
}
