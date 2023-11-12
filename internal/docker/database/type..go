package database

var allowedTypes = []string{"mysql", "postgres", "mongo"}

var envBasedType = map[string]string{
	"mysql":    "MYSQL_ROOT_PASSWORD=my-secret-pw",
	"postgres": "POSTGRES_PASSWORD=my-secret-pw",
	"mongo":    "",
}

func typeExists(typeDB string) bool {
	for _, t := range allowedTypes {
		if t == typeDB {
			return true
		}
	}

	return false
}
