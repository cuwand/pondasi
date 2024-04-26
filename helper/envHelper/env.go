package envHelper

import (
	"fmt"
	"os"
	"strconv"
)

func GetEnvAndValidateBool(key string) bool {
	value := os.Getenv(key)

	if len(value) == 0 {
		panic(fmt.Sprintf("env [%s] not found", key))
	}

	b, err := strconv.ParseBool(value)

	if err != nil {
		panic(err)
	}

	return b
}

func GetEnvAndValidateInt(key string) int {
	value := os.Getenv(key)

	if len(value) == 0 {
		panic(fmt.Sprintf("env [%s] not found", key))
	}

	b, err := strconv.Atoi(value)

	if err != nil {
		panic(err)
	}

	return b
}

func GetEnvAndValidate(key string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		panic(fmt.Sprintf("env [%s] not found", key))
	}

	return value
}

func GetEnv(key string) string {
	value := os.Getenv(key)

	return value
}
