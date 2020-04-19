package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("value:", strings.Join(v, ""))
	}

	fmt.Fprintf(w, "Hello Jaewon!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)

	if r.Method == "GET" {
		currentTime := time.Now().Unix()
		hash := md5.New()

		io.WriteString(hash, strconv.FormatInt(currentTime, 10))
		token := fmt.Sprintf("%x", hash.Sum(nil))

		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseForm()

		token := r.Form.Get("token")
		if token != "" {
			// check token validity
		} else {
			// give error if no token
		}

		fmt.Println("username length:", len(r.Form["username"][0]))
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username")))
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username")))
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)

	if r.Method == "GET" {
		currentTime := time.Now().Unix()
		hash := md5.New()

		io.WriteString(hash, strconv.FormatInt(currentTime, 10))
		token := fmt.Sprintf("%x", hash.Sum(nil))

		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)

	} else {
		r.ParseMultipartForm(32 << 20)

		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}

		defer file.Close()

		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

// Simulate upload file in this client

// func postFile(filename string, targetUrl string) error {
// 	bodyBuf := &bytes.Buffer{}
// 	bodyWriter := multipart.NewWriter(bodyBuf)

// 	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
// 	if err != nil {
// 		fmt.Println("error writing to buffer")
// 		return err
// 	}

// 	fh, err := os.Open(filename)
// 	if err != nil {
// 		fmt.Println("error opening file")
// 		return err
// 	}

// 	defer fh.Close()

// 	_, err = io.Copy(fileWriter, fh)
// 	if err != nil {
// 		return err
// 	}

// 	contentType := bodyWriter.FormDataContentType()
// 	bodyWriter.Close()

// 	resp, err := http.Post(targetUrl, contentType, bodyBuf)
// 	if err != nil {
// 		return err
// 	}

// 	defer resp.Body.Close()

// 	respBody, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println(resp.Status)
// 	fmt.Println(string(respBody))

// 	return nil
// }

func main() {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
