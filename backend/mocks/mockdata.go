package mocks

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/phaicom/golang-vuejs-example/backend/models"
)

var usernames = []string{
	"John Doe",
	"Mark Jacop",
	"Alice Smith",
}

var messages = []string{
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad.",
	"Minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit.",
	"In voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia.",
	"Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa.",
}

func CreateJson() {
	posts := make([]models.Post, 0)
	for i := 0; i < 5; i++ {
		user := models.User{
			CreatedBy:   usernames[rand.Intn(len(usernames)-1)],
			TimeCreated: time.Now(),
		}
		rand := rand.Intn(len(messages) - 1)
		post := models.Post{
			PostID:      int64(i),
			Caption:     messages[rand][:20],
			MessageBody: messages[rand],
			ImageURL:    "https://via.placeholder.com/150x150",
			User:        user,
		}

		posts = append(posts, post)
	}

	data, err := json.Marshal(posts)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("./posts.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Profiles data written to ", file.Name())
}

func OpenJson() []models.Post {
	posts := make([]models.Post, 0)

	file, err := os.Open("./posts.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&posts)
	if err != nil {
		log.Fatal(err)
	}
	return posts
}

// func FindByUsername(username string) models.Profile {
// 	profile := models.Profile{}
// 	profiles := OpenJson()
// 	for _, p := range profiles {
// 		if p.Username == username {
// 			profile = p
// 			break
// 		}
// 	}
// 	return profile
// }
