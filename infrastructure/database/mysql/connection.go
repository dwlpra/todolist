package mysql

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dwlpra/todolist/domain/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateConnection() *gorm.DB {

	// docker run -e MYSQL_HOST=172.17.0.1 -e MYSQL_USER=root -e MYSQL_PASSWORD=my-secret-pw -e MYSQL_DBNAME=todolist -e MYSQL_PORT=3306 -p 8090:3030 fiber-todo:v1

	// docker run --network="host" -e API_URL=http://172.17.0.3:3030 monsterup/devcode-unit-test-1
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DBNAME")
	// docker build -t dwlpra/todolist-api:v2 .
	// docker push dwlpra/todolist-api:v1
	// docker build -t fiber-todo:v1 .

	// user := "root"
	// pass := "my-secret-pw"
	// host := "127.0.0.1"
	// port := "3306"
	// dbName := "todolist"

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:            true,
		DisableAutomaticPing:   true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	// Get generic database object sql.DB to use its functions
	sqlDB, _ := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(1000)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(1000)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	db.AutoMigrate(&entity.Activity{})
	db.AutoMigrate(&entity.Todo{})

	return db

}
