package main

// В пакете database/sql для выполнения SELECT-запросов есть функция Query() (англ. query — «запрос»). Она принимает SQL-запрос в виде строки и параметры запроса. Про параметры запроса мы подробно расскажем ниже, а пока взгляните на выполнение простых запросов.
// Вот как будет выглядеть фрагмент кода, который позволит сделать запрос и получить все записи из таблицы products: 

// запрашиваем только идентификатор, название и цену продукта
// rows, err := db.Query("SELECT id, product, price FROM products")
// if err != nil {
//     fmt.Println(err)
//     return
// }
// defer rows.Close()
package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "modernc.org/sqlite"
)

type Product struct {
    ID      int
    Product string
    Price   int
}

func (p Product) String() string {
    return fmt.Sprintf("ID: %d, Product: %s, Price: %d", p.ID, p.Product, p.Price)
}

func main() {
    db, err := sql.Open("sqlite", "demo.db")
    if err != nil {
        log.Println(err)
        return
    }
    defer db.Close()

    rows, err := db.Query("SELECT id, product, price FROM products")
    if err != nil {
        log.Println(err)
        return
    }
    defer rows.Close()

    for rows.Next() {
        product := Product{}

        err := rows.Scan(&product.ID, &product.Product, &product.Price)
        if err != nil {
            log.Println(err)
            return
        }

        fmt.Println(product)
    }

    if err := rows.Err(); err != nil {
        log.Println(err)
        return
    }
}