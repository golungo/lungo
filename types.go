package lungo

import "go.mongodb.org/mongo-driver/bson/primitive"

type Refs = map[string]string
type Filter = map[string]interface{}
type Fields = []string
type ObjectID = primitive.ObjectID
