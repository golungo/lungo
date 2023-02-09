package query

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (m Model) Exec() ([]byte, error) {
	if m.ifError() {
		fmt.Println(m.errors)
		return nil, errors.New(m.errors)
	}

	collection := m.collection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pipeline := bson.A{}

	for _, query := range m.query {
		pipeline = append(pipeline, query)
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}

	if err := cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	data, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	return data, nil
}
