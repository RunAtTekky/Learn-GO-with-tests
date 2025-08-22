package maps

import "errors"

type Dictionary map[string]string

var (
	ErrNoEntryFound = errors.New("No entry found for your word")
	ErrWordExists   = errors.New("The word you are trying to add already exists")
)

func (d Dictionary) Search(key string) (val string, err error) {
	definition, ok := d[key]

	if !ok {
		return "", ErrNoEntryFound
	}
	return definition, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNoEntryFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}
