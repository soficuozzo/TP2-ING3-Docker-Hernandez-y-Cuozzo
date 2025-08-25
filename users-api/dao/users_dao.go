package dao

type Usuario struct {
	UsuarioID    int64  `gorm:"primary_key;AUTO_INCREMENT; not null"`
	Nombre       string `gorm:"type:varchar(255)"`
	Apellido     string `gorm:"type:varchar(255)"`
	Email        string `gorm:"type:varchar(255)"`
	Tipo         string `gorm:"type:varchar(255)"`
	Passwordhash string `gorm:"type:varchar(255)"`
}

type Usuarios []Usuario
