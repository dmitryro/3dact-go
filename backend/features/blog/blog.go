package blog

import ("net/http",
        "github.com/go-chi/chi",
        "github.com/go-chi/render")

type Post struct {
    Id int `json:"id"`
    Author_id int `json:"author_id"`
    Title string `json:"title"`
    Body string `json:"body"`
}

type Comment struct {
    Post_id int `json:"post_id"`
    Author_id int `json:"author_id"`
    Body string `json:"body"`
}

func Routes() *chi.Mux {

}

func GetAPost(w http.ResponseWriter, r *http.Request) {

}

func DeletePost(w http.ResponseWriter, r *http.Request) {

}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {

}

func CreatePost(w http.ResponseWriter, r *http.Request) {

}

func main() {
	fmt.Println("vim-go")
}
