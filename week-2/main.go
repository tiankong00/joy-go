package main

import (
	"database/sql"
	"log"

	"github.com/pkg/errors"
)

func DBQuery() error {
	return errors.Wrap(sql.ErrNoRows, "Not Found")
}

func DBConn() error {
	return errors.WithMessage(DBQuery(), "Conn Failed")
}

func DB() error {
	err := DBConn()
	if errors.Cause(err) == sql.ErrNoRows {
		log.Printf("%+v\n", err)
		return nil
	}
	return err
}

func main() {
	err := DB()
	if err != nil {
		log.Println(err)
	}
}
