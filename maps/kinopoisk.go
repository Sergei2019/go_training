package main

import "fmt"

var mapa = make(map[string]float64)

func main() {
	mapa["Дюна"] = 9.4
	mapa["Интерстеллар"] = 8.0
	mapa["Довод"] = 7.5
	addFilm("Дюна", 7)
	addFilm("Робокоп", 7)
	fmt.Println(topFilm())
	fmt.Println(avgRat())

}

func addFilm(s string, rat float64) {
	if _, ok := mapa[s]; ok {
		mapa[s] = rat
	} else {
		fmt.Println("Фильм уже добавлен")
	}
	return

}

func topFilm() map[string]float64 {
	res_mapa := make(map[string]float64)
	for key, val := range mapa {
		if val > 7.0 {
			res_mapa[key] = val
		}
	}
	return res_mapa
}

func avgRat() float64 {
	var sum = 0.0
	for _, val := range mapa {
		sum = sum + val
	}
	aver := sum / float64(len(mapa))
	return aver
}
