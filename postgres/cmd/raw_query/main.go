package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/jackc/pgx/v4"
)

const (
	dbDSN = "host=localhost port=54321 dbname=auth user=auth-user password=auth-password sslmode=disable"
)

func main() {
	ctx := context.Background()

	// Создаем соединение с базой данных
	con, err := pgx.Connect(ctx, dbDSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer con.Close(ctx)

	// Делаем запрос на вставку записи в таблицу users
	res, err := con.Exec(ctx, "INSERT INTO users (name, email, password, role) VALUES ($1, $2, $3, $4)", gofakeit.Name(), gofakeit.Email(), "525252", "ADMIN") //знак доллара в качестве placeholder
		if err != nil {
		log.Fatalf("failed to insert users: %v", err)
	}

	log.Printf("inserted %d rows", res.RowsAffected())

	// Делаем запрос на выборку записей из таблицы users
	rows, err := con.Query(ctx, "SELECT * FROM users")
	if err != nil {
		log.Fatalf("failed to select users: %v", err)
	}
	defer rows.Close()

	// Итерируемся по всем строчкам данных и парсим их для вывода
	for rows.Next() {
		var id int
		var name, email, password, role string
		var createdAt time.Time
		var updatedAt sql.NullTime

		err = rows.Scan(&id, &name, &email, &password, &role, &createdAt, &updatedAt)
		if err != nil {
			log.Fatalf("failed to scan users: %v", err)
		}

		log.Printf("id: %d, name: %s, email: %s, password: %s, role: %s, created_at: %v, updated_at: %v\n", id, name, email, password, role, createdAt, updatedAt)
	}
}
