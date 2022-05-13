package util

import (
	"context"
	"fmt"
	"vidlearn-final-projcect/config"

	"github.com/labstack/gommon/log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseDriver string

const (
	// MySQL
	MySQl DatabaseDriver = "mysql"

	// MongoDB
	MongoDB DatabaseDriver = "mongodb"
)

type DatabaseConnection struct {
	Driver DatabaseDriver

	// MySQL
	MySQlDB *gorm.DB

	// MongoDB
	MongoDB     *mongo.Database
	mongoClient *mongo.Client
}

func CreateDatabaseConnection(config *config.AppConfig) *DatabaseConnection {
	var db DatabaseConnection

	switch config.Database.Driver {
	case "mysql":
		db.MySQlDB = CreateMySQLClient(config)
		db.Driver = MySQl
	case "mongodb":
		db.mongoClient = CreateMongoClient(config)
		db.MongoDB = db.mongoClient.Database(config.Log.Name)
		db.Driver = MongoDB
	default:
		panic("Database driver not supported")
	}

	return &db
}

func CreateLogConnection(config *config.AppConfig) *DatabaseConnection {
	var db DatabaseConnection

	switch config.Log.Driver {
	case "mongodb":
		db.mongoClient = CreateMongoClient(config)
		db.MongoDB = db.mongoClient.Database(config.Log.Name)
		db.Driver = MongoDB
	default:
		panic("Database driver not supported")
	}

	return &db
}

func (db *DatabaseConnection) CloseConnection() {
	if db.MySQlDB != nil {
		db, _ := db.MySQlDB.DB()
		db.Close()
	}

	ctx := context.Background()
	if db.mongoClient != nil {
		db.mongoClient.Disconnect(ctx)
	}
}

func CreateMySQLClient(config *config.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Info("Error connecting to database")
		panic(err)
	}

	return db
}

func CreateMongoClient(config *config.AppConfig) *mongo.Client {
	uri := fmt.Sprintf("mongodb://%v:%v",
		config.Log.Host,
		config.Log.Port,
	)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Info("Error connecting to database")
		panic(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Info("Error connecting to database")
		panic(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Info("Error connecting to database")
		panic(err)
	}

	return client
}
