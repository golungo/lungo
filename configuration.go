package lungo

import (
	"regexp"
)

type DatabaseConfig struct {
	Name string
	URI  string
}

var configuration DatabaseConfig

func Init(URI string) error {
	re, err := regexp.Compile(`mongodb\:\/\/(?P<user>([^"]*))\:(?P<password>([^"]*))\@(?P<host>([^"]*))\:(?P<port>([^"]*))\/(?P<database>([^"]*))`)

	if err != nil {
		return err
	}

	matches := re.FindStringSubmatch(URI)
	names := re.SubexpNames()
	result := make(map[string]string)

	for i, name := range names {
		if name != "" {
			result[name] = matches[i]
		}
	}

	configuration = DatabaseConfig{Name: result["database"], URI: URI}

	return nil
}
