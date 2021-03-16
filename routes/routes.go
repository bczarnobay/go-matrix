package routes

import (
	"net/http"
)

func Setup() {
	http.HandleFunc("/sum", sumController)
	http.HandleFunc("/multiply", multiplyController)
	http.HandleFunc("/flatten", flattenController)
	http.HandleFunc("/invert", invertController)
	http.HandleFunc("/echo", echoController)
}
