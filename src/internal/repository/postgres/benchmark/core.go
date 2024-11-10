package benchmark

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"golang.org/x/exp/rand"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"src/internal/model"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}

func benchCreateClientGorm(repo IUserRepositoryGorm, n int) func(b *testing.B) {
	return func(b *testing.B) {
		b.N = n
		for i := 0; i < b.N; i++ {
			rand.Seed(uint64(time.Now().UnixNano()))
			name := randomString(5)
			surname := randomString(10)
			err := repo.Create(context.TODO(), &model.User{
				Name:     name,
				Surname:  surname,
				Email:    name + "@gmail.com",
				Role:     model.UserRoleCustomer,
				Password: "123456",
			})
			if err != nil {
				panic(err)
			}
		}
	}
}

func benchGetClientGorm(repo IUserRepositoryGorm, n int) func(b *testing.B) {
	return func(b *testing.B) {
		b.N = n
		for i := 0; i < b.N; i++ {
			_, err := repo.GetUserByID(context.TODO(), 1)
			if err != nil {
				panic(err)
			}
		}
	}
}

func benchCreateClientSquirrel(repo IUserRepositorySquirrel, n int) func(b *testing.B) {
	return func(b *testing.B) {
		b.N = n
		for i := 0; i < b.N; i++ {
			rand.Seed(uint64(time.Now().UnixNano()))
			name := randomString(5)
			surname := randomString(10)
			err := repo.Create(context.TODO(), &model.User{
				Name:     name,
				Surname:  surname,
				Email:    name + "@gmail.com",
				Role:     model.UserRoleCustomer,
				Password: "123456",
			})
			if err != nil {
				panic(err)
			}
		}
	}
}

func benchGetClientGormSquirrel(repo IUserRepositorySquirrel, n int) func(b *testing.B) {
	return func(b *testing.B) {
		b.N = n
		for i := 0; i < b.N; i++ {
			_, err := repo.GetUserByID(context.TODO(), 1)
			if err != nil {
				panic(err)
			}
		}
	}
}

const (
	USER     = "postgres"
	PASSWORD = "admin"
	DBNAME   = "Shop"
)

func SetupTestDatabaseGorm() (testcontainers.Container, *gorm.DB, error) {
	containerReq := testcontainers.ContainerRequest{
		Image:        "postgres:16",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
		Env: map[string]string{
			"POSTGRES_DB":       DBNAME,
			"POSTGRES_PASSWORD": PASSWORD,
			"POSTGRES_USER":     USER,
		},
	}

	dbContainer, _ := testcontainers.GenericContainer(
		context.Background(),
		testcontainers.GenericContainerRequest{
			ContainerRequest: containerReq,
			Started:          true,
		})

	host, _ := dbContainer.Host(context.Background())
	port, _ := dbContainer.MappedPort(context.Background(), "5432")

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port.Int(), USER, PASSWORD, DBNAME)
	pureDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, nil, fmt.Errorf("gorm open: %w", err)
	}

	text, err := os.ReadFile("../../../sql/init/init.sql")
	if err != nil {
		return nil, nil, fmt.Errorf("read file: %w", err)
	}

	if err := pureDB.Exec(string(text)).Error; err != nil {
		return nil, nil, fmt.Errorf("exec: %w", err)
	}

	return dbContainer, pureDB, nil
}

func SetupTestDatabase() (testcontainers.Container, *sql.DB) {
	containerReq := testcontainers.ContainerRequest{
		Image:        "postgres:16",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
		Env: map[string]string{
			"POSTGRES_DB":       DBNAME,
			"POSTGRES_PASSWORD": PASSWORD,
			"POSTGRES_USER":     USER,
		},
	}

	dbContainer, _ := testcontainers.GenericContainer(
		context.Background(),
		testcontainers.GenericContainerRequest{
			ContainerRequest: containerReq,
			Started:          true,
		})

	host, _ := dbContainer.Host(context.Background())
	port, _ := dbContainer.MappedPort(context.Background(), "5432")

	dsnPGConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port.Int(), USER, PASSWORD, DBNAME)
	db, err := sql.Open("pgx", dsnPGConn)
	if err != nil {
		return dbContainer, nil
	}

	err = db.Ping()
	if err != nil {
		return dbContainer, nil
	}
	db.SetMaxOpenConns(10)

	text, err := os.ReadFile("../../../sql/init/init.sql")
	if err != nil {
		return dbContainer, nil
	}

	if _, err := db.Exec(string(text)); err != nil {
		fmt.Println(err)
		return dbContainer, nil
	}

	return dbContainer, db
}

func Bench() []string {

}
