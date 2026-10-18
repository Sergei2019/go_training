// // Задание
// // Напишите функцию olderThanDataClients, которая принимает дату bday и возвращает список из первых 10 клиентов старше указанной даты. Список должен быть отсортирован по убыванию дат рождения. Даты рождения представлены в формате YYYYMMDD. Каждое значение в списке должно содержать все поля из структуры ClientInfo. Функция main и структура VideoInfo уже реализованы. bday = 19630225.

// SELECT id, fio, email, login, birthday
// FROM clients
// WHERE birthday < ?
// ORDER BY birthday DESC
// LIMIT 10

package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

const bday = 19630225

type ClientInfo struct {
	ID       int
	FIO      string
	Email    string
	Login    string
	Birthday int
}

func olderThanDataClients(ctx context.Context, db *sql.DB, bday int) ([]ClientInfo, error) {
	// Готовим SQL запрос
	query := `SELECT id, fio, email, login, birthday 
            FROM clients
            WHERE birthday < ?
            ORDER BY birthday DESC
            LIMIT 10`

	// Выполняем запрос
	rows, err := db.QueryContext(ctx, query, bday)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []ClientInfo

	// Итерируем по результатам запроса
	for rows.Next() {
		var c ClientInfo
		if err := rows.Scan(&c.ID, &c.FIO, &c.Email, &c.Login, &c.Birthday); err != nil {
			return nil, err
		}
		clients = append(clients, c)
	}

	// Проверяем наличие ошибок после итерации
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return clients, nil
}

func main() {
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	clients, err := olderThanDataClients(context.Background(), db, bday)
	if err != nil {
		fmt.Println(err)
	}

	for _, c := range clients {
		fmt.Println(c)
	}
}
