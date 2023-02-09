package query

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (m Model) Delete(ID interface{}) Model {
	if m.ifError() {
		return m
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := m.collection.DeleteOne(ctx, bson.M{"_id": ID})
	if err != nil {
		m.setError(err.Error())
	}

	return m
}
