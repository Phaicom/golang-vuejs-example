package models

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
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

type User struct {
	CreatedBy    string    `json:"createdBy"`
	ModifiedBy   string    `json:"modifiedBy"`
	TimeCreated  time.Time `json:"timeCreated"`
	TimeModified time.Time `json:"timeModified"`
}

type Post struct {
	PostID      int64  `json:"postID"`
	Caption     string `json:"caption"`
	MessageBody string `json:"messageBody"`
	ImageURL    string `json:"imageURL"`
	User               // Embedded type
}

func CreatePost(username string, caption string, messageBody string, imageURL string) (*Post, error) {
	user := User{CreatedBy: username, TimeCreated: time.Now()}
	post := Post{User: user, Caption: caption, MessageBody: messageBody, ImageURL: imageURL}

	// add post to db
	// if err return (zero value, err)
	// else if add completed
	// return (completed post, nil)
	return &post, nil
}

func CreatePosts(size int) (*[]Post, error) {
	posts, err := ReadJson()
	if err != nil {
		return nil, err
	}
	start := len(posts)

	for i := start; i <= start+size; i++ {
		user := User{
			CreatedBy:   usernames[rand.Intn(len(usernames)-1)],
			TimeCreated: time.Now(),
		}
		rand := rand.Intn(len(messages) - 1)
		post := Post{
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
		return nil, err
	}

	file, err := os.Create("./posts.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return nil, err
	}
	fmt.Println("Profiles data written to ", file.Name())
	return &posts, nil
}

func GetPosts() (*[]Post, error) {
	// get collection of post from db
	posts, err := ReadJson()
	if err != nil {
		return nil, err
	}
	// if err return (zero value, err)
	// else if add completed
	// return (completed []post, nil)
	return &posts, nil
}

func UpdatePost(username string, caption string, messageBody string, imageURL string) (*Post, error) {
	// get post by ID
	// post := getPostByID()
	user := User{ModifiedBy: username, TimeModified: time.Now()}
	post := Post{User: user, Caption: caption, MessageBody: messageBody, ImageURL: imageURL}

	// update post to db
	// if err return (zero value, err)
	// else if add completed
	// return (completed post, nil)
	return &post, nil
}

func DeletePost(username string, caption string, messageBody string, imageURL string) (*Post, error) {
	// get post by ID
	// post := getPostByID()
	post := new(Post)

	// delete post from db
	// if err return (zero value, err)
	// else if add completed
	// return (completed post, nil)
	return post, nil
}

func ReadJson() ([]Post, error) {
	posts := make([]Post, 0)

	file, err := os.Open("./posts.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&posts)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func ResetJSON() (*[]Post, error) {
	posts := make([]Post, 0)
	for i := 0; i < 5; i++ {
		user := User{
			CreatedBy:   usernames[rand.Intn(len(usernames))],
			TimeCreated: time.Now(),
		}
		rand := rand.Intn(len(messages))
		post := Post{
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
		return nil, err
	}

	file, err := os.Create("./posts.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return nil, err
	}
	fmt.Println("Profiles data written to ", file.Name())
	return &posts, nil
}
