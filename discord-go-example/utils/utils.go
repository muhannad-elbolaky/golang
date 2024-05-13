package utils

import (
	"log"
	"os"
)

func GetEnvVar(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatal(key + " must be set")
	}
	return value
}
