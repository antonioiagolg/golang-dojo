package main

const (
	ErrorUnkonwWord          = DictionaryErr("word not in dictionary")
	ErrorAlreadyExists       = DictionaryErr("word exists")
	ErrorWordDoesNotExist    = DictionaryErr("cannot update: word not in dictionary")
	ErrorDeleteWordNotExists = DictionaryErr("cannot delete: word does not exist")
)

type Dictionary map[string]string
type DictionaryErr string

func (d DictionaryErr) Error() string {
	return string(d)
}

func (d Dictionary) Search(word string) (string, error) {
	if _, ok := d[word]; !ok {
		return "", ErrorUnkonwWord
	}
	return d[word], nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorUnkonwWord:
		d[word] = definition
	case nil:
		return ErrorAlreadyExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definiton string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorUnkonwWord:
		return ErrorWordDoesNotExist
	case nil:
		d[word] = definiton
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorUnkonwWord:
		return ErrorDeleteWordNotExists
	case nil:
		delete(d, word)
	default:
		return err
	}
	return nil
}
