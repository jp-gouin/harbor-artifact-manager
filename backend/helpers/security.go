package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"backend/args"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
)

var httpClient = &http.Client{}

// TokenInfo struct
type Tokeninfo struct {
	Iss string `json:"iss"`
	HD  string `json:"hd"`
	// userId
	Sub string `json:"sub"`
	Azp string `json:"azp"`
	// clientId
	Aud string `json:"aud"`
	Iat int64  `json:"iat"`
	// expired time
	Exp int64 `json:"exp"`

	Email         string      `json:"email"`
	EmailVerified bool        `json:"email_verified"`
	AtHash        string      `json:"at_hash"`
	Name          string      `json:"name"`
	GivenName     string      `json:"given_name"`
	FamilyName    string      `json:"family_name"`
	Picture       string      `json:"picture"`
	Local         string      `json:"locale"`
	Data          interface{} `json:"-"`
}

func GetUsername(r *http.Request) (string, error) {
	var token string
	signIn := r.Header.Get("signIn")
	signInWS := r.URL.Query()["signIn"]
	if len(signInWS) != 0 && signInWS[0] == "google" {
		tokens, _ := r.URL.Query()["token"]
		token = tokens[0]
		parsedToken, err := VerifyIdToken(token)
		if err != nil {
			return "", err
		}
		user := parsedToken.Email
		return user, err
	} else if signIn == "Google" || (len(signInWS) != 0 && signInWS[0] == "google") {
		authHeader := r.Header.Get("Authorization")
		splitToken := strings.Split(authHeader, "Bearer ")
		reqToken := splitToken[1]
		parsedToken, err := VerifyIdToken(reqToken)
		if err != nil {
			return "", err
		}
		user := parsedToken.Email
		return user, err
	} else {
		tokens, _ := r.URL.Query()["token"]
		jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
			ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
				return []byte(args.Args.JwtKey), nil
			},
			SigningMethod: jwt.SigningMethodHS256,
		})
		if len(tokens) == 0 {
			token, _ = jwtMiddleware.Options.Extractor(r)
		} else {
			token = tokens[0]
		}
		// Now parse the token
		parsedToken, err := jwt.Parse(token, jwtMiddleware.Options.ValidationKeyGetter)
		user := parsedToken.Claims.(jwt.MapClaims)["user"].(string)
		return user, err
	}

}
func IsAdmin(username string) bool {
	admins := strings.Split(args.Args.Admins, ",")
	for _, admin := range admins {
		fmt.Println("isAdmin : " + admin)
		if admin == username {
			return true
		}
	}
	return false
}
func VerifyIdToken(idToken string) (*Tokeninfo, error) {
	req, err := http.NewRequest("GET", "https://oauth2.googleapis.com/tokeninfo?id_token="+idToken, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var tokeninfo Tokeninfo
	json.Unmarshal([]byte(body), &tokeninfo)

	return &tokeninfo, nil
}
