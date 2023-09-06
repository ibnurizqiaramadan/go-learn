package Functions

import (
	"strconv"

	// "github.com/go-errors/errors"
	"github.com/gofiber/fiber/v2/log"
)

func StoI(s string) int {
	result, err := strconv.Atoi(s)
	// log.Error(err.(*errors.Error).Unwrap().Error())
	if err != nil {
		log.Error(err)
		return -0
	}
	return result
}

func StoB(s string) bool {
	result, error := strconv.ParseBool(s)
	if error != nil {
		log.Error(error)
		return false
	}
	return result
}

func IsEmpty(value any, defaultVal any) any {
	if value == nil || value == "" {
		return defaultVal
	}
	return value
}
