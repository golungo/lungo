package query

import (
	"github.com/golungo/lungo"

	"go.mongodb.org/mongo-driver/bson"
)

func (m Model) Sort(filter lungo.Filter) Model {
	if m.ifError() {
		return m
	}

	if filter != nil {
		var match bson.D

		for fieldName, fieldValue := range filter {
			match = append(match, bson.E{fieldName, fieldValue})
		}

		result := bson.D{{"$sort", match}}

		m.query = append(m.query, result)
	}

	return m
}
