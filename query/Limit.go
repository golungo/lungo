package query

import (
	"go.mongodb.org/mongo-driver/bson"
)

func (m Model) Limit(number int) Model {
	if m.ifError() {
		return m
	}

	if number != 0 {
		m.query = append(m.query, bson.D{
			bson.E{
				Key:   "$limit",
				Value: number,
			},
		})
	}

	return m
}
