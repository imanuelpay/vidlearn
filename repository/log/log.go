package log

import (
	"context"
	"vidlearn-final-projcect/business/log"
	"vidlearn-final-projcect/util"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func LogRepository(dbCon *util.DatabaseConnection) log.Repository {
	var logRepository log.Repository
	logRepository = CreateMongoDBRepository(dbCon.MongoDB)
	return logRepository
}

type MongoDBRepository struct {
	col *mongo.Collection
}

type collection struct {
	ID        primitive.ObjectID `bson:"_id"`
	At        string             `bson:"at"`
	Host      string             `bson:"host"`
	Method    string             `bson:"method"`
	Path      string             `bson:"path"`
	IP        string             `bson:"ip"`
	Status    int                `bson:"status"`
	UserAgent string             `bson:"user_agent"`
}

func newCollection(log log.Log) (*collection, error) {
	UIDString := primitive.NewObjectID().Hex()
	objectID, err := primitive.ObjectIDFromHex(UIDString)

	if err != nil {
		return nil, err
	}

	return &collection{
		objectID,
		log.At,
		log.Host,
		log.Method,
		log.Path,
		log.IP,
		log.Status,
		log.UserAgent,
	}, nil
}

func (col *collection) MongoDBToContent() *log.Log {
	var log log.Log
	log.ID = col.ID.Hex()
	log.At = col.At
	log.Host = col.Host
	log.Method = col.Method
	log.Path = col.Path
	log.IP = col.IP
	log.Status = col.Status
	log.UserAgent = col.UserAgent

	return &log
}

func CreateMongoDBRepository(db *mongo.Database) *MongoDBRepository {
	return &MongoDBRepository{
		db.Collection("log"),
	}
}

func (repository *MongoDBRepository) GetAllLog() (logs []*log.Log, err error) {
	cursor, err := repository.col.Find(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var col *collection
		err := cursor.Decode(&col)
		if err != nil {
			return nil, err
		}

		logs = append(logs, col.MongoDBToContent())
	}

	return logs, nil
}

func (repository *MongoDBRepository) CreateLog(log *log.Log) (*log.Log, error) {
	col, err := newCollection(*log)
	if err != nil {
		return nil, err
	}

	_, err = repository.col.InsertOne(context.TODO(), col)
	if err != nil {
		return nil, err
	}

	return log, nil
}
