package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"github.com/student-management/models"

	"github.com/joho/godotenv"

	"github.com/dgrijalva/jwt-go"
)

func Authentication(h http.Handler) http.Handler {
	// load project env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("error loading env file %v", err)
	}

	authJ := http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		// get jwt token from the header prefixed with Bearer
		authHeader := strings.Split(req.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			wr.Header().Set("content-type", "application/json")
			wr.WriteHeader(http.StatusUnauthorized)
			respErr, _ := json.Marshal(models.HttpErrs{ErrMsg: "Unauthorized", ErrCode: http.StatusUnauthorized})
			_, _ = wr.Write(respErr)
		} else {
			// get the jwt token if it exists
			jwtToken := authHeader[1]

			// check if the token is signed with correct algo and signed with the given secret key
			token, _ := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(os.Getenv("SECRETKEY")), nil
			})

			// check if the token is valid
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				v := true
				if !(claims["studentname"] != "" && strings.Split(claims["studentname"].(string), "@")[1] == "gmail.com") {
					v = false
				}
				// can get from the database
				if !(claims["password"] != "" && claims["password"].(string) == "1234") {
					v = false
				}

				if v {
					h.ServeHTTP(wr, req)
				} else {
					wr.Header().Set("content-type", "application/json")
					wr.WriteHeader(http.StatusUnauthorized)
					respErr, _ := json.Marshal(models.HttpErrs{ErrMsg: "Unauthorized", ErrCode: http.StatusUnauthorized})
					_, _ = wr.Write(respErr)
				}
			} else {
				wr.Header().Set("content-type", "application/json")
				wr.WriteHeader(http.StatusUnauthorized)
				respErr, _ := json.Marshal(models.HttpErrs{ErrMsg: "Unauthorized", ErrCode: http.StatusUnauthorized})
				_, _ = wr.Write(respErr)
			}
		}
	})

	return authJ
}
