package dao

type Curso struct {
	CursoID     string `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Categoria   string `json:"categoria"`
	Capacidad   int64  `json:"capacidad"`
	Requisito   string `json:"requisito"`
	Duracion    int64  `json:"duracion"`
	Imagen      string `json:"imagen"`
	Valoracion  int64  `json:"valoracion"`
	Profesor    string `json:"profesor"`
}
