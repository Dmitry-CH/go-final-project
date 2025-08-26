package api

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Dmitry-CH/go-final-project/pkg/config"
	"github.com/golang-jwt/jwt"
)

var ErrBadSignToken = errors.New("ошибка не удалось подписать токен")

var secretKey = []byte("my_secret_key")

type user struct {
	Password string `json:"password"`
}

func signinHandler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	var user user

	conf := config.New()

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		writeJson(w, ErrResp{err.Error()}, http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(buf.Bytes(), &user)
	if err != nil {
		writeJson(w, ErrResp{err.Error()}, http.StatusBadRequest)
		return
	}

	ePass := conf.Password
	if user.Password != ePass {
		writeJson(w, ErrResp{"invalid password"}, http.StatusBadRequest)
		return
	}

	claims := jwt.MapClaims{
		"pass": generateSum(ePass),
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := jwtToken.SignedString(secretKey)
	if err != nil {
		writeJson(w, ErrResp{err.Error()}, http.StatusBadRequest)
		return
	}

	writeJson(w, TokenResp{signedToken})
}

func generateSum(s string) string {
	hash := md5.Sum([]byte(s))

	return hex.EncodeToString(hash[:])
}
