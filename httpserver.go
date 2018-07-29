package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type EmailPostFormat struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.URL.Path)
}

func sendEmailHandler(w http.ResponseWriter, r *http.Request) {
	ret := make(map[string]interface{})
	ret["msg"] = "send ok"
	ret["errmsg"] = 0
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Read failed:", err)
		}
		defer r.Body.Close()

		body_str := string(body)
		log.Println("body_str:", body_str)

		var post EmailPostFormat

		if err := json.Unmarshal(body, &post); err == nil {
			//log.Println(post)
			go (&Email{To: post.To, Subject: post.Subject, Body: post.Body}).Send()
			ret, _ := json.Marshal(ret)
			io.WriteString(w, string(ret))
		} else {
			log.Warn("json format error:", err)
		}

	} else {
		ret["msg"] = "Must post method"
		ret["code"] = 500
		ret, _ := json.Marshal(ret)
		io.WriteString(w, string(ret))
		log.Warn("ONly support Post")
	}

}

func HttpServ() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", echoHandler)
	mux.HandleFunc("/sendmail", sendEmailHandler)

	if err := http.ListenAndServe(":12345", mux); err != nil {
		log.Fatal("http boot fail")
	}

}

func FileServ() {
	http.ListenAndServe(":12346", http.FileServer(http.Dir("./storage")))
}
