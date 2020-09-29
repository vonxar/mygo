package main

import(
	"testing"
)

func TestHAndleGet(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest)

	writer := httptest.NewRwquest()
	request, _ := http.NewRequest("GET", "/post/1", nil)
	mux.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Errorf("Response cose is %v", writer.Code)
	}
	var post post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 1{
		t.Error("Cannot retrieve JSON post")
	}
}