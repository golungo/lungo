package query

import (
	"github.com/golungo/lungo"

	"go.mongodb.org/mongo-driver/bson"
)

func (m Model) Lookup(fields lungo.Fields) Model {
	if m.ifError() {
		return m
	}

	for _, fieldName := range fields {
		for _, v := range m.virtuals {
			if v.as == fieldName {
				req := bson.D{
					bson.E{
						Key: "$lookup",
						Value: bson.D{
							bson.E{
								Key:   "from",
								Value: v.from,
							},
							bson.E{
								Key:   "localField",
								Value: v.localField,
							},
							bson.E{
								Key:   "foreignField",
								Value: v.foreignField,
							},
							bson.E{
								Key:   "as",
								Value: v.as,
							},
						},
					},
				}

				m.query = append(m.query, req)
			}
		}
	}

	return m
}
