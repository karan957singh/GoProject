package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"./models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func getData(db *sqlx.DB, table string) *sql.Rows {

	var query string
	if table == "Genres" {
		query = "select * from genres;"
	}

	if table == "Movies" {
		query = "select * from movies;"
	}
	selectall, err := db.Query(query)
	// fmt.Println(selectal)
	if err != nil {
		panic(err.Error())
	}
	// fmt.Println(selectall)
	// var selectall *sql.Rows
	return selectall
}
func getDbConnection() *sqlx.DB {
	db, err := sqlx.Connect("mysql", "root:@tcp(127.0.0.1:3306)/react_app")
	if err != nil {
		panic(err.Error())
	}
	return db
}

type myData struct {
	Id int
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//to set response status
	// w.WriteHeader(500)
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	// fmt.Println(models.GenreSchema)
	// q := r.URL.Query()

	fmt.Println("handler called!     ")
	params := r.URL.Query()
	// paramsBody := r.Body
	// paramB, _ := ioutil.ReadAll(paramsBody)
	// bodyString := string(paramB)
	// fmt.Println(bodyString)
	// var data myData
	// urData := make(map[byte]string)
	// deco := json.NewDecoder(r.Body)
	// err := deco.Decode(&urData)
	// fmt.Println("urData:  ", urData)
	fmt.Println(params)
	db := getDbConnection()
	var timme = time.Now()
	if len(params["id"]) > 0 {
		for _, id := range params["id"] {
			fmt.Println(id)
			rowsAffected, _ := db.Exec("update movies set deleted_at=?  where id=?", timme, id)
			affCount, _ := rowsAffected.RowsAffected()
			fmt.Println("rows affected: ", affCount)
			if affCount == 0 {
				http.Error(w, "movie not found", 404)
				return
			}

			// rows, _ := db.Query("select * from movies where id=?", id)
			// if rows.Next() {
			// 	_, _ = db.Exec("update movies set deleted_at=?  where id=?", timme, id)
			// } else {
			// 	http.Error(w, "movie not found", 404)
			// 	return
			// }
		}
	}
}
func getMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//to set response status
	// w.WriteHeader(500)
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	// fmt.Println(models.GenreSchema)
	// q := r.URL.Query()

	fmt.Println("get movie handler called!     ")
	params := r.URL.Query()
	fmt.Println("URL params: ", params)
	db := getDbConnection()
	// var timme = time.Now()
	if len(params["id"]) > 0 {
		for _, id := range params["id"] {
			fmt.Println(id)
			rows, _ := db.Query("select * from movies where id=?", id)

			var movies models.Movies
			// moviesCol := []models.Movies{}
			for rows.Next() {
				rows.Scan(&movies.Id, &movies.Genre_id, &movies.Title, &movies.NumberInStock, &movies.DailyRentalRate, &movies.Liked, &movies.Created_at, &movies.Updated_at, &movies.Deleted_at)

				// moviesCol = append(moviesCol, movies)
			}
			if movies.Id == 0 {
				http.Error(w, "movie not found", 404)
				return
			}
			json.NewEncoder(w).Encode(movies)

			// rows, _ := db.Query("select * from movies where id=?", id)
			// if rows.Next() {
			// 	_, _ = db.Exec("update movies set deleted_at=?  where id=?", timme, id)
			// } else {
			// 	http.Error(w, "movie not found", 404)
			// 	return
			// }
		}
	}

	// fmt.Println(err)

	// tx, err := db.Begin()

	// some, err := tx.Prepare("update movies set deleted_at=curdate() where id=3")

	// res, err := some.Exec()
	// some.Close()
	// // defer db.Close()
	// rowsAffected, err := res.RowsAffected()
	// fmt.Println("rows affected: ", rowsAffected)

	// r.ParseForm()
	// postParams := r.Form.Get("5")

	// w.Write([]byte("doneeee\n"))
	// fmt.Fprintf(w, "all goood!")
	// io.WriteString(w, "test successful")
	// keyPassed := q.Get("id")

	// fmt.Println(data)
	// fmt.Fprintf(w, string(data))

	// fmt.Println(r.Method, data, keyPassed)
}

func saveMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//to set response status
	// w.WriteHeader(500)
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	// fmt.Println(models.GenreSchema)
	// q := r.URL.Query()

	fmt.Println("save movie handler called!     ")
	params := r.URL.Query()
	fmt.Println("URL params: ", params)
	var movies models.Movies
	paramsBody := r.Body
	// paramB, _ := ioutil.ReadAll(paramsBody)
	// paramB.Scan(&movies)
	json.NewDecoder(paramsBody).Decode(&movies)
	// json.NewEncoder(w).Encode(genreCol)
	// bodyString := string(movies)
	fmt.Println(movies)
	fmt.Println("Form params: ")
	db := getDbConnection()
	// var timme = time.Now()
	if len(params["id"]) > 0 {
		for _, id := range params["id"] {
			fmt.Println(id, movies.Title)
			// rows, _ := db.Query("select * from movies where id=?", id)

			// for rows.Next() {
			// 	rows.Scan(&movies.Id, &movies.Genre_id, &movies.Title, &movies.NumberInStock, &movies.DailyRentalRate, &movies.Liked, &movies.Created_at, &movies.Updated_at, &movies.Deleted_at)

			// 	// moviesCol = append(moviesCol, movies)
			// }

			rowsAffected, _ := db.Exec("update movies set genre_id=?, title=?, numberInStock=?, dailyRentalRate=?, liked=? where id=?", movies.Genre_id, movies.Title, movies.NumberInStock, movies.DailyRentalRate, movies.Liked, id)
			affCount, _ := rowsAffected.RowsAffected()
			fmt.Println("rows affected: ", affCount)
			if affCount == 0 {
				http.Error(w, "movie not found", 404)
				return
			}

			// if movies.Id == 0 {
			// 	http.Error(w, "movie not found", 404)
			// 	return
			// }
			json.NewEncoder(w).Encode(movies)

			// rows, _ := db.Query("select * from movies where id=?", id)
			// if rows.Next() {
			// 	_, _ = db.Exec("update movies set deleted_at=?  where id=?", timme, id)
			// } else {
			// 	http.Error(w, "movie not found", 404)
			// 	return
			// }
		}
	}
}

func main() {
	fmt.Println("Go MySQL Connection")

	http.HandleFunc("/get-genre", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		db := getDbConnection()
		defer db.Close()
		selectall := getData(db, "Genres")

		var genre models.Genres
		genreCol := []models.Genres{}
		for selectall.Next() {
			selectall.Scan(&genre.Id, &genre.Name, &genre.Created_at, &genre.Updated_at, &genre.Deleted_at)
			genreCol = append(genreCol, genre)
		}
		// fmt.Println(genreCol)
		json.NewEncoder(w).Encode(genreCol)
		fmt.Println("Query Successful")
	})
	http.HandleFunc("/get-movies", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		db := getDbConnection()
		defer db.Close()

		selectall := getData(db, "Movies")

		var movies models.Movies
		moviesCol := []models.Movies{}
		for selectall.Next() {
			selectall.Scan(&movies.Id, &movies.Genre_id, &movies.Title, &movies.NumberInStock, &movies.DailyRentalRate, &movies.Liked, &movies.Created_at, &movies.Updated_at, &movies.Deleted_at)
			moviesCol = append(moviesCol, movies)
		}
		json.NewEncoder(w).Encode(moviesCol)
		// fmt.Println("Query Successful", moviesCol)
	})
	http.HandleFunc("/delete-movie", handler)
	http.HandleFunc("/get-movie", getMovieHandler)
	http.HandleFunc("/save-movie", saveMovieHandler)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Connection successful")
}
