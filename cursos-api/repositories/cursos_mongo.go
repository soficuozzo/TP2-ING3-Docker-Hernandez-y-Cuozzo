package repositories

import (
	"context"
	"fmt"
	"log"
	cursosDAO "proyecto_arqui_soft_2/cursos-api/dao"
	"proyecto_arqui_soft_2/cursos-api/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	//agregue yo para inscripcion
	"proyecto_arqui_soft_2/cursos-api/dao"
)

// lo agregue para inscripcion
type CursosRepository interface {
	ObtenerCursoPorID(cursoID string) (dao.Curso, error)
	ActualizarCurso(curso dao.Curso) error
}

type CursosRepositoryImpl struct{}

func NewCursosRepository() CursosRepository {
	return &CursosRepositoryImpl{}
}

// Definición de la interfaz InscripcionRepository
type InscripcionRepository interface {
	Inscribir(ctx context.Context, cursoID string) error
	// Otros métodos según tus necesidades
}

// -----------------------------------------------------------------------------------------
type MongoConfig struct {
	Host       string
	Port       string
	Username   string
	Password   string
	Database   string
	Collection string
}

type Mongo struct {
	client     *mongo.Client
	database   string
	collection string
}

// GetInscripcionByUserId implements services.Repository.
func (repository Mongo) GetInscripcionByUserId(ctx context.Context, userid int64) ([]dao.Inscripcion, error) {
	panic("unimplemented")
}

// InscribirCurso implements services.Repository.
func (repository Mongo) InscribirCurso(ctx context.Context, inscripcion dao.Inscripcion) error {
	panic("unimplemented")
}

const (
	connectionURI = "mongodb://%s:%s"
)

func NewMongo(config MongoConfig) Mongo {
	credentials := options.Credential{
		Username: config.Username,
		Password: config.Password,
	}

	ctx := context.Background()
	uri := fmt.Sprintf(connectionURI, config.Host, config.Port)
	cfg := options.Client().ApplyURI(uri).SetAuth(credentials)

	client, err := mongo.Connect(ctx, cfg)
	if err != nil {
		log.Panicf("error connecting to MongoDB: %v", err)
	}

	return Mongo{
		client:     client,
		database:   config.Database,
		collection: config.Collection,
	}
}

func (repository Mongo) GetCursoByID(ctx context.Context, id string) (cursosDAO.Curso, error) {
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return cursosDAO.Curso{}, fmt.Errorf("error converting id to mongo ID: %w", err)
	}
	result := repository.client.Database(repository.database).Collection(repository.collection).FindOne(ctx, bson.M{"_id": objectID})
	if result.Err() != nil {
		return cursosDAO.Curso{}, fmt.Errorf("error finding document: %w", result.Err())
	}

	var curso cursosDAO.Curso
	if err := result.Decode(&curso); err != nil {
		return cursosDAO.Curso{}, fmt.Errorf("error decoding result: %w", err)
	}
	return curso, nil
}

func (repository Mongo) GetCursoByName(ctx context.Context, name string) (cursosDAO.Curso, error) {

	result := repository.client.Database(repository.database).Collection(repository.collection).FindOne(ctx, bson.M{"nombre": name})
	if result.Err() != nil {
		return cursosDAO.Curso{}, fmt.Errorf("error finding document: %w", result.Err())
	}

	var curso cursosDAO.Curso
	if err := result.Decode(&curso); err != nil {
		return cursosDAO.Curso{}, fmt.Errorf("error decoding result: %w", err)
	}
	return curso, nil
}

func (repository Mongo) GetCursosByIds(ctx context.Context, ids []string) ([]domain.CursoData, error) {
	// Convertir cada ID en ObjectID
	objectIDs := make([]primitive.ObjectID, 0, len(ids))
	for _, id := range ids {
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, fmt.Errorf("error converting id to mongo ID: %w", err)
		}
		objectIDs = append(objectIDs, objectID)
	}

	fmt.Println("ObjectIDs:", objectIDs) // Verifica los ObjectIDs

	// Crear filtro para encontrar múltiples documentos
	filter := bson.M{"_id": bson.M{"$in": objectIDs}}

	// Ejecutar la consulta
	cursor, err := repository.client.Database(repository.database).
		Collection(repository.collection).Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("error finding documents: %w", err)
	}
	defer cursor.Close(ctx)

	// Decodificar resultados
	var cursos []domain.CursoData
	for cursor.Next(ctx) {
		var curso struct {
			CursoID     primitive.ObjectID `bson:"_id"`
			Nombre      string             `bson:"nombre"`
			Descripcion string             `bson:"descripcion"`
			Categoria   string             `bson:"categoria"`
			Capacidad   int64              `bson:"capacidad"`
			Requisito   string             `bson:"requisito"`
			Duracion    int64              `bson:"duracion"`
			Imagen      string             `bson:"imagen"`
			Valoracion  int64              `bson:"valoracion"`
			Profesor    string             `bson:"profesor"`
		}

		if err := cursor.Decode(&curso); err != nil {
			return nil, fmt.Errorf("error decoding result: %w", err)
		}
		fmt.Println("Curso decodificado:", curso) // Verifica el contenido decodificado

		// Convertir el ObjectID a string para el dominio
		cursoData := domain.CursoData{
			CursoID:     curso.CursoID.Hex(), // Convertimos el ObjectID a string
			Nombre:      curso.Nombre,
			Descripcion: curso.Descripcion,
			Categoria:   curso.Categoria,
			Capacidad:   curso.Capacidad,
			Requisito:   curso.Requisito,
			Duracion:    curso.Duracion,
			Imagen:      curso.Imagen,
			Valoracion:  curso.Valoracion,
			Profesor:    curso.Profesor,
		}

		cursos = append(cursos, cursoData)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return cursos, nil
}

func (repository Mongo) Create(ctx context.Context, curso cursosDAO.Curso) (string, error) {
	result, err := repository.client.Database(repository.database).Collection(repository.collection).InsertOne(ctx, curso)
	if err != nil {
		return "", fmt.Errorf("error creating document: %w", err)
	}

	objectID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("error converting mongo ID to object ID")
	}
	return objectID.Hex(), nil
}

func (repository Mongo) Update(ctx context.Context, cursoID string, updateData bson.M) error {
	objectID, err := primitive.ObjectIDFromHex(cursoID)
	if err != nil {
		return fmt.Errorf("error converting id to mongo ID: %w", err)
	}

	if len(updateData) == 0 {
		return fmt.Errorf("no fields to update for course ID %s", cursoID)
	}

	filter := bson.M{"_id": objectID}
	result, err := repository.client.Database(repository.database).Collection(repository.collection).UpdateOne(ctx, filter, bson.M{"$set": updateData})
	if err != nil {
		return fmt.Errorf("error updating document: %w", err)
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("no document found with ID %s", cursoID)
	}

	return nil
}

func (repository Mongo) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("error converting id to mongo ID: %w", err)
	}

	filter := bson.M{"_id": objectID}
	result, err := repository.client.Database(repository.database).Collection(repository.collection).DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("error deleting document: %w", err)
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("no document found with ID %s", id)
	}

	return nil
}

func (repository Mongo) TestConnection(ctx context.Context) error {
	if err := repository.client.Ping(ctx, nil); err != nil {
		return fmt.Errorf("error al hacer ping a MongoDB: %w", err)
	}
	return nil
}

//lo agregue para inscripcion

// Obtiene un curso por su ID
func (repo *CursosRepositoryImpl) ObtenerCursoPorID(cursoID string) (dao.Curso, error) {
	// Simulamos la obtención de un curso desde la base de datos
	return dao.Curso{
		CursoID:     cursoID,
		Nombre:      "Curso de Ejemplo",
		Descripcion: "Descripción del curso",
		Categoria:   "Tecnología",
		Capacidad:   10,
	}, nil
}

// Actualiza un curso en la base de datos
func (repo *CursosRepositoryImpl) ActualizarCurso(curso dao.Curso) error {
	// Simulamos la actualización exitosa del curso en la base de datos
	return nil
}

func (repository Mongo) GetAllCursos(ctx context.Context) ([]domain.CursoData, error) {
	// Crear un filtro vacío para obtener todos los cursos
	filter := bson.D{} // Este filtro selecciona todos los documentos

	// Ejecutar la consulta para obtener todos los documentos
	cursor, err := repository.client.Database(repository.database).
		Collection(repository.collection).Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("error finding documents: %w", err)
	}
	defer cursor.Close(ctx)

	// Decodificar los resultados
	var cursos []domain.CursoData
	for cursor.Next(ctx) {
		var curso struct {
			CursoID     primitive.ObjectID `bson:"_id"`
			Nombre      string             `bson:"nombre"`
			Descripcion string             `bson:"descripcion"`
			Categoria   string             `bson:"categoria"`
			Capacidad   int64              `bson:"capacidad"`
			Requisito   string             `bson:"requisito"`
			Duracion    int64              `bson:"duracion"`
			Imagen      string             `bson:"imagen"`
			Valoracion  int64              `bson:"valoracion"`
			Profesor    string             `bson:"profesor"`
		}

		if err := cursor.Decode(&curso); err != nil {
			return nil, fmt.Errorf("error decoding result: %w", err)
		}

		// Convertir el ObjectID a string para el dominio
		cursoData := domain.CursoData{
			CursoID:     curso.CursoID.Hex(), // Convertimos el ObjectID a string
			Nombre:      curso.Nombre,
			Descripcion: curso.Descripcion,
			Categoria:   curso.Categoria,
			Capacidad:   curso.Capacidad,
			Requisito:   curso.Requisito,
			Duracion:    curso.Duracion,
			Imagen:      curso.Imagen,
			Valoracion:  curso.Valoracion,
			Profesor:    curso.Profesor,
		}

		cursos = append(cursos, cursoData)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return cursos, nil
}
