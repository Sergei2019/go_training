// Задание
// Представьте, что вы обновляете информацию о клиентах. Если во время обновления происходит ошибка, нужно отменить все изменения, чтобы не нарушить целостность данных. В функцию updateClients добавьте использование транзакций:
// Функция должна начинаться с транзакции с уровнем изоляции LevelSerializable.
// В случае успешного обновления, функция подтверждает транзакцию.
// При любой ошибке функция должна отменять транзакцию с помощью и возвращать ошибку.

package main

import (
	"context"
	"database/sql"
)

type Client struct {
	ID       int
	FIO      string
	Login    string
	Birthday string
	Email    string
}

func updateClients(ctx context.Context, db *sql.DB, clients []Client) error {
	// начинаем транзакцию
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}

	// для каждого клиента обновляем персональные данные
	for _, client := range clients {
		_, err := tx.ExecContext(ctx,
			"UPDATE clients SET fio = :fio, login = :login, birthday = :birthday, email = :email WHERE id = :id",
			sql.Named("id", client.ID),
			sql.Named("fio", client.FIO),
			sql.Named("login", client.Login),
			sql.Named("birthday", client.Birthday),
			sql.Named("email", client.Email),
		)
		if err != nil {
			// если произошла ошибка, отменяем транзакцию
			tx.Rollback()
			return err
		}
	}

	// подтверждаем транзакцию
	return tx.Commit()
}
