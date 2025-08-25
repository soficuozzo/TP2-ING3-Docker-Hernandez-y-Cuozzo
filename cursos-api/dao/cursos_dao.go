package dao

type Curso struct {
	CursoID     string `bson:"_id,omitempty"`
	Nombre      string `bson:"nombre"`
	Descripcion string `bson:"descripcion"`
	Categoria   string `bson:"categoria"`
	Capacidad   int64  `bson:"capacidad"`
	Requisito   string  `bson:"requisito"`
	Duracion   int64  `bson:"duracion"`
	Imagen   string  `bson:"imagen"`
	Valoracion   int64  `bson:"valoracion"`
	Profesor   string  `bson:"profesor"`
}

type CursosData []Curso

type Inscripcion struct {
	UsuarioID int64  `json:"usuario_id"`
	CursoID   string `json:"curso_id"`
}

// Estructura para la respuesta de disponibilidad
type DisponibilidadCurso struct {
	CursoID        string `json:"curso_id"`
	Disponibilidad int64  `json:"disponibilidad"`
}

// type Cursos []Curso no se si va aca

type Inscripciones []Inscripcion