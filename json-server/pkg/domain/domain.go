package domain

import "fmt"

// Todo -
type Todo struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// String -
func (t *Todo) String() string {
	return fmt.Sprintf("Todo{Id: %d, UserId: %d, Title: %s, Completed: %v}",
		t.ID, t.UserID, t.Title, t.Completed)
}

// Post -
type Post struct {
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// String -
func (p *Post) String() string {
	return fmt.Sprintf("Post{Id: %d, UserId: %d, Title: %s}", p.ID, p.UserID, p.Title)
}

// Album -
type Album struct {
	UserID int    `json:"userId,omitempty"`
	ID     int    `json:"id"`
	Title  string `json:"title,omitempty"`
}

// String -
func (a *Album) String() string {
	return fmt.Sprintf("Album{Id: %d, UserId: %d, Title: %s}", a.ID, a.UserID, a.Title)
}

// User -
type User struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Address  *Address `json:"address,omitempty"`
}

// Address -
type Address struct {
	Street      string            `json:"address"`
	Suite       string            `json:"suite,omitempty"`
	City        string            `json:"city,omitempty"`
	ZipCode     string            `json:"zipcode,omitempty"`
	GeoLocation map[string]string `json:"geo,omitempty"`
}

// String -
func (u *User) String() string {
	return fmt.Sprintf("User{Id: %d, Name: %s, Username: %s, Email: %s}", u.ID, u.Name, u.Username, u.Email)
}

// Comment -
type Comment struct {
	ID     int    `json:"id"`
	PostID int    `json:"postId"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

// String -
func (c *Comment) String() string {
	return fmt.Sprintf("Comment{Id: %d, UserId: %d, Title: %s, Completed: %v}", c.ID, c.PostID, c.Name, c.Email)
}

// Photo -
type Photo struct {
	AlbumID      int    `json:"albumId,omitemtpy"`
	ID           int    `json:"id,omitempty"`
	Title        string `json:"title,omitempty"`
	URL          string `json:"url,omitempty"`
	ThumbnailURL string `json:"thumbnailUrl,omitempty"`
}

// String -
func (p *Photo) String() string {
	return fmt.Sprintf("Photo{Id: %d, UserId: %d, Title: %s, URL: %s}", p.ID, p.AlbumID, p.Title, p.URL)
}
