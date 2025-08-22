package maps

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrNoEntryFound      = DictionaryErr("No entry found for your word")
	ErrWordExists        = DictionaryErr("Cannot ADD word as it already exists")
	ErrWordDoesNotExists = DictionaryErr("The word cannot be UPDATED as it does not exist")
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

func (d Dictionary) Update(key, new_value string) error {
	_, err := d.Search(key)

	switch err {
	case nil:
		d[key] = new_value
	case ErrNoEntryFound:
		return ErrWordDoesNotExists
	default:
		return err
	}

	return nil
}
