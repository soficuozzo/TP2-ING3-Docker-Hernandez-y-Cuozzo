package domain

type CursoData struct {
	CursoID     string `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Categoria   string `json:"categoria"`
	Capacidad   int64  `json:"capacidad"`
	Imagen      string `json:"imagen,omitempty"`
	Valoracion  int64  `json:"valoracion"`
	Requisito   string `json:"requisito"`
	Profesor    string `json:"profesor"`
	Duracion    int64  `json:"duracion"`
}
type CursoNew struct {
	Operation string `json:"operation"`
	CursoID   string `json:"curso_id"`
}
