package query

import (
	"github.com/golungo/lungo"

	"go.mongodb.org/mongo-driver/bson"
)

func (m Model) Find(filter lungo.Filter) Model {
	if m.ifError() {
		return m
	}

	if filter != nil {
		var match bson.D

		for fieldName, fieldValue := range filter {
			if _type, exists := m.types[fieldName]; exists {
				switch _type {
				case "slice":
					match = append(match, bson.E{
						Key: fieldName,
						Value: bson.D{
							bson.E{
								Key:   "$in",
								Value: bson.A{fieldValue},
							},
						},
					})
				default:
					match = append(match, bson.E{
						Key:   fieldName,
						Value: fieldValue,
					})
				}
			}
		}

		result := bson.D{
			bson.E{
				Key:   "$match",
				Value: match,
			},
		}

		m.query = append(m.query, result)
	}

	return m
}
