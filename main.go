package main
import {
	"log"
	"net/http"
}

func homepage(w http.ResponseWriter, r *http.Request) {
	http.