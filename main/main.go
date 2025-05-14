package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"go_server/database"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Dex string `json:"dex"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		http.Error(w, "Błąd podczas pobierania użytkowników", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			http.Error(w, "Błąd przy skanowaniu danych", http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func main() {
	db.InitDatabase()
	defer db.DB.Close()

	http.HandleFunc("/users", getUsers)

	fmt.Println("Serwer działa na porcie 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
