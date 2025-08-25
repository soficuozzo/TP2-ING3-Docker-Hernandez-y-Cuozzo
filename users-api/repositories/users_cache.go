package repositories

import (
	//"fmt"
	"fmt"
	"proyecto_arqui_soft_2/users-api/dao"
	"proyecto_arqui_soft_2/users-api/domain"

	"time"

	"github.com/karlseguin/ccache"
)

type CacheConfig struct {
	TTL time.Duration // Cache expiration time
}

type Cache struct {
	client *ccache.Cache
	ttl    time.Duration
}

// GenerarJWT implements services.Repository.
func (repository Cache) GenerarJWT(email string) (string, error) {
	panic("unimplemented")
}

func NewCache(config CacheConfig) Cache {
	// Initialize ccache with default settings
	cache := ccache.New(ccache.Configure())
	return Cache{
		client: cache,
		ttl:    config.TTL,
	}
}

// el email es una key en la cache ya que es unico. no van haber dos usuarios con el mismo email

func (repository Cache) GetUsuariobyEmail(email string) (dao.Usuario, error) {
	eKey := fmt.Sprintf("usuario:email:%s", email)

	item := repository.client.Get(eKey)

	if item != nil && !item.Expired() {
		user, ok := item.Value().(domain.UsuarioData)

		if !ok {
			return dao.Usuario{}, fmt.Errorf("Error al recuperar usuario de la cache")
		}

		userf := dao.Usuario{
			UsuarioID:    user.UsuarioID,
			Nombre:       user.Nombre,
			Apellido:     user.Apellido,
			Email:        user.Email,
			Passwordhash: user.Passwordhash,
			Tipo:         user.Tipo,
		}

		return userf, nil
	}

	return dao.Usuario{}, fmt.Errorf("Usuario no encontrado en la cache con email %s", email)
}

func (repository Cache) GetUsuariobyID(id int64) (dao.Usuario, error) {
	idKey := fmt.Sprintf("user:id:%d", id)

	item := repository.client.Get(idKey)

	if item != nil && !item.Expired() {

		user, ok := item.Value().(domain.UsuarioData)

		if !ok {
			return dao.Usuario{}, fmt.Errorf("Error al recuperar usuario de la cache")
		}

		userf := dao.Usuario{
			UsuarioID:    user.UsuarioID,
			Nombre:       user.Nombre,
			Apellido:     user.Apellido,
			Email:        user.Email,
			Passwordhash: user.Passwordhash,
			Tipo:         user.Tipo,
		}

		return userf, nil
	}

	return dao.Usuario{}, fmt.Errorf("Usuario no encontrado en la cache con email %d", id)
}

func (repository Cache) Actualizar(usuario dao.Usuario) error {

	idKey := fmt.Sprintf("user:id:%d", usuario.UsuarioID)
	eKey := fmt.Sprintf("user:email:%s", usuario.Email)

	repository.client.Set(idKey, usuario, repository.ttl)
	repository.client.Set(eKey, usuario, repository.ttl)

	return nil

}

func (repository Cache) CrearUsuario(newusuario dao.Usuario) (dao.Usuario, error) {

	idKey := fmt.Sprintf("user:id:%d", newusuario.UsuarioID)
	eKey := fmt.Sprintf("user:email:%s", newusuario.Email)

	repository.client.Set(idKey, newusuario, repository.ttl)
	repository.client.Set(eKey, newusuario, repository.ttl)

	return newusuario, nil

}
