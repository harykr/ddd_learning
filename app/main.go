package main

import (
	"awesomeProject/app/handler"
	"log"
	"net/http"
)

func main() {
	addr := ":8080"
	shoppingHandler := handler.ShoppingHandler{}
	http.HandleFunc("/add-to-card", shoppingHandler.AddToCart)
	log.Println("listen on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
