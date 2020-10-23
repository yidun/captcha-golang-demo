package main

import (
	"fmt"
	"github.com/galaxy-book/captcha-golang-demo/sdk"
	"log"
	"net/http"
)

const(
	captchaId = "YOUR_CAPTCHA_ID"
	secretId = "YOUR_SECRET_ID"
	secretKey = "YOUR_SECRET_KEY"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("web")))
	http.HandleFunc("/verify", verify)

	fmt.Println("server is success listen on 8080.")
	fmt.Println("web entry: http://127.0.0.1:8080/index.html")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func verify(w http.ResponseWriter, r *http.Request) {
	validate := r.FormValue("validate")
	verifier, err := sdk.New(captchaId, secretId, secretKey)
	user := "nico"
	if err != nil{
		w.Write([]byte(err.Error()))
		return
	}
	verifyResult, err := verifier.Verify(validate, user)
	if err != nil{
		w.Write([]byte(err.Error()))
		return
	}
	if verifyResult.Result{
		w.Write([]byte("验证成功！"))
	}else{
		w.Write([]byte("验证失败！"))
	}
}