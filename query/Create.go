package query

import (
	"context"
	"time"

	"github.com/golungo/lungo"
)

func (m Model) Create(data interface{}) Model {
	if m.ifError() {
		return m
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := m.collection.InsertOne(ctx, data)
	if err != nil {
		return m.setError(err.Error())
	}

	filter := lungo.Filter{
		"_id": result.InsertedID,
	}

	return m.Match(filter)
}
