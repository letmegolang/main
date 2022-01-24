package main

import (
	"fmt"
	"net/http"
)

type application struct {
	auth struct {
		username string
		password string
	}
}

func main() {
	app := new(application)

	app.auth.username = "username"
	app.auth.password = "password"

	mux := http.NewServeMux()
	mux.HandleFunc("/public", app.unprotectedHandler)

	// set protection to /protected page
	mux.HandleFunc("/protected", app.basicAuth(app.protectedHandler))

	srv := &http.Server{
		Addr:    ":80",
		Handler: mux,
	}
	srv.ListenAndServe()
}

func (app *application) protectedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the protected page")
}

func (app *application) unprotectedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the public page")
}

func (app *application) basicAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()

		if !ok {
			fmt.Println("Error parsing basic auth")
			w.WriteHeader(401)
			return
		}

		if app.auth.username != username {
			fmt.Printf("Wrong username \n")
			w.WriteHeader(401)
			return
		}

		if app.auth.password != password {
			fmt.Printf("Wrong password \n")
			w.WriteHeader(401)
			return
		}

		fmt.Printf("Username: %s\n", username)
		fmt.Printf("Password: %s\n", password)
		next.ServeHTTP(w, r)
		return
	})
}
