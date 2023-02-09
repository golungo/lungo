package lungo

import (
	"reflect"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetCollection(name string) *mongo.Collection {
	collection := connection.Database(configuration.Name).Collection(name)

	return collection
}

func GetRefs(t reflect.Type) map[string]string {
	result := map[string]string{}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldName := field.Tag.Get("bson")
		ref := field.Tag.Get("lungo")

		if ref != "" {
			result[fieldName] = ref
		}
	}

	return result
}
