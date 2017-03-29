package controllers

import (
	"net/http"
	"strings"
	"fmt"
	"github.com/lifei6671/go-git-webhook-client/conf"
	"github.com/lifei6671/go-git-webhook-client/models"
	"github.com/lifei6671/go-git-webhook-client/cache"
	"crypto/sha256"
	"encoding/hex"
	"crypto/md5"
	"time"
)


func Token(w http.ResponseWriter, r *http.Request)  {

	result := &models.JsonResult{
		ErrorCode:500,
		Message:"Method not allow.",
	}

	w.Header().Add("Content-Type","application/json; charset=UTF-8")
	if !strings.EqualFold("POST",r.Method) {

		s,_ := result.JsonString()

		fmt.Fprint(w,s)
		return
	}

	if err := r.ParseForm();err != nil {
		result.Message = err.Error()
		s,_ := result.JsonString()

		fmt.Fprint(w,s)
		return
	}

	account := r.FormValue("account")
	password := r.FormValue("password")
	t := r.FormValue("time")

	if !strings.EqualFold(conf.GetString("account",""),account) {
		result.Message = "Permission denied."
		s,_ := result.JsonString()
		fmt.Fprint(w,s)
		return
	}

	rawPassword := conf.GetString("password","")

	hash := sha256.New()
	hash.Write([]byte(account + rawPassword + t))
	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)

	if !strings.EqualFold(mdStr,password) {
		result.Message = "Permission denied:50001."
		s,_ := result.JsonString()
		fmt.Fprint(w,s)
		return
	}

	data := &models.JsonResult{
		ErrorCode:0,
		Message:"ok",
	}

	hash = md5.New()
	hash.Write([]byte(account + password + time.Now().String()))
	md = hash.Sum(nil)
	token := hex.EncodeToString(md)

	data.Data = token

	s ,_ := data.JsonString()

	cache.TokenCache.Add(token,models.Member{
		Account:account,
		Password:rawPassword,
		Token:token,
	})

	fmt.Fprint(w,s)

}


