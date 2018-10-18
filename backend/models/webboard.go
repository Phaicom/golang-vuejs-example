package models

import "time"

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

func GetPosts() (*[]Post, error) {
	// get collection of post from db
	posts := new([]Post)
	// if err return (zero value, err)
	// else if add completed
	// return (completed []post, nil)
	return posts, nil
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
