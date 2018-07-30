package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type EmailPostFormat struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type ConfPostFormat struct {
	Name     string `json:"name"`
	Contents string `json:"contents"`
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
			Log.Println("Read failed:", err)
		}
		defer r.Body.Close()

		body_str := string(body)
		Log.Println("body_str:", body_str)

		var post EmailPostFormat

		if err := json.Unmarshal(body, &post); err == nil {
			//Log.Println(post)
			go (&Email{To: post.To, Subject: post.Subject, Body: post.Body}).Send()
			ret, _ := json.Marshal(ret)
			io.WriteString(w, string(ret))
		} else {
			Log.Warn("json format error:", err)
		}

	} else {
		ret["msg"] = "Must post method"
		ret["code"] = 500
		ret, _ := json.Marshal(ret)
		io.WriteString(w, string(ret))
		Log.Warn("ONly support Post")
	}

}

func confHandler(w http.ResponseWriter, r *http.Request) {
	var post ConfPostFormat
	ret := make(map[string]interface{})
	ret["msg"] = "ok"
	ret["errmsg"] = 0
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			Log.Println("Read failed:", err)
		}
		defer r.Body.Close()

		if err := json.Unmarshal(body, &post); err == nil {
			go (&FileInfo{FileName: post.Name, FilePath: "."}).WriteWithIo(post.Contents)
			ret, _ := json.Marshal(ret)
			io.WriteString(w, string(ret))
		} else {
			Log.Warn("json format error:", err)
		}

	} else {
		ret["msg"] = "Must post method"
		ret["code"] = 500
		ret, _ := json.Marshal(ret)
		io.WriteString(w, string(ret))
		Log.Warn("ONly support Post")
	}

}

func HttpServ() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", echoHandler)
	mux.HandleFunc("/conf", confHandler)
	mux.HandleFunc("/sendmail", sendEmailHandler)

	port := fmt.Sprintf(":%d", Config["server_http_port"])
	if err := http.ListenAndServe(port, mux); err != nil {
		Log.Fatal("http boot fail")
	}

}

func FileServ() {
	http.ListenAndServe(":12346", http.FileServer(http.Dir("./storage")))
}
