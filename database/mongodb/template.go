package mongodb

import (
	"context"
	"github.com/bpdlampung/banklampung-core-backend-go/helpers/structs"
	"go.mongodb.org/mongo-driver/bson"
)

type Templates interface {
	Save(obj interface{}, ctx context.Context) error
	Update(obj interface{}, ctx context.Context) error
	FindById(id string, result interface{}, ctx context.Context) error
}

type implementTemplateRepository struct {
	mongodb Collections
}

func ImplementTemplateRepository(mongoDb Collections) Templates {
	return implementTemplateRepository{
		mongodb: mongoDb,
	}
}

func (i implementTemplateRepository) Save(obj interface{}, ctx context.Context) error {
	if len(structs.GetStringValueFromStruct(obj, "bson", "_id")) == 0 {
		panic("_id is required")
	}

	err := i.mongodb.InsertOne(InsertOne{
		Document: structs.DoCreateMongoEntity(obj),
	}, ctx)

	if err != nil {
		return err
	}

	return nil
}

func (i implementTemplateRepository) Update(obj interface{}, ctx context.Context) error {
	objId := structs.GetStringValueFromStruct(obj, "bson", "_id")

	if len(objId) == 0 || objId == "-" {
		panic("_id is required ::" + objId)
	}

	err := i.mongodb.UpdateOne(UpdateOne{
		Filter: bson.M{
			"_id": structs.GetStringValueFromStruct(obj, "bson", "_id"),
		},
		Document: structs.DoUpdateMongoEntity(obj),
	}, ctx)

	if err != nil {
		return err
	}

	return nil
}

func (i implementTemplateRepository) FindById(id string, result interface{}, ctx context.Context) error {
	if err := i.mongodb.FindOne(FindOne{
		Result: result,
		Filter: bson.M{
			"_id": id,
		},
	}, ctx); err != nil {
		return err
	}

	return nil
}
