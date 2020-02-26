package dict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound = errors.New("Not Found")
	errCantUpdate = errors.New("Can't update non-existing word")
	errWordExists = errors.New("The word already exists")
)



// Searching a word
func (d Dictionary) Search(word string) (string, error){
	value, exists := d[word]
	if exists{
		return value, nil
	}
	return "", errNotFound
}

// Add word to dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err{
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err{
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Deletes a word from dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}