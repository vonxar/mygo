package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Post struct {
	Id       int
	Content  string
	// Author   string
	AuthorName string `db: author`
}

// type Comment struct {
// 	Id      int
// 	Content string
// 	Author  string
// 	Post    *Post
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

// func process(w http.ResponseWriter, r *http.Request) {
// 	rand.Seed(time.Now().Unix())
// 	var t *template.Template
// 	if rand.Intn(10) > 5 {
// 		t, _ = template.ParseFiles("layout.html", "red.html")
// 	} else {
// 		t, _ = template.ParseFiles("layout.html")
// 	}
// 	t.ExecuteTemplate(w, "layout", "")
// }

// func form(w http.ResponseWriter, r *http.Request) {
// 	t, _ := template.ParseFiles("form.html")
// 	t.Execute(w, nil)
// }

// var PostById map[int]*Post
// var PostByAuthor map[string][]*Post

// func store(data interface{}, filename string) {
// 	buffer := new(bytes.Buffer)
// 	encoder := gob.NewEncoder(buffer)
// 	err := encoder.Encode(data)
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func load(data interface{}, filename string) {
// 	raw, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		panic(err)
// 	}
// 	buffer := bytes.NewBuffer(raw)
// 	dec := gob.NewDecoder(buffer)
// 	err = dec.Decode(data)
// 	if err != nil {
// 		panic(err)
// 	}
// }

var Db *sqlx.DB

func init() {
	var err error
	Db, err = sqlx.Open("postgres", "user=gest dbname=gest password=gest sslmode=disable")
	if err != nil {
		panic(err)
	}
}

// func (comment *Comment) Create() (err error) {
// 	if comment.Post == nil {
// 		err = errors.New("投稿が見つかりません")
// 		// err = errors.New("Post not found")
// 		return
// 	}
// 	err = Db.QueryRow("insert into comments (content, author, post_id) values ($1, $2, $3) returning id", comment.Content, comment.Author, comment.Post.Id).Scan(&comment.Id)
// 	return
// }

// func Posts(limit int) (posts []Post, err error) {
// 	rows, err := Db.Query("select id, content, author from posts limit $1", limit)
// 	if err != nil {
// 		return
// 	}
// 	for rows.Next() {
// 		post := Post{}
// 		err = rows.Scan(&post.Id, &post.Content, &post.Author)
// 		if err != nil {
// 			return
// 		}
// 		posts = append(posts, post)
// 	}
// 	rows.Close()
// 	return
// }

func GetPost(id int) (post Post, err error) {
	post = Post{}
	// post.Comments = []Comment{}
	err = Db.QueryRowx("select id, content, author from posts where id = $1", id).StructScan(&post)

	// rows, err := Db.Query("select id, content, author from comments where post_id = $1", id)
	if err != nil {
		// return
		panic(err)
	}
	// for rows.Next() {
	// 	comment := Comment{Post: &post}
	// 	err = rows.Scan(&comment.Id, &comment.Content, &comment.Author)
	// 	if err != nil {
	// 		return
	// 	}
	// 	post.Comments = append(post.Comments, comment)
	// }
	// rows.Close()
	return
}

func (post *Post) Create() (err error) {
	err = Db.QueryRow("insert into posts (content, author) values ($1, $2) returning id", post.Content, post.AuthorName).Scan(&post.Id)
	return
}

// 	statement := "insert into posts (content, author) values ($1, $2) returning id"
// 	stmt, err := Db.Prepare(statement)
// 	if err != nil {
// 		return
// 	}
// 	defer stmt.Close()
// 	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
// 	return
// }

// func (post *Post) Update() (err error) {
// 	_, err = Db.Exec("update posts set content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
// 	return
// }

// func (post *Post) Delete() (err error) {
// 	_, err = Db.Exec("delete from posts where id = $1", post.Id)
// 	return
// }

// func DeleteAll() (err error) {
// 	_, err = Db.Exec("delete from posts")
// 	return
// }

type Post struct {
	Id int
	Content string
	AuthorName string `db: "author"`
}

func main() {
	post := Post{Content: "Hello World!", AuthorName: "Sau Sheong"}
	// fmt.Println(post)
	post.Create()
	// comment := Comment{Content: "いい投稿だね", Author: "Joe", Post: &post}
	// comment.Create()
	fmt.Println(post)
	// readPost, _ := GetPost(post.Id)

	// fmt.Println(readPost)
	// fmt.Println(readPost.Comments)
	// fmt.Println(readPost.Comments[0].Post)
	// fmt.Println(post)
	// readPost, _ := GetPost(post.Id)
	// fmt.Println(readPost)

	// readPost.Content = "Bonjour Monde"
	// readPost.Author = "Pierre"
	// readPost.Update()
	// posts, _ := Posts(10)
	// fmt.Println(posts)

	// readPost.Delete()
	// csvFile, err := os.Create("posts.csv")
	// if err != nil {
	// 	panic(err)
	// }
	// defer csvFile.Close()

	// allPosts := []Post{
	// 	Post{Id: 1, Content: "Hello world", Author: "Sau Sheong"},
	// 	Post{Id: 2, Content: "Bonjpur Monde", Author: "Prierre"},
	// 	Post{Id: 3, Content: "Hello Mundo", Author: "Pedro"},
	// 	Post{Id: 4, Content: "Greetings Earthlings", Author: "Sau Sheong"},
	// }

	// writer := csv.NewWriter(csvFile)
	// for _, post := range allPosts {
	// 	line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
	// 	err := writer.Write(line)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	// writer.Flush()

	// file, err := os.Open("posts.csv")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// reader := csv.NewReader(file)
	// reader.FieldsPerRecord = -1
	// record, err := reader.ReadAll()
	// if err != nil {
	// 	panic(err)
	// }
	// var posts []Post
	// for _, item := range record {
	// 	id, _ := strconv.ParseInt(item[0], 0, 0)
	// 	post := Post{Id: int(id), Content: item[1], Author: item[2]}
	// 	posts = append(posts, post)
	// }

	// fmt.Println(posts[0].Id)
	// fmt.Println(posts[0].Content)
	// fmt.Println(posts[0].Author)
	// PostById = make(map[int]*Post)
	// PostByAuthor = make(map[string][]*Post)

	// post1 := Post{Id: 1, Content: "Hello world", Author: "Sau Sheong"}
	// post2 := Post{Id: 2, Content: "Bonjpur Monde", Author: "Prierre"}
	// post3 := Post{Id: 3, Content: "Hello Mundo", Author: "Pedro"}
	// post4 := Post{Id: 4, Content: "Greetings Earthlings", Author: "Sau Sheong"}

	// store(post1)
	// store(post2)
	// store(post3)
	// store(post4)

	// fmt.Println(PostById[1])
	// fmt.Println(PostById[2])

	// for _, post := range PostByAuthor["Sau Sheong"] {
	// 	fmt.Println(post)
	// }
	// for _, post := range PostByAuthor["Pedro"] {
	// 	fmt.Println(post)
	// }

	// mux := httprouter.New()
	// mux.GET("/hello/:name", hello)
	// server := http.Server{
	// Addr: "127.0.0.1:8080",
	// Handler: mux,
	// }
	// http.HandleFunc("/write", writerExample)
	// http.HandleFunc("/writeheandler", writeHeandlerExample)
	// http.HandleFunc("/redirect", headlerExample)
	// http.HandleFunc("/json", jsonExample)
	// http.HandleFunc("/set_cookie", setCookie)
	// http.HandleFunc("/get_cookie", getCookie)
	// http.HandleFunc("/set_message", setMessage)
	// http.HandleFunc("/show_message", showMessage)
	// http.HandleFunc("/process", process)
	// http.HandleFunc("/form", form)
	// server.ListenAndServe()

	// data := []byte("hello world\n")
	// err := ioutil.WriteFile("datal1", data, 0644)
	// if err != nil {
	// 	panic(err)
	// }
	// read1, _ := ioutil.ReadFile("data1")
	// fmt.Print(string(read1))

	// file1, _ := os.Create("data2")
	// defer file1.Close()
	// bytes, _ := file1.Write(data)
	// fmt.Printf("Wrote %d bytes to file\n", bytes)

	// file2, _ := os.Open("data2")
	// defer file2.Close()

	// read2 := make([]byte, len(data))
	// bytes, _ = file2.Read(read2)
	// fmt.Printf("REad %d bytes from file\n", bytes)
	// fmt.Println(string(read2))
}
