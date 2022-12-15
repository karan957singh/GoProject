package models

var GenreSchema = `create table genres(
	id int(10) auto_increment not null,
	name varchar(50),
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP null default NULL,
	deleted_at TIMESTAMP null default NULL,
	primary key(id));`

var MoviesSchema = `create table movies(
	id int(10) auto_increment,
	genre_id int(10),
	title varchar(50),
	numberInStock int(3),
	dailyRentalRate float(5,2),
	liked tinyInt default 0,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP null default NULL,
	deleted_at TIMESTAMP null default NULL,
	primary key(id));`

// _ = db.MustExec(models.GenreSchema)
// _ = db.MustExec(models.MoviesSchema)

// var timme = time.Now()
// params := r.URL.Query()
// _, err = db.Exec("update movies set deleted_at=?  where id=?", timme, 5)

//db := getDbConnection()
// db.MustExec(models.MoviesSchema)

// paramsBody := r.Body
// 	paramB, _ := ioutil.ReadAll(paramsBody)
// 	bodyString := string(paramB)
// 	fmt.Println(bodyString)
