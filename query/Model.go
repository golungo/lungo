package query

import (
	"reflect"
	"strings"

	"github.com/golungo/lungo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type virtualField struct {
	from         string
	localField   string
	foreignField string
	as           string
	single       bool
}

type Model struct {
	collection *mongo.Collection `json:"-" bson:"-"`
	refs       map[string]string `json:"-" bson:"-"`
	query      []bson.D          `json:"-" bson:"-"`
	types      map[string]string `json:"-" bson:"-"`
	errors     string            `json:"-" bson:"-"`
	virtuals   []virtualField    `json:"-" bson:"-"`
}

func (m Model) Init(collectionName string, types reflect.Type) Model {
	m.collection = lungo.GetCollection(collectionName)

	_types := map[string]string{}

	for i := 0; i < types.NumField(); i++ {
		_field := types.Field(i)

		fieldName := strings.ToLower(_field.Name)

		if fieldName == "id" {
			fieldName = "_id"
		}

		var typeName string

		switch _field.Type.Kind() {
		case reflect.Slice:
			typeName = "slice"
		default:
			typeName = _field.Type.Name()
		}

		_types[fieldName] = typeName
	}

	m.types = _types

	return m
}
