package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"backend/args"
	"backend/helpers"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
)

/*
User struct
*/
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	AuthCode string `json:"authcode"`
}

var (
	conf = &oauth2.Config{
		ClientID:     args.Args.GoogleClientID,
		ClientSecret: args.Args.GoogleSecret,
		RedirectURL:  "postmessage",
		Endpoint:     google.Endpoint,
	}
)

// TokenHandler is our handler to take a username and password and,
// if it's valid, return a token used for future requests.
func TokenHandler(w http.ResponseWriter, r *http.Request) {
	nm.Log("TokenHandler")
	w.Header().Add("Content-Type", "application/json")

	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(reqBody, &user)

	if !authentUser(user.Username, user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, `{"error":"invalid_credentials"}`)
		return
	}

	// We are happy with the credentials, so build a token. We've given it
	// an expiry of 1 hour.
	var token *jwt.Token
	if helpers.IsAdmin(user.Username) {
		token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user": user.Username,
			"exp":  time.Now().Add(time.Hour * time.Duration(1)).Unix(),
			"iat":  time.Now().Unix(),
			"id":   "1001-0000-1001-aigh",
		})
	} else {
		token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user": user.Username,
			"exp":  time.Now().Add(time.Hour * time.Duration(1)).Unix(),
			"iat":  time.Now().Unix(),
		})
	}

	tokenString, err := token.SignedString([]byte(args.Args.JwtKey))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error":"token_generation_failed"}`)
		return
	}
	io.WriteString(w, `{"token":"`+tokenString+`"}`)
	return
}

// AuthMiddleware is our middleware to check our token is valid. Returning
// a 401 status to the client if it is not valid.
func AuthMiddleware(next http.Handler) http.Handler {
	nm.Log("AuthMiddleware")
	if len(args.Args.JwtKey) == 0 {
		log.Fatal("HTTP server unable to start, expected an APP_KEY for JWT auth")
	}
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(args.Args.JwtKey), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	return authHandler(jwtMiddleware, next)

	//return jwtMiddleware.Handler(next)

}

func myHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
}

func authHandler(jwtMiddleware *jwtmiddleware.JWTMiddleware, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		signIn := r.Header.Get("signIn")
		if signIn == "Google" {
			authHeader := r.Header.Get("Authorization")
			splitToken := strings.Split(authHeader, "Bearer ")
			reqToken := splitToken[1]
			result, err := helpers.VerifyIdToken(reqToken)
			if err != nil {
				nm.Log("error")
				return
			}
			if result.Email != "" {
				h.ServeHTTP(w, r)
			}
		} else {
			err := jwtMiddleware.CheckJWT(w, r)
			if err != nil {
				return
			}
			h.ServeHTTP(w, r)
		}
	})
}
