package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// SetupRoutes настраивает все маршруты для сервера
func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Главная страница
	router.HandleFunc("/", homeHandler).Methods("GET")

	// Страница "О нас"
	router.HandleFunc("/about", aboutHandler).Methods("GET")

	// Пример маршрута с параметром {name}
	router.HandleFunc("/user/{name}", userHandler).Methods("GET")

	return router
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Добро пожаловать на главный экран!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Это страница 'О нас'")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	fmt.Fprintf(w, "Привет, %s!\n", name)
}
