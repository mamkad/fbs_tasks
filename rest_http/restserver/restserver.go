/*
 * REST HTTP сервис
 */

package RESTapi

import (
	"./fibonacci"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//GET по адресу http://localhost:8080/input
//Загрузка начальной страницы input_form.html
//Форма для ввода
func InputHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("input_form.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

//POST по адресу http://localhost:8080/calc
//Функция получает диапазон номеров из формы ввода и вычисляет числа фибоначчи
//Затем редеректить на страницу с ответом
func CalcHandler(writer http.ResponseWriter, request *http.Request) {
	Rang := request.FormValue("range")
	i := strings.Index(Rang, ",")
	address := "/result"
	
	if i == -1 {
		address = "/input"
	} else {
		start, err := strconv.Atoi(strings.TrimSpace(Rang[:i]))
		if err != nil {
			address = "/input"
		} else {
			end, err := strconv.Atoi(strings.TrimSpace(Rang[i+1:]))
			if err != nil {
				address = "/input"
			} else {
				fibData, err := fibonacci.Fibonacci(int64(start), int64(end))
				if err != nil {
					address = "/input"
				} else {
					options := os.O_WRONLY | os.O_CREATE
					file, err := os.OpenFile("data.json", options, os.FileMode(0600))
					check(err)
					json.NewEncoder(file).Encode(fibData)
					err = file.Close()
					check(err)
				}
			}
		}
	}
	http.Redirect(writer, request, address, http.StatusFound)
}

//GET по адресу http://localhost:8080/result
//Выводит таблицу с ответом. Страницу result.html
func ResultHandler(writer http.ResponseWriter, request *http.Request) {
	file, err := os.Open("data.json")
	check(err)
	defer file.Close()

	var fibdata map[int64]int64
	json.NewDecoder(file).Decode(&fibdata)

	html, err := template.ParseFiles("result.html")
	check(err)
	err = html.Execute(writer, fibdata)
	check(err)
}

func Start() {
	http.HandleFunc("/input", InputHandler)
	http.HandleFunc("/calc", CalcHandler)
	http.HandleFunc("/result", ResultHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
