package mongodb

import (
	"context"
	"fmt"
	"github.com/cuwand/pondasi/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
	"strings"
)

type Mongodb struct {
	client *mongo.Client
	dbname string
	logger logger.Logger
}

var mongoClient Mongodb

func GenerateUri(host, port, name, username, password string) string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", username, password, host, port, name)
}

func InitConnection(masterDBUrl string, logger logger.Logger) Mongodb {
	mClient, err := newClient(masterDBUrl)

	if err != nil {
		panic(err)
	}

	session, _ := mClient.StartSession()
	session.StartTransaction()
	mongo.WithSession(context.Background(), session, func(sessionContext mongo.SessionContext) error {

		return nil
	})

	mongoClient = Mongodb{
		client: mClient,
		dbname: strings.Split(strings.ReplaceAll(masterDBUrl, "//", ""), "/")[1],
		logger: logger,
	}

	logger.Info("MongoDB Connected")

	return mongoClient
}

func newClient(mongoUri string) (*mongo.Client, error) {
	tM := reflect.TypeOf(bson.M{})
	reg := bson.NewRegistryBuilder().RegisterTypeMapEntry(bsontype.EmbeddedDocument, tM).Build()

	client, err := mongo.Connect(
		context.Background(),
		options.Client().SetRetryWrites(true),
		options.Client().SetRetryReads(true),
		options.Client().SetMaxPoolSize(1000),
		options.Client().ApplyURI(mongoUri),
		options.Client().SetRegistry(reg),
	)

	if err != nil {
		return nil, err
	}

	//if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
	//	return nil, err
	//}

	return client, nil
}

func (r Mongodb) GetMongoClient() *mongo.Client {
	return r.client
}

func (r Mongodb) GetMongoLogger() logger.Logger {
	return r.logger
}
