package mysql

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/pressly/goose"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var testdb *gorm.DB

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect docker %s", err)
	}

	runOptions := &dockertest.RunOptions{
		Repository: "mysql",
		Tag:        "8.0",
		Env: []string{
			"MYSQL_ROOT_PASSWORD=root",
			"MYSQL_DATABASE=mysqltest",
		},
		Cmd: []string{
			"mysqld",
			"--character-set-server=utf8mb4",
			"--collation-server=utf8mb4_unicode_ci",
		},
	}

	resource, err := pool.RunWithOptions(runOptions)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	dsn := fmt.Sprintf("root:root@tcp(127.0.0.1:%s)/mysqltest?charset=utf8mb4&parseTime=true", resource.GetPort("3306/tcp"))
	if err := pool.Retry(func() error {
		testdb, err = gorm.Open(mysql.New(mysql.Config{DSN: dsn}), &gorm.Config{})
		if err != nil {
			log.Println("gorm.Open()")
			return err
		}

		sqlDB, err := testdb.DB()
		if err != nil {
			log.Println("testdb.DB()")
			return err
		}

		return sqlDB.Ping()
	}); err != nil {
		log.Println("pool.Retry()")
		log.Fatalf("Could not connect to docker: %s", err)
	}
	defer func() {
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	}()

	// migration
	sqlDB, err := testdb.DB()
	if err != nil {
		log.Fatalf("Could not get sql.DB: %v", err)
	}

	if err := goose.SetDialect("mysql"); err != nil {
		log.Fatalf("Could not set dialect: %v", err)
	}

	if err := goose.Up(sqlDB, "../../../../../database/migrations"); err != nil {
		log.Fatalf("Could not run migrations: %v", err)
	}

	os.Exit(m.Run())

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}
