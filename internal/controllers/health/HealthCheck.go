package health

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"totmapi/internal/config"
	"totmapi/internal/controllers"

	_ "github.com/lib/pq"
)

func SetRoutes(router *mux.Router) {
	router.HandleFunc("/hello", Hello)
	router.HandleFunc("/database", DatabaseHealth)
}

func init() {
	controllers.RegisterRouteSetter(SetRoutes)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func DatabaseHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "checking db health")

	connStr := *config.GetDBConfig().ConnectionString

	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	// Call the function to query the 'person' table
	err = getAllPersons(db)
	if err != nil {
		log.Fatalf("Error querying 'person' table: %v", err)
	}
}

// getAllPersons queries the 'person' table and prints the results.
// Adjust the code to match your actual schema (column names/types).
func getAllPersons(db *sql.DB) error {
	rows, err := db.Query("SELECT * FROM totm.person")
	if err != nil {
		return fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	// Example: Suppose 'person' has columns: id (int), name (text)
	// Adjust the scanning to match your actual columns.
	for rows.Next() {
		var id int
		var name string

		err := rows.Scan(&id, &name)
		if err != nil {
			return fmt.Errorf("scan error: %w", err)
		}

		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("rows error: %w", err)
	}

	return nil
}
