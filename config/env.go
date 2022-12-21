package config

import "os"

var (
	DatabaseTable = os.Getenv("DATABASE_TABLE")
	Region        = os.Getenv("REGION")
)
