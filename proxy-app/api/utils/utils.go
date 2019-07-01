package utils

import (
	env "github.com/joho/godotenv"
)

// LoadEnv - loads environment
func LoadEnv() {
	env.Load("/Users/donovan/Documents/Wizeline/GolangProjects/go/src/github.com/kiraco/proxy-app/.env")
}
