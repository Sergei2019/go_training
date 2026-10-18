// Задание
// Реализуйте сервис отслеживания посылок со следующими функциями:
// регистрация посылки,
// получение списка посылок клиента,
// изменение статуса посылки,
// изменение адреса доставки,
// удаление посылки.
// План такой: информация о посылках хранится в БД. Посылка может быть зарегистрирована, отправлена или доставлена. При регистрации посылки создаётся новая запись в БД. У только что зарегистрированной должен быть статус «зарегистрирована». Трек-номер посылки равен её идентификатору в таблице. Если посылка в статусе «зарегистрирована», можно изменить адрес доставки или удалить посылку.
// После создания сервиса напишите тесты для проверки функций, работающих с БД.
// Структура программы
// Для вас уже подготовлен каркас приложения и БД. В функции main() проверяется основная функциональность сервиса. Структура ParcelService реализует логику работы с посылками и использует объект типа ParcelStore для работы с данными о посылке в БД.
// В качестве СУБД используется SQLite. Файл с БД называется tracker.db. В БД всего одна таблица parcel со следующими колонками:
// number — номер посылки, целое число, автоинкрементное поле.
// client — идентификатор клиента, целое число.
// status — статус посылки, строка.
// address — адрес посылки, строка.
// created_at — дата и время создания посылки, строка.
// Вам нужно реализовать пустые функции в файле parcel.go. Внимательно изучите, где и как используются функции, и комментарии к ним. Также в файле main.go объявлены три константы для трёх статусов посылки, которые вам тоже пригодятся:
// ParcelStatusRegistered имеет значение «registered», соответствует статусу посылки «зарегистрирована».
// ParcelStatusSent имеет значение «sent», соответствует статусу «отправлена».
// ParcelStatusDelivered имеет значение «delivered», соответствует статусу «доставлена».
// В функции main() необходимо создать подключение к БД и объекту типа ParcelStore.
// В файле parcel_test.go находятся тесты для функций из parcel.go. Реализуйте их, чтобы проверить, правильно ли работают запросы. Эти тесты интеграционные, и значит, понадобится подключение к БД для их запуска. Изучите описание теста и комментарии в нём — это поможет вам понять, что и как он тестирует.
// Чтобы убедиться, что вы всё сделали правильно, запустите написанные вами тесты — они должны выполниться без ошибок.
// А потом запустите приложение. Приложение тоже должно выполниться без ошибок, а на консоли будут примерно следующие сообщения:

// Список посылок может быть больше, если вы запускали приложение несколько раз.
// Репозиторий с прекодом находится здесь.
// Когда решите, что вы выполнили задание, нажмите на кнопку Готово ниже!

package main

import (
	"database/sql"
)

type ParcelStore struct {
	db *sql.DB
}

func NewParcelStore(db *sql.DB) ParcelStore {
	return ParcelStore{db: db}
}

func (s ParcelStore) Add(p Parcel) (int, error) {
	res, err := s.db.Exec("INSERT INTO parcel (client, status, address, created_at) VALUES (:client, :status, :address, :created_at)",
		sql.Named("client", p.Client),
		sql.Named("status", p.Status),
		sql.Named("address", p.Address),
		sql.Named("created_at", p.CreatedAt))
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s ParcelStore) Get(number int) (Parcel, error) {
	p := Parcel{}

	row := s.db.QueryRow("SELECT number, client, status, address, created_at FROM parcel WHERE number = :number",
		sql.Named("number", number))
	err := row.Scan(&p.Number, &p.Client, &p.Status, &p.Address, &p.CreatedAt)
	if err != nil {
		return p, err
	}

	return p, nil
}

func (s ParcelStore) GetByClient(client int) ([]Parcel, error) {
	rows, err := s.db.Query("SELECT number, client, status, address, created_at FROM parcel WHERE client = :client",
		sql.Named("client", client))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []Parcel
	for rows.Next() {
		p := Parcel{}

		err := rows.Scan(&p.Number, &p.Client, &p.Status, &p.Address, &p.CreatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

func (s ParcelStore) SetStatus(number int, status string) error {
	_, err := s.db.Exec("UPDATE parcel SET status = :status WHERE number = :number",
		sql.Named("status", status),
		sql.Named("number", number))

	return err
}

func (s ParcelStore) SetAddress(number int, address string) error {
	_, err := s.db.Exec("UPDATE parcel SET address = :address WHERE number = :number AND status = :status",
		sql.Named("address", address),
		sql.Named("number", number),
		sql.Named("status", ParcelStatusRegistered))

	return err
}

func (s ParcelStore) Delete(number int) error {
	_, err := s.db.Exec("DELETE FROM parcel WHERE number = :number AND status = :status",
		sql.Named("number", number),
		sql.Named("status", ParcelStatusRegistered))

	return err
}

// package main

// import (
//     "database/sql"
//     "math/rand"
//     "testing"
//     "time"

//     "github.com/stretchr/testify/require"
// )

// var (
//     // randSource источник псевдо случайных чисел.
//     // Для повышения уникальности в качестве seed
//     // используется текущее время в unix формате (в виде числа)
//     randSource = rand.NewSource(time.Now().UnixNano())
//     // randRange использует randSource для генерации случайных чисел
//     randRange = rand.New(randSource)
// )

// // getTestParcel возвращает тестовую посылку
// func getTestParcel() Parcel {
//     return Parcel{
//         Client:    1000,
//         Status:    ParcelStatusRegistered,
//         Address:   "test",
//         CreatedAt: time.Now().UTC().Format(time.RFC3339),
//     }
// }

// // TestAddGetDelete проверяет добавление, получение и удаление посылку
// func TestAddGetDelete(t *testing.T) {
//     // prepare
//     db, err := sql.Open("sqlite", "tracker.db")
//     if err != nil {
//         require.NoError(t, err)
//     }
//     defer db.Close()
//     store := NewParcelStore(db)
//     parcel := getTestParcel()

//     // add
//     parcel.Number, err = store.Add(parcel)

//     require.NoError(t, err)
//     require.NotEmpty(t, parcel.Number)

//     // get
//     stored, err := store.Get(parcel.Number)

//     require.NoError(t, err)
//     require.Equal(t, parcel, stored)

//     // delete
//     err = store.Delete(parcel.Number)

//     stored, err = store.Get(parcel.Number)
//     require.Equal(t, sql.ErrNoRows, err)
// }

// // TestSetAddress проверяет обновление адреса
// func TestSetAddress(t *testing.T) {
//     // prepare
//     db, err := sql.Open("sqlite", "tracker.db")
//     if err != nil {
//         require.NoError(t, err)
//     }
//     defer db.Close()
//     store := NewParcelStore(db)
//     parcel := getTestParcel()

//     // add
//     parcel.Number, err = store.Add(parcel)

//     require.NoError(t, err)
//     require.NotEmpty(t, parcel.Number)

//     // set address
//     newAddress := "new test address"
//     err = store.SetAddress(parcel.Number, newAddress)

//     require.NoError(t, err)

//     // check
//     stored, err := store.Get(parcel.Number)

//     require.NoError(t, err)
//     require.Equal(t, newAddress, stored.Address)
// }

// // TestSetStatus проверяет обновление статуса
// func TestSetStatus(t *testing.T) {
//     // prepare
//     db, err := sql.Open("sqlite", "tracker.db")
//     if err != nil {
//         require.NoError(t, err)
//     }
//     defer db.Close()
//     store := NewParcelStore(db)
//     parcel := getTestParcel()

//     // add
//     parcel.Number, err = store.Add(parcel)

//     require.NoError(t, err)
//     require.NotEmpty(t, parcel.Number)

//     // set status
//     err = store.SetStatus(parcel.Number, ParcelStatusSent)

//     require.NoError(t, err)

//     // check
//     stored, err := store.Get(parcel.Number)

//     require.NoError(t, err)
//     require.Equal(t, ParcelStatusSent, stored.Status)
// }

// // TestGetByClient проверяет получение посылок по идентификатору клиента
// func TestGetByClient(t *testing.T) {
//     // prepare
//     db, err := sql.Open("sqlite", "tracker.db")
//     if err != nil {
//         require.NoError(t, err)
//     }
//     defer db.Close()
//     store := NewParcelStore(db)

//     parcels := []Parcel{
//         getTestParcel(),
//         getTestParcel(),
//         getTestParcel(),
//     }
//     parcelMap := map[int]Parcel{}

//     // задаём всем посылкам одного клиента
//     client := randRange.Intn(10_000_000)
//     parcels[0].Client = client
//     parcels[1].Client = client
//     parcels[2].Client = client

//     // add
//     for i := 0; i < len(parcels); i++ {
//         id, err := store.Add(parcels[i])

//         require.NoError(t, err)
//         require.NotEmpty(t, id)

//         parcels[i].Number = id
//         parcelMap[id] = parcels[i]
//     }

//     // get by client
//     storedParcels, err := store.GetByClient(client)

//     require.NoError(t, err)
//     require.Len(t, storedParcels, len(parcels))

//     // check
//     for _, parcel := range storedParcels {
//         expectedParcel, ok := parcelMap[parcel.Number]

//         require.True(t, ok)
//         require.Equal(t, expectedParcel, parcel)
//     }
// }
