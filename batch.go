package main

import (
	"banwire/services/file_tokenizer/db"
)

// BatchTest is a function only for test
func BatchTest() {

	defer func() {
		db.Connection.Close(nil)
	}()

	/*
		bla bla bla ... more code
	*/
}
