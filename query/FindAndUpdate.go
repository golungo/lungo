package query

import (
	"context"
	"time"

	"github.com/golungo/lungo"
)

func (m Model) FindAndUpdate(filter lungo.Filter, update interface{}) Model {
	if m.ifError() {
		return m
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := m.collection.ReplaceOne(ctx, filter, update)
	if err != nil {
		return m.setError(err.Error())
	}

	return m.Find(filter)
}
