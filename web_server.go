package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Обробник запитів
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	// Запуск веб-серверу на порту 8080
	fmt.Println("Сервер запущено на порту 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}
