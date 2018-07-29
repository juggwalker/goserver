package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type User struct {
	Name string `json:"name1"`
	Age  int    `json:"age1"`
}

type EmailPostFormat struct {
	Form string `json:"from"`
	To   string `json:"to"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("Form: ", r.Form)
	fmt.Println("Path: ", r.URL.Path)
	fmt.Println(r.Form["a"])
	fmt.Println(r.Form["b"])
	fmt.Println(r.Method)
	for k, v := range r.Form {
		fmt.Println(k, "=>", v, strings.Join(v, "-"))
	}

	io.WriteString(w, "Hello, GoServer!\n")
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.URL.Path)
}

func sendEmailHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "It works !")

	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Read failed:", err)
		}
		defer r.Body.Close()

		body_str := string(body)
		fmt.Println("body_str:", body_str)

		var user User

		if err := json.Unmarshal(body, &user); err == nil {
			fmt.Println(user)
			user.Age += 100
			fmt.Println(user)
			ret, _ := json.Marshal(user)
			fmt.Fprint(w, string(ret))
		} else {
			log.Println("json format error:", err)
		}

	} else {

		log.Println("ONly support Post")
		fmt.Fprintf(w, "Only support post")
	}

	//go SendMail()
	io.WriteString(w, "发送邮件\n")
}

func HttpServ() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/", echoHandler)
	mux.HandleFunc("/sendmail", sendEmailHandler)

	if err := http.ListenAndServe(":12345", mux); err != nil {
		log.Fatal("http boot fail")
	}

}

func FileServ() {
	http.ListenAndServe(":12346", http.FileServer(http.Dir("./storage")))
}
