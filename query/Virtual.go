package query

func (m Model) Virtual(fromCollection string, localField string, foreignField string, as string, single bool) Model {
	virtual := virtualField{
		from:         fromCollection,
		localField:   localField,
		foreignField: foreignField,
		as:           as,
		single:       single,
	}

	m.virtuals = append(m.virtuals, virtual)

	return m
}
