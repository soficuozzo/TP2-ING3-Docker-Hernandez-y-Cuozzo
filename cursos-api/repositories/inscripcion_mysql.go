package repositories

import (
	"context"
	"fmt"
	"proyecto_arqui_soft_2/cursos-api/dao"
	"proyecto_arqui_soft_2/cursos-api/domain"

	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

// GetCursoByName implements services.Repository.
func (m MySQL) GetCursoByName(ctx context.Context, name string) (dao.Curso, error) {
	panic("unimplemented")
}

// GetAllCursos implements services.Repository.
func (m MySQL) GetAllCursos(ctx context.Context) ([]domain.CursoData, error) {
	panic("unimplemented")
}

// GetCursosByIds implements services.Repository.
func (m MySQL) GetCursosByIds(tx context.Context, id []string) ([]domain.CursoData, error) {
	panic("unimplemented")
}

func (m MySQL) GetInscripcionByUserId(ctx context.Context, id int64) ([]dao.Inscripcion, error) {

	var inscripciones []dao.Inscripcion

	result := m.db.Where("usuario_id = ?", id).Find(&inscripciones)

	if result.Error != nil {
		return inscripciones, result.Error
	}

	return inscripciones, nil

}

// InscribirCurso implements services.Repository.
func (m MySQL) InscribirCurso(ctx context.Context, inscripcion dao.Inscripcion) error {
	result := m.db.Create(&inscripcion)

	if result.Error != nil {
		log.Error("")
		return result.Error
	}

	return nil
}

// Create implements services.Repository.
func (m MySQL) Create(ctx context.Context, curso dao.Curso) (string, error) {
	panic("unimplemented")
}

// Delete implements services.Repository.
func (m MySQL) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// GetCursoByID implements services.Repository.
func (m MySQL) GetCursoByID(ctx context.Context, id string) (dao.Curso, error) {
	panic("unimplemented")
}

// Update implements services.Repository.
func (m MySQL) Update(ctx context.Context, id string, updateData primitive.M) error {
	panic("unimplemented")
}

var (
	migrate = []interface{}{
		dao.Inscripcion{},
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
