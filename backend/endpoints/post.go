package endpoints

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/phaicom/golang-vuejs-example/backend/models"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	// mock data
	posts := make([]models.Post, 0)
	for i := 1; i <= 5; i++ {
		user := models.User{
			CreatedBy:   "John Doe",
			TimeCreated: time.Now(),
		}

		post := models.Post{
			PostID:      int64(i),
			Caption:     "Caption",
			ImageURL:    "https://via.placeholder.com/150x150",
			MessageBody: "Message Body",
			User:        user,
		}

		posts = append(posts, post)
	}
	json.NewEncoder(w).Encode(posts)
}
