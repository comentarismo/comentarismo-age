package age_test

import (
	"comentarismo-age/age"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"runtime"
	"testing"
	"time"
)

func TestFemaleAgeHandler(t *testing.T) {

	Convey("Should Learn female names in english and report age female for a female name", t, func() {

		//if coldstart == true && serialized exist
		if age.LEARNAGE != "" {
			runtime.GOMAXPROCS(runtime.NumCPU())

			year := time.Now().Year()
			var start = year - 64
			var end = year - 18

			log.Println("Will start server on learning mode for (year, start, end) -> ", year, start, end)

			done := make(chan bool, end-start)
			for i := start; i <= end; i++ {
				targetFile := fmt.Sprintf("/en/yob%d.txt", i)
				age_range := ""
				thisage := year - i

				if thisage >= 18 && thisage <= 24 {
					age_range = "18_24"
				} else if thisage >= 25 && thisage <= 34 {
					age_range = "25_34"
				} else if thisage >= 35 && thisage <= 44 {
					age_range = "35_44"
				} else if thisage >= 45 && thisage <= 54 {
					age_range = "45_54"
				} else if thisage >= 55 && thisage <= 64 {
					age_range = "55_64"
				} else {
					log.Println("Will skip year outside interest age range --> ", i)
					done <- true
					continue
				}
				log.Println("Year ", i, " ,age ", thisage, " ,classified as ", age_range, " ,Will learn ", targetFile)

				go age.StartLanguageAge(age_range, targetFile, done)
			}
			for j := start; j <= end; j++ {
				<-done
			}
			//save serialized
			age.WriteToFile("classifier.serialized")
		} else {
			//read serialized
			_, err := age.LearnFromFile("classifier.serialized")
			So(err, ShouldBeNil)
		}

		//now ask question

		lang := "en"

		targetWord := "matt"
		class := age.Classify(targetWord, lang)
		if class != "18_24" {
			t.Errorf("Classify failed, word (%s) should be 18_24, result: %s", targetWord, class)
		}

		targetWord = "rachael"
		class = age.Classify(targetWord, lang)
		if class != "18_24" {
			t.Errorf("Classify failed, word (%s) should be 18_24, result: %s", targetWord, class)
		}

		targetWord = "hannah"
		class = age.Classify(targetWord, lang)
		if class != "35_44" {
			t.Errorf("Classify failed, word (%s) should be 35_44, result: %s", targetWord, class)
		}

		targetWord = "norah"
		class = age.Classify(targetWord, lang)
		if class != "18_24" {
			t.Errorf("Classify failed, word (%s) should be 18_24, result: %s", targetWord, class)
		}

		targetWord = "claire"
		class = age.Classify(targetWord, lang)
		if class != "18_24" {
			t.Errorf("Classify failed, word (%s) should be 18_24, result: %s", targetWord, class)
		}

		targetWord = "lauren"
		class = age.Classify(targetWord, lang)
		if class != "18_24" {
			t.Errorf("Classify failed, word (%s) should be 18_24, result: %s", targetWord, class)
		}

		targetWord = "marsha"
		class = age.Classify(targetWord, lang)
		if class != "18_24" {
			t.Errorf("Classify failed, word (%s) should be 18_24, result: %s", targetWord, class)
		}

		//MALE

		targetWord = "hank"
		class = age.Classify(targetWord, lang)
		if class != "18_24" {
			t.Errorf("Classify failed, word (%s) should be 18_24, result: %s", targetWord, class)
		}
		targetWord = "mark"
		class = age.Classify(targetWord, lang)
		if class != "18_24" {
			t.Errorf("Classify failed, word (%s) should be 18_24, result: %s", targetWord, class)
		}
		targetWord = "edward"
		class = age.Classify(targetWord, lang)
		if class != "18_24" {
			t.Errorf("Classify failed, word (%s) should be 18_24, result: %s", targetWord, class)
		}
		targetWord = "ben"
		class = age.Classify(targetWord, lang)
		if class != "18_24" {
			t.Errorf("Classify failed, word (%s) should be 18_24, result: %s", targetWord, class)
		}
	})

}

func init() {
	age.Flush()
}
