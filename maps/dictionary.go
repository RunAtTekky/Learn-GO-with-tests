package maps

import "errors"

type Dictionary map[string]string

var ErrNoEntryFound = errors.New("No entry found for your word")

func (d Dictionary) Search(key string) (val string, err error) {
	definition, ok := d[key]

	if !ok {
		return "", ErrNoEntryFound
	}
	return definition, nil
}
