package dict

import "errors"

type Dictionary map[string]string // type alias

var errNotFound = errors.New("no word found")
var errKeyExists = errors.New("the word already exists")
var errKeyDoesNotExist = errors.New("the word does not exists")

func (d Dictionary) Search(key string) (string, error) {
	value, exists := d[key]
	if exists {
		return value, nil
	}

	return "", errNotFound
}

func (d Dictionary) Add(key string, value string) error {
	_, err := d.Search(key)

	switch err {
	case errNotFound:
		d[key] = value
	case nil:
		return errKeyExists
	}

	return nil
}

func (d Dictionary) Update(key string, value string) error {
	_, err := d.Search(key)

	switch err {
	case errNotFound:
		return errKeyDoesNotExist
	case nil:
		d[key] = value
	}

	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
