package query

func (m Model) setError(message string) Model {
	m.errors = message

	return m
}

func (m Model) ifError() bool {
	return m.errors != ""
}
