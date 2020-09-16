package main

import (
	"math/rand"
	"net/http"
	"text/template"
	"time"
)

// 送信構造体
// type Post struct {
// 	User    string
// 	Threads []string
// }

// 書き込み
// func writerExample(w http.ResponseWriter, r *http.Request) {
// 	str := `<html>
// 	<head><title>Go Web Programming</title></head>
// 	<body>hello world</body>
// 	</html>`
// 	w.Write([]byte(str))
// }

// func writeHeandlerExample(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(501)
// 	fmt.Fprintln(w, "そのようなサービスはありません。他を当てってください")
// }

// func headlerExample(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Location", "http://google.com")
// 	w.WriteHeader(302)
// }

// json形式で渡す
// func jsonExample(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	post := &Post{
// 		User:    "Sau Sheong",
// 		Threads: []string{"1番目", "2番目", "3番目"},
// 	}
// 	json, _ := json.Marshal(post)
// 	w.Write(json)
// }

//helloの表示
// func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	fmt.Fprintf(w, "hello, %s\n", p.ByName("name"))
// }

// func log(h http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
// 		fmt.Println("Hendler function called -" + name)
// 		h(w, r)
// 	}
// }

// func world(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "World!")
// }

// func headers(w http.ResponseWriter, r *http.Request) {
// 	h := r.Header
// 	fmt.Fprintln(w, h)
// }

// func body(w http.ResponseWriter, r *http.Request) {
// 	len := r.ContentLength
// 	body := make([]byte, len)
// 	r.Body.Read(body)
// 	fmt.Fprintln(w, string(body))
// }

// func setMessage(w http.ResponseWriter, r *http.Request) {
// 	msg := []byte("hello world")
// 	c := http.Cookie{
// 		Name:  "flash",
// 		Value: base64.URLEncoding.EncodeToString(msg),
// 	}
// 	http.SetCookie(w, &c)
// }

// func showMessage(w http.ResponseWriter, r *http.Request) {
// 	c, err := r.Cookie("flash")
// 	if err != nil {
// 		if err == http.ErrNoCookie {
// 			fmt.Fprintln(w, "メッセージがありません")
// 		}
// 	} else {
// 		rc := http.Cookie{
// 			Name:    "flash",
// 			MaxAge:  -1,
// 			Expires: time.Unix(1, 0),
// 		}
// 		http.SetCookie(w, &rc)
// 		val, _ := base64.URLEncoding.DecodeString(c.Value)
// 		fmt.Fprintln(w, string(val))
// 	}
// }

// func setCookie(w http.ResponseWriter, r *http.Request) {
// 	c1 := http.Cookie{
// 		Name:     "first_cookie",
// 		Value:    "Go Web Progrmming",
// 		HttpOnly: true,
// 	}
// 	c2 := http.Cookie{
// 		Name:     "second_cookie",
// 		Value:    "Manning Publication",
// 		HttpOnly: true,
// 	}
// 	http.SetCookie(w, &c1)
// 	http.SetCookie(w, &c2)
// }

// func getCookie(w http.ResponseWriter, r *http.Request) {
// 	c1, err := r.Cookie("first_cookie")
// 	if err != nil {
// 		fmt.Fprintln(w, "Cannnot get the first cookie")
// 	}
// 	cs := r.Cookies()
// 	fmt.Fprintln(w, c1)
// 	fmt.Fprintln(w, cs)
// }

// func formatDate(t time.Time) string {
// 	layout := "2006-01-02"
// 	return t.Format(layout)
// }

func process(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	var t *template.Template
	if rand.Intn(10) > 5 {
		t, _ = template.ParseFiles("layout.html", "red.html")
	} else {
		t, _ = template.ParseFiles("layout.html")
	}
	t.ExecuteTemplate(w, "layout", "")
}

func form(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("form.html")
	t.Execute(w, nil)
}

func main() {
	// mux := httprouter.New()
	// mux.GET("/hello/:name", hello)
	server := http.Server{
		Addr: "127.0.0.1:8080",
		// Handler: mux,
	}
	// http.HandleFunc("/write", writerExample)
	// http.HandleFunc("/writeheandler", writeHeandlerExample)
	// http.HandleFunc("/redirect", headlerExample)
	// http.HandleFunc("/json", jsonExample)
	// http.HandleFunc("/set_cookie", setCookie)
	// http.HandleFunc("/get_cookie", getCookie)
	// http.HandleFunc("/set_message", setMessage)
	// http.HandleFunc("/show_message", showMessage)
	http.HandleFunc("/process", process)
	http.HandleFunc("/form", form)
	server.ListenAndServe()
}
