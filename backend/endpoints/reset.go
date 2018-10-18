package endpoints

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/phaicom/golang-vuejs-example/backend/models"
)

func ResetHandler(w http.ResponseWriter, r *http.Request) {
	_, err := models.ResetJSON()
	if err != nil {
		log.Fatal(err)
	}

	posts, err := models.GetPosts()
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(posts)
}
