package middleware

import "os"

const dbFile = "blockchain_%s.db"

func DbExists(dbFile string) bool {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return false
	}

	return true
}
