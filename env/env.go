package env

import (
	"log"
	"os"
	"strconv"
)

// It's common that you don't want a program to run if the environment is missing configuration
func GetenvOrFatal(environmentVariableName string) string {
	variable, ok := os.LookupEnv(environmentVariableName)
	if !ok {
		log.Fatalf("environment variable %s not found", environmentVariableName)
	}
	return variable
}

func GetenvOrFatalInt(environmentVariableName string) int {
	variableString := GetenvOrFatal(environmentVariableName)
	variableInt, err := strconv.Atoi(variableString)
	if err != nil {
		log.Fatalf("environment variable %s could not be interpreted as an int", environmentVariableName)
	}
	return variableInt
}

func GetenvOrFatalFloat(environmentVariableName string) float64 {
	variableString := GetenvOrFatal(environmentVariableName)
	variableFloat, err := strconv.ParseFloat(variableString, 64)
	if err != nil {
		log.Fatalf("environment variable %s could not be interpreted as a float64", environmentVariableName)
	}
	return variableFloat
}

func GetenvOrFatalBool(environmentVariableName string) bool {
	variableString := GetenvOrFatal(environmentVariableName)
	variableBool, err := strconv.ParseBool(variableString)
	if err != nil {
		log.Fatalf("environment variable %s could not be interpreted as a bool", environmentVariableName)
	}
	return variableBool
}
