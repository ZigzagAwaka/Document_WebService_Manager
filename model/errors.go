package model

import "fmt"

type ErrDocNotFound int

func (e ErrDocNotFound) Error() string {
	return fmt.Sprintf("The requested document with ID %v was not found", int(e))
}

type ErrDocAlreadyExists int

func (e ErrDocAlreadyExists) Error() string {
	return fmt.Sprintf("The document with ID %v already exists", int(e))
}
