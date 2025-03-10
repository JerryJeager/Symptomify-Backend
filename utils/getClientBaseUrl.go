package utils

import "os"

func GetClientBaseUrl() string {
	environment := os.Getenv("ENVIRONMENT")
	var baseUrl string
	if environment == "development" {
		baseUrl = "http://localhost:3000"
	} else {
		baseUrl = "https://symptomify.vercel.app"
	}

	return baseUrl
}
