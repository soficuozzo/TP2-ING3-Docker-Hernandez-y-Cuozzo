package repositories

import (
	dao "proyecto_arqui_soft_2/search-api/dao"
)

type Mock struct {
	data map[int64]dao.Curso
}

func NewMock() Mock {
	return Mock{
		data: make(map[int64]dao.Curso),
	}
}