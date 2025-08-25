package repositories

import (
	//"errors"
	"fmt"
	"log"
	"proyecto_arqui_soft_2/users-api/dao"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLConfig struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

type MySQL struct {
	db *gorm.DB
}

// GenerarJWT implements services.Repository.
func (repository MySQL) GenerarJWT(email string) (string, error) {
	panic("unimplemented")
}

// Actualizar implements services.Repository.
func (repository MySQL) Actualizar(usuario dao.Usuario) error {
	if err := repository.db.Save(&usuario).Error; err != nil {
		return fmt.Errorf("error actualizando usuario: %w", err)
	}
	return nil
}

var (
	migrate = []interface{}{
		dao.Usuario{},
	}
)

func NewMySQL(config MySQLConfig) MySQL {
	// Build DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username, config.Password, config.Host, config.Port, config.Database)

	// Open connection to MySQL using GORM
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to MySQL: %s", err.Error())
	}

	// Automigrate structs to Gorm
	for _, target := range migrate {
		if err := db.AutoMigrate(target); err != nil {
			log.Fatalf("error automigrating structs: %s", err.Error())
		}
	}

	return MySQL{
		db: db,
	}
}

func (repository MySQL) GetUsuariobyEmail(email string) (dao.Usuario, error) {
	var usuario dao.Usuario

	result := repository.db.Where("email = ?", email).First(&usuario)
	if result.Error != nil {
		return usuario, result.Error

	}

	return usuario, nil
}

func (repository MySQL) GetUsuariobyID(id int64) (dao.Usuario, error) {
	var usuario dao.Usuario

	result := repository.db.Where("usuario_id = ?", id).First(&usuario)
	if result.Error != nil {
		return usuario, result.Error
	}

	return usuario, nil
}

func (repository MySQL) CrearUsuario(newusuario dao.Usuario) (dao.Usuario, error) {

	result := repository.db.Create(&newusuario)

	if result.Error != nil {
		return newusuario, fmt.Errorf("error creating user: %w", result.Error)
	}

	return newusuario, nil
}
