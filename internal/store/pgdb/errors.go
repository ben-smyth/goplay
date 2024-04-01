package pgdb

import "fmt"

func NewSQLNotFound(entityType, name string) error {
	return fmt.Errorf("%s: %s not found", entityType, name)
}

func NewSQLDeleteFailed(entityType, name string, err error) error {
	return fmt.Errorf("%s: unable to delete %s: %s", entityType, name, err.Error())
}
func NewSQLUpdateFailed(entityType, name string, err error) error {
	return fmt.Errorf("%s: unable to update %s: %s", entityType, name, err.Error())
}

func NewSQLCreateFailed(entityType, name string, err error) error {
	return fmt.Errorf("%s: unable to create %s: %s", entityType, name, err.Error())
}

func NewSQLListFailed(entityType string, err error) error {
	return fmt.Errorf("%s: unable to list all: %s", entityType, err.Error())
}
