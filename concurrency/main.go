package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

// Result -
type Result struct {
	Error    error
	Todo     *Todo
	Todos    []*Todo
	Comment  *Comment
	Comments []*Comment
	User     *User
	Users    []*User
	Album    *Album
	Albums   []*Album
	Post     *Post
	Posts    []*Post
	Photo    *Photo
	Photos   []*Photo
}

func parse(response []byte) *Result {
	var todo Todo
	err := json.Unmarshal(response, &todo)
	if err != nil {
		return &Result{
			Error:   err,
			Todo:    nil,
			Comment: nil,
		}
	}
	return &Result{
		Error: nil,
		Todo:  &todo,
	}
}

func parseComment(response []byte) *Result {
	var comment Comment
	err := json.Unmarshal(response, &comment)
	if err != nil {
		return &Result{
			Error:   err,
			Comment: nil,
		}
	}
	return &Result{
		Error:   nil,
		Comment: &comment,
	}
}

// 'https://jsonplaceholder.typicode.com/todos/1'
var sampleResponse = `{"userId": 1, "id": 1, "title": "delectus aut autem", "completed": false}`

var (
	totalPhotos   = 5000
	totalComments = 500
	totalPosts    = 100
	totalAlbums   = 100
	totalTodos    = 200
	totalUsers    = 10

	baseURL = "http://localhost:3000"
	//baseURL        = "https://jsonplaceholder.typicode.com"
	todoURLBase    = fmt.Sprintf("%s/todos", baseURL)
	commentURLBase = fmt.Sprintf("%s/comments", baseURL)
	postsURLBase   = fmt.Sprintf("%s/posts", baseURL)
	usersURLBase   = fmt.Sprintf("%s/users", baseURL)
	photosURLBase  = fmt.Sprintf("%s/photos", baseURL)
	albumsURLBase  = fmt.Sprintf("%s/albums", baseURL)
)

func main() {
	//concurrency := 4
	start := time.Now()

	// fetchAllArtifactsSync()
	//fetchAllArtifactsWithSingleGoRoutines()
	fetchItems()
	elapsed := time.Since(start)
	log.Printf("Requests took: %s", elapsed)
}

func fetchItems() {
	log.Printf("Started retrieving: %d photos.", totalPhotos)
	var wg sync.WaitGroup
	wg.Add(6)

	go func() {
		defer wg.Done()
		todoRes := fetchTodos(totalTodos, 5, fetchSingleTodoSync)
		if todoRes != nil {
			log.Printf("Retrieved %d todos.", len(todoRes.Todos))
		}
	}()

	go func() {
		defer wg.Done()
		photoRes := fetchPhotos(totalPhotos, 5, fetchSinglePhotoSync)
		if photoRes != nil {
			log.Printf("Retrieved %d photos.", len(photoRes.Photos))
		}
	}()

	go func() {
		defer wg.Done()
		albumRes := fetchAlbums(totalAlbums, 5, fetchSingleAlbumSync)
		if albumRes != nil {
			log.Printf("Retrieved %d albums.", len(albumRes.Albums))
		}
	}()

	go func() {
		defer wg.Done()
		commentRes := fetchComments(totalComments, 5, fetchSingleCommentSync)
		if commentRes != nil {
			log.Printf("Retrieved %d comments.", len(commentRes.Comments))
		}
	}()

	go func() {
		defer wg.Done()
		userRes := fetchUsers(totalUsers, 5, fetchSingleUserSync)
		if userRes != nil {
			log.Printf("Retrieved %d users.", len(userRes.Users))
		}
	}()
	go func() {
		defer wg.Done()
		postsRes := fetchPosts(totalPosts, 5, fetchSinglePostSync)
		if postsRes != nil {
			log.Printf("Retrieved %d posts.", len(postsRes.Posts))
		}
	}()

	wg.Wait()
}

func fetchAllArtifactsWithSingleGoRoutines() {
	var wg sync.WaitGroup
	wg.Add(6)
	log.Printf("Started fetching individually with single go routines.")

	go func() {
		defer wg.Done()
		commentsRes := <-fetchUrls(totalComments, fetchAllCommentsSync)
		if commentsRes != nil {
			if commentsRes.Error != nil {
				log.Printf("Fetching comments resulted in an error %v", commentsRes.Error.Error())
			} else {
				log.Printf("fetched %d comments ", len(commentsRes.Comments))
			}
		} else {
			log.Printf("results was nil")
		}
	}()

	go func() {
		defer wg.Done()
		postsRes := <-fetchUrls(totalPosts, fetchAllPostsSync)
		if postsRes != nil {
			if postsRes.Error != nil {
				log.Printf("Fetching posts resulted in an error %v", postsRes.Error.Error())
			} else {
				log.Printf("fetched %d posts ", len(postsRes.Posts))
			}
		} else {
			log.Printf("results was nil")
		}
	}()

	go func() {
		defer wg.Done()
		albumsRes := <-fetchUrls(totalAlbums, fetchAllAlbumsSync)
		if albumsRes != nil {
			if albumsRes.Error != nil {
				log.Printf("Fetching posts resulted in an error %v", albumsRes.Error.Error())
			} else {
				log.Printf("fetched %d albums", len(albumsRes.Albums))
			}
		} else {
			log.Printf("results was nil")
		}
	}()

	go func() {
		defer wg.Done()
		todosRes := <-fetchUrls(totalTodos, fetchAllTodosSync)
		if todosRes != nil {
			if todosRes.Error != nil {
				log.Printf("Fetching todos resulted in an error %v", todosRes.Error.Error())
			} else {
				log.Printf("fetched %d todos ", len(todosRes.Todos))
			}
		} else {
			log.Printf("results was nil")
		}
	}()

	go func() {
		defer wg.Done()
		usersRes := <-fetchUrls(totalUsers, fetchAllUsersSync)
		if usersRes != nil {
			if usersRes.Error != nil {
				log.Printf("Fetching users resulted in an error %v", usersRes.Error.Error())
			} else {
				log.Printf("fetched %d users ", len(usersRes.Users))
			}
		} else {
			log.Printf("results was nil")
		}
	}()

	go func() {
		defer wg.Done()
		photosRes := <-fetchUrls(totalPhotos, fetchAllPhotosSync)
		if photosRes != nil {
			if photosRes.Error != nil {
				log.Printf("Fetching photos resulted in an error %v", photosRes.Error.Error())
			} else {
				log.Printf("fetched %d photos", len(photosRes.Photos))
			}
		} else {
			log.Printf("results was nil")
		}
	}()
	wg.Wait()
}

func fetchUrls(size int, f func(int) *Result) <-chan *Result {
	out := make(chan *Result)
	go func() {
		out <- f(size)
		close(out)
	}()
	return out
}

// comments only
// 24.795822621s
// 30.484542731s
// 25.434060587s

// with posts
//  29.418193215s
// 32.566759872s
// 31.720287564s

// with albums
// 48.361736606s
// 36.859959839s
// 33.749999766s

// with todos
// 2m9.278809966s - 100 todos
// 45.15940894s - 100 todos
// 46.841035549s - 100 todos
// 46.227141831s - 200 todos
// 50.590853835s - 200 todos
// 50.579789924s - 200 todos
func fetchAllArtifactsSync() {
	log.Printf("started fetching all artifacts...")
	commentsRes := fetchAllCommentsSync(500)
	if commentsRes.Error != nil {
		log.Printf("Fetching comments resulted in an error %v", commentsRes.Error.Error())
	} else {
		log.Printf("fetched %d comments", len(commentsRes.Comments))
	}

	postsRes := fetchAllPostsSync(100)
	if postsRes.Error != nil {
		log.Printf("Fetching posts resulted in an error %v", postsRes.Error.Error())
	} else {
		log.Printf("fetched %d posts", len(postsRes.Posts))
	}

	albumsRes := fetchAllAlbumsSync(100)
	if albumsRes.Error != nil {
		log.Printf("Fetching posts resulted in an error %v", albumsRes.Error.Error())
	} else {
		log.Printf("fetched %d albums", len(albumsRes.Albums))
	}
	todosRes := fetchAllTodosSync(200)
	if todosRes.Error != nil {
		log.Printf("Fetching todos resulted in an error %v", todosRes.Error.Error())
	} else {
		log.Printf("fetched %d todos", len(todosRes.Todos))
	}
	usersRes := fetchAllUsersSync(10)
	if usersRes.Error != nil {
		log.Printf("Fetching users resulted in an error %v", usersRes.Error.Error())
	} else {
		log.Printf("fetched %d users", len(usersRes.Users))
	}
	photosRes := fetchAllPhotosSync(1000)
	if photosRes.Error != nil {
		log.Printf("Fetching photos resulted in an error %v", photosRes.Error.Error())
	} else {
		log.Printf("fetched %d photos", len(photosRes.Photos))
	}

	log.Printf("completed fetching all artifacts")
}

// 655.859287ms
func fetchSingleArtifacts() {
	postRes := fetchSinglePostSync(1)
	if postRes.Error != nil {
		log.Printf("Fetching post resulted in an error %v", postRes.Error.Error())
	} else {
		log.Printf(postRes.Post.ToString())
	}

	userRes := fetchSingleUserSync(1)
	if userRes.Error != nil {
		log.Printf("Fetching user resulted in an error %v", userRes.Error.Error())
	} else {
		log.Printf(userRes.User.ToString())
	}

	albumRes := fetchSingleAlbumSync(1)
	if albumRes.Error != nil {
		log.Printf("Fetching album resulted in an error %v", albumRes.Error.Error())
	} else {
		log.Printf(albumRes.Album.ToString())
	}

	photosRes := fetchSinglePhotoSync(1)
	if photosRes.Error != nil {
		log.Printf("Fetching photos resulted in an error %v", photosRes.Error.Error())
	} else {
		log.Printf(photosRes.Photo.ToString())
	}

	todoRes := fetchSingleTodoSync(1)
	if todoRes.Error != nil {
		log.Printf("Fetching todo resulted in an error %v", todoRes.Error.Error())
	} else {
		log.Printf(todoRes.Todo.ToString())
	}

	commentRes := fetchSingleCommentSync(1)
	if albumRes.Error != nil {
		log.Printf("Fetching comment resulted in an error %v", commentRes.Error.Error())
	} else {
		log.Printf(commentRes.Comment.ToString())
	}
}

// func fetchSingleTodoSync(id int) *Result {
// 	url := fmt.Sprintf("%s/%d", todoURLBase, id)
// 	body, err := fetch(url)
// 	if err != nil {
// 		log.Fatalf("error: %v", err.Error())
// 	}
// 	var todo Todo
// 	err = json.Unmarshal(body, &todo)
// 	if err != nil {
// 		return &Result{
// 			Error: err,
// 			Todo:  nil,
// 		}
// 	}
// 	return &Result{
// 		Error: nil,
// 		Todo:  &todo,
// 	}
// }
// func fetchSinglePhotosSync(id int) *Result {
// 	url := fmt.Sprintf("%s/%d", photosURLBase, id)
// 	body, err := fetch(url)
// 	if err != nil {
// 		log.Fatalf("error: %v", err.Error())
// 	}
// 	var photo Photo
// 	err = json.Unmarshal(body, &photo)
// 	if err != nil {
// 		return &Result{
// 			Error: err,
// 			Photo: nil,
// 		}
// 	}
// 	return &Result{
// 		Error: nil,
// 		Photo: &photo,
// 	}
// }

// func fetchSingleAlbumSync(id int) *Result {
// 	url := fmt.Sprintf("%s/%d", albumsURLBase, id)
// 	body, err := fetch(url)
// 	if err != nil {
// 		log.Fatalf("error: %v", err.Error())
// 	}
// 	var album Album
// 	err = json.Unmarshal(body, &album)
// 	if err != nil {
// 		return &Result{
// 			Error: err,
// 			Album: nil,
// 		}
// 	}
// 	return &Result{
// 		Error: nil,
// 		Album: &album,
// 	}
// }

func requestToDos(concurrency, size int) <-chan *Result {
	var wg sync.WaitGroup
	wg.Add(concurrency)

	// fan out tasks into the channel and create as many worker go routines
	// as the concurrency demands
	tasks := make(chan string, 100)
	go func() {
		for i := 1; i <= size; i++ {
			tasks <- fmt.Sprintf("%s/%d", todoURLBase, i)
		}
		close(tasks)
	}()

	// create workers and schedule closing results when all work is done.
	results := make(chan *Result, 100)
	go func() {
		wg.Wait()
		close(results)
	}()

	// create as many go routines as concurrency allows
	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()
			for t := range tasks {
				r, err := fetch(t)
				if err != nil {
					log.Printf("could not fetch %v: %v", t, err)
					continue
				}

				results <- parse(r)
			}
		}()
	}
	return results
}

func requestComments() {
	for i := 1; i <= 500; i++ {
		body, err := request(i, commentURLBase)
		if err != nil {
			log.Fatalf("error: %v", err.Error())
		}
		commentResult := parseComment(body)
		if commentResult != nil {
			if commentResult.Error != nil {
				log.Fatalf("error: %v", commentResult.Error.Error())
			}
			log.Printf("%s", commentResult.Comment.ToString())
		}
	}
}
func requestTodo(url string) *Result {
	body, err := fetch(url)
	if err != nil {
		log.Fatalf("error: %v", err.Error())
	}
	todoResult := parse(body)
	if todoResult != nil {
		if todoResult.Error != nil {
			log.Fatalf("error: %v", todoResult.Error.Error())
		}
		//log.Printf("%s", todoResult.Todo.ToString())
	}
	return todoResult
}

func requestTodos() {
	for i := 1; i <= 100; i++ {
		body, err := request(i, todoURLBase)
		if err != nil {
			log.Fatalf("error: %v", err.Error())
		}
		todoResult := parse(body)
		if todoResult != nil {
			if todoResult.Error != nil {
				log.Fatalf("error: %v", todoResult.Error.Error())
			}
			log.Printf("%s", todoResult.Todo.ToString())
		}
	}
}

func fetch(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("could not get %s: %v", url, err)
	}
	defer res.Body.Close()
	if res != nil {
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			if res.StatusCode == http.StatusTooManyRequests {
				return nil, fmt.Errorf("you are being rate limited")
			}

			return nil, fmt.Errorf("bad response from server: %s", res.Status)
		}
		resp, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}

	return nil, nil
}

func request(ID int, URLBase string) ([]byte, error) {
	url := fmt.Sprintf("%s/%d", URLBase, ID)
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("could not get %s: %v", url, err)
	}
	defer res.Body.Close()
	if res != nil {
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			if res.StatusCode == http.StatusTooManyRequests {
				return nil, fmt.Errorf("you are being rate limited")
			}

			return nil, fmt.Errorf("bad response from server: %s", res.Status)
		}
		resp, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}

	return nil, nil
}

// https://jsonplaceholder.typicode.com/albums
/*{
	"userId"
	"id"
	"title"
}*/

// https://jsonplaceholder.typicode.com/photos
/*
{
	albumId:
	id
	title
	url
	thumbnailUrl
}
*/

//https://jsonplaceholder.typicode.com/users
/*
{
"id": 1,
"name": "Leanne Graham",
"username": "Bret",
"email": "Sincere@april.biz",
"address": {
"street": "Kulas Light",
"suite": "Apt. 556",
"city": "Gwenborough",
"zipcode": "92998-3874",
"geo": {
"lat": "-37.3159",
"lng": "81.1496"
}
}
*/

// posts -> 100
// comments -> 100
// albums -> 100
// photos -> 5000
// todos -> 200
// users -> 10
