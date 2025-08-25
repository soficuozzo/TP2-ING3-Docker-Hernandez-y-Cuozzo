package services

import (
	"proyecto_arqui_soft_2/users-api/dao"
	"proyecto_arqui_soft_2/users-api/domain"
)

type Mock struct{}

func NewMock() Mock {
	return Mock{}
}

func (service Mock) GetUsuariobyEmail(email string) (dao.Usuario, error) {
	//TODO implement me
	panic("implement me")
}

func (service Mock) GetUsuariobyID(id int64) (dao.Usuario, error) {
	//TODO implement me
	panic("implement me")
}

func (service Mock) CrearUsuario(usuario dao.Usuario) (domain.UsuarioData, error) {
	//TODO implement me
	panic("implement me")

}

func (service Mock) Actualizar(usuario dao.Usuario) error {
	//TODO implement me
	panic("implement me")
}

func (service Mock) GenerarJWT(email string) (string, error) {
	//TODO implement me
	panic("implement me")
}
