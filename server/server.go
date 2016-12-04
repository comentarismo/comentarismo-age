package server

import (
	"fmt"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/gorilla/pat"
	"log"
	"net/http"
	"os"
	"comentarismo-age/age"
	"time"
)

var (
	router *pat.Router
)

type WebError struct {
	Error string
}

func init() {
	//
	if age.LEARNAGE == "true" {
		//will train with world know age words
		var start = 1950
		var end = 2012
		log.Println("Will start server on learning mode")
		year := time.Now().Year()

		done := make(chan bool, end - start)
		for i := start; i <= end; i++ {
			targetFile := fmt.Sprintf("/age/en/yob%d.txt", i)
			age_range := ""
			thisage := year - i

			if (thisage >= 18 && thisage <= 24) {
				age_range = "18_24"
			} else if (thisage >= 25 && thisage <= 34) {
				age_range = "25_34"
			} else if (thisage >= 35 && thisage <= 44) {
				age_range = "35_44"
			} else if (thisage >= 45 && thisage <= 54) {
				age_range = "45_54"
			} else if (thisage >= 55 && thisage <= 64) {
				age_range = "55_64"
			} else {
				log.Println("Will skip year outside interest age range --> ", i)
				done <- true
				continue
			}
			log.Println("Year ", i, " ,age ",thisage, " ,classified as ", age_range, " ,Will learn ", targetFile)

			go age.StartLanguageAge(age_range, targetFile, done)
		}
		go func() {
			for j := start; j <= end; j++ {
				targetFile := fmt.Sprintf("/age/en/yob%d.txt", j)
				log.Println("Finished learning ", <-done, targetFile)
			}
		}()

	}
}

//NewServer return pointer to new created server object
func NewServer(Port string) *http.Server {
	router = InitRouting()
	return &http.Server{
		Addr:    ":" + Port,
		Handler: router,
	}
}

//StartServer start and listen @server
func StartServer(Port string) {
	log.Println("Starting server")
	s := NewServer(Port)
	fmt.Println("Server starting --> " + Port)
	err := gracehttp.Serve(
		s,
	)
	if err != nil {
		log.Fatalln("Error: %v", err)
		os.Exit(0)
	}

}

func InitRouting() *pat.Router {

	r := pat.New()

	/** bayes classifier age **/
	r.Post("/age", AgeHandler)
	r.Post("/revoke", RevokeAgeHandler)
	r.Post("/report", ReportAgeHandler)

	return r
}
