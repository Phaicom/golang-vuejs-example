package endpoints

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/phaicom/golang-vuejs-example/backend/models"
)

type postData struct {
	Size int64 `json:"size"`
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetPost(w, r)
	case "POST":
		CreatePost(w, r)
	default:
		GetPost(w, r)
	}
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	// mock data
	posts, err := models.GetPosts()
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(posts)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	data := new(postData)
	json.NewDecoder(r.Body).Decode(data)
	posts, err := models.CreatePosts(int(data.Size))
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(posts)
}
