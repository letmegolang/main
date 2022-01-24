package main

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"net/http"
)

var app_sec_key = "some_secret_key" //this should be in an env, not in code in production!

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/login/{username}/{password}", LoginHandler)
	mux.HandleFunc("/protected/{token}", ProtectedHandler)

	http.Handle("/", mux)

	srv := &http.Server{
		Addr:    ":80",
		Handler: mux,
	}
	srv.ListenAndServe()
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	password := vars["password"]

	w.Header().Set("Content-Type", "text/html")

	// check username and password to get JWT token
	if username != "username" || password != "password" {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "No login or password\n")
		return
	}
	w.WriteHeader(http.StatusOK)

	token, ok := CreateToken(100)
	if ok == nil {
		println(token)

		fmt.Fprintf(w, "Token: %v\n", token)
		fmt.Fprintf(w, "<br><a href=/protected/%v>to protected page</a>\n", token)
	}

}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	token := vars["token"]
	w.WriteHeader(http.StatusOK)
	data_from_token := ParseToken(token, "some_secret_key")
	fmt.Fprintf(w, "Data from token: %v\n", data_from_token)
}

func CreateToken(userId uint64) (string, error) {
	var err error

	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["some_data"] = "Some data"
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(app_sec_key))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(tokenString string, secretkey string) string {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretkey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Printf("Data from token:\n %v \n %v \n %s \n", claims["authorized"], claims["user_id"], claims["some_data"])
		some_result := fmt.Sprintf("%v\n%v\n%s\n", claims["authorized"], claims["user_id"], claims["some_data"])
		return some_result

	} else {
		fmt.Println(err)
		return "no data to display"
	}
}
