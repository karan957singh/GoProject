package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Genres struct {
	Name string `json:"name"` //name
	Id   int    `json:"id"`   //id
	// name string
	// id   int
}

type Movies struct {
	Id              int     `json:"id"`
	Title           string  `json:"title"`
	Genre_id        int     `json:"genre_id"`
	NumberInStock   int     `json:"numberInStock"`
	DailyRentalRate float32 `json:"dailyRentalRate"`
	Liked           bool    `json:"liked"`
}

func getData(db *sql.DB, table string) *sql.Rows {
	selectall, err := db.Query("select * from " + table)
	if err != nil {
		panic(err.Error())
	}
	// fmt.Println("result of type: ", reflect.TypeOf(selectall))
	// defer selectall.Close()
	return selectall
}
func getDbConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/react_app")
	if err != nil {
		panic(err.Error())
	}
	return db
}

type myData struct {
	Id string
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//to set response status
	// w.WriteHeader(500)
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	q := r.URL.Query()

	var data myData
	deco := json.NewDecoder(r.Body)
	err := deco.Decode(&data)
	fmt.Println(err)

	r.ParseForm()
	postParams := r.Form.Get("5")

	w.Write([]byte("doneeee\n"))
	fmt.Fprintf(w, "all goood!")
	io.WriteString(w, "test successful")
	keyPassed := q.Get("id")

	fmt.Println(r.Method, data, postParams, keyPassed)

}

func main() {
	fmt.Println("Go MySQL Connection")

	http.HandleFunc("/get-genre", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		db := getDbConnection()
		defer db.Close()
		selectall := getData(db, "Genre")

		var genre Genres
		genreCol := []Genres{}
		for selectall.Next() {
			selectall.Scan(&genre.Id, &genre.Name)
			genreCol = append(genreCol, genre)
		}
		json.NewEncoder(w).Encode(genreCol)
		fmt.Println("Query Successful")
	})
	http.HandleFunc("/get-movies", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		db := getDbConnection()
		defer db.Close()
		selectall := getData(db, "Movies")

		var movies Movies
		moviesCol := []Movies{}
		for selectall.Next() {
			selectall.Scan(&movies.Id, &movies.Genre_id, &movies.Title, &movies.NumberInStock, &movies.DailyRentalRate, &movies.Liked)
			moviesCol = append(moviesCol, movies)
		}
		json.NewEncoder(w).Encode(moviesCol)
		fmt.Println("Query Successful")
	})
	http.HandleFunc("/delete-movie", handler)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Connection successful")
}
