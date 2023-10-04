package main

import (
	"fmt"      // this package is for working with std input and output
	"log"      // this package is log errors
	"net/http" // this is the most important package to deal with http rest API and routes
)

// / Every go program execution starts from this `main` function.
func main() {
	/// we are telling go-lang to check for the static files present in the `static` directory.
	/// by default `index.html` is the file which is loaded on the root route `/`
	fileServer := http.FileServer(http.Dir("./static"))
	/// handle the root route and pass the base static page
	http.Handle("/", fileServer)

	/// handler other routes here.
	///
	/// each handler function should provide 2 input params: ResponseWrite and Request.
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/admin", adminHandler)

	/// start the server on the port.
	port := ":8080"
	fmt.Printf("Starting new Server at Port: %s", port)

	/// this is a important statement to start and listen to the server.
	/// if error starting the server, log a fatal error.
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("Error Parsing Form :", err)
		return
	}

	// if r.Method != http.MethodPost {
	// 	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	// 	return
	// }

	fmt.Fprintf(w, "Post Form Request Successfull")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "\n\nName: %s \n", name)
	fmt.Fprintf(w, "\nAddress: %s \n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	/// check if this is a valid route or not.
	/// this step is not mandatory but makes sure that this function is not accidentally called.
	/// if wrong path, throw 404 error.
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	/// check if the method is correct or not. We are supporting only get method of this route.
	/// if method is anything apart from GET, throw 405 error, method not allowed
	if r.Method != http.MethodGet {
		http.Error(w, "method is not supported", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Hello! This is cool. We passed all the validations for this API.")
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "This is a private route for the admin.", http.StatusNotFound)
}
