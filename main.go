package main
import {
	"log"
	"net/http"
}

func homepage(w http.ResponseWriter, r *http.Request) {
	//Render the home html page from static folder
	http.ServeFile(w, r, "static/home.html")
}

func coursePage(w http.ResponseWriter, r *http.Request) {
	//Render the course html page
	http.ServeFile(w, r, "static/courses.html")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	//Render the about html page
	http.ServeFile(w, r, "static/about.html")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	//Render the contact html page
	http.ServeFile(w, r, "static/contact.html")
}

func main() {
	//Handle the routes for the web application
	http.HandleFunc("/", homepage)
	http.HandleFunc("/courses", coursePage)
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/contact", contactPage)

	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
