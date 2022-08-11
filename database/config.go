package database

import "fmt"

var (
	host     = "localhost"
	user     = ""
	password = ""
	name     = "realcargo"
	port     = 5432
)

// Configdan foydalanib "connection string" yasab olamiz
var DB_CONFIG = fmt.Sprintf(
	"host=%s user=%s password=%s dbname=%s port=%d",
	host, user, password, name, port,
)
