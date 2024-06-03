package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hrishikeshgoyal/quest/go-microservice-fk/repository/mysql"
	"github.com/hrishikeshgoyal/quest/go-microservice-fk/resource"
	"github.com/hrishikeshgoyal/quest/go-microservice-fk/service/impl"
	"log"
	"net/http"
	"os"
	"time"
)

func getUrl(config Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?pareseTime=true", config.Database.Username, &config.Database.Password, config.Database.Host, config.Database.Name)
}

const DNS = "root:rootroot@tcp(127.0.0.1:3306)/go-microservice-fk?parseTime=true"

type Config struct {
	Database struct {
		Host     string `json:"host"`
		Password string `json:"password"`
		Username string `json:"username"`
		Name     string `json:"name"`
	} `json:"database"`
	Host string `json:"host"`
	Port string `json:"port"`
}

func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	json.NewDecoder(configFile).Decode(&config)
	return config
}

func main() {

	r := mux.NewRouter()

	config := LoadConfiguration("config.json")

	//httpAddr := os.Getenv("HTTP_ADDR")
	httpAddr := "8080"
	log.Printf(getUrl(config))
	ur := mysql.NewUserRepository(DNS)
	//ur := sqlite.NewUserRepository(DNS)
	us := impl.NewUserService(ur)

	resource.RegisterUserResource(r, us)
	s := &http.Server{
		Addr:           ":" + httpAddr,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
