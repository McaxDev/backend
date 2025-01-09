package utils

import "os"

func GetEnv(key, defValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	} else {
		return defValue
	}
}
