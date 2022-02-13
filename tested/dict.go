package tested

type Dictionary map[string]string

const (
	ErrKeyNotFound = DictErr("could not find the word")
	ErrKeyExists   = DictErr("Key already exists")
)

type DictErr string

func (e DictErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]

	if !ok {
		return "", ErrKeyNotFound
	}
	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrKeyNotFound:
		d[key] = value
	case nil:
		return ErrKeyExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrKeyNotFound:
		return ErrKeyNotFound
	case nil:
		d[key] = value
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)

	switch err {
	case ErrKeyNotFound:
		return ErrKeyNotFound
	case nil:
		delete(d, key)
	default:
		return err
	}
	return nil
}
