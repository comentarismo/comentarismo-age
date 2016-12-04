package server_test

import (
	"comentarismo-age/age"
	"comentarismo-age/server"
	"encoding/json"
	"github.com/drewolson/testflight"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestFemaleAgeHandler(t *testing.T) {

	testflight.WithServer(server.InitRouting(), func(r *testflight.Requester) {

		Convey("Should Learn female names in english and report age range for a female name", t, func() {
			targetWord := "rachael"
			response := r.Post("/report?lang=en&age=18_24", testflight.FORM_ENCODED, "text="+targetWord)
			log.Println(response.Body)
			assert.Equal(t, 200, response.StatusCode)

			targetWord = "hannah"
			response = r.Post("/report?lang=en&age=18_24", testflight.FORM_ENCODED, "text="+targetWord)
			log.Println(response.Body)
			assert.Equal(t, 200, response.StatusCode)

			targetWord = "norah"
			response = r.Post("/report?lang=en&age=18_24", testflight.FORM_ENCODED, "text="+targetWord)
			log.Println(response.Body)
			assert.Equal(t, 200, response.StatusCode)

			targetWord = "claire"
			response = r.Post("/report?lang=en&age=18_24", testflight.FORM_ENCODED, "text="+targetWord)
			log.Println(response.Body)
			assert.Equal(t, 200, response.StatusCode)

			targetWord = "marsha"
			response = r.Post("/report?lang=en&age=18_24", testflight.FORM_ENCODED, "text="+targetWord)
			log.Println(response.Body)
			assert.Equal(t, 200, response.StatusCode)

			targetWord = "rachael"
			response = r.Post("/report?lang=en&age=18_24", testflight.FORM_ENCODED, "text="+targetWord)
			log.Println(response.Body)
			assert.Equal(t, 200, response.StatusCode)

			//now try with a name
			textTarget := "rachael"
			response = r.Post("/age?lang=en", testflight.FORM_ENCODED, "text="+textTarget)

			So(response.StatusCode, ShouldEqual, 200)
			So(len(response.Body), ShouldBeGreaterThan, 0)

			log.Println(response.Body) //{"code":200,"error":"","spam":false}

			ageReport := age.AgeReport{}
			err := json.Unmarshal(response.RawBody, &ageReport)
			So(err, ShouldBeNil)

			So(ageReport.Error, ShouldBeBlank)
			So(ageReport.Code, ShouldEqual, 200)
			So(ageReport.Age, ShouldEqual, "18_24")

			//now try with a 55_64 name
			textTarget = "marsha"
			response = r.Post("/age?lang=en", testflight.FORM_ENCODED, "text="+textTarget)

			So(response.StatusCode, ShouldEqual, 200)
			So(len(response.Body), ShouldBeGreaterThan, 0)

			log.Println(response.Body) //{"code":200,"error":"","age":true}

			ageReport = age.AgeReport{}
			err = json.Unmarshal(response.RawBody, &ageReport)
			So(err, ShouldBeNil)

			So(ageReport.Error, ShouldBeBlank)
			So(ageReport.Code, ShouldEqual, 200)
			So(ageReport.Age, ShouldEqual, "18_24")

			//now revoke the age name
			textTarget = "marsha"
			response = r.Post("/revoke?lang=en&age=18_24", testflight.FORM_ENCODED, "text="+textTarget)

			So(response.StatusCode, ShouldEqual, 200)
			So(len(response.Body), ShouldBeGreaterThan, 0)
			log.Println(response.Body) //{"code":200,"error":"","age":true}

			//now it should not be age anymore
			textTarget = "marsha"
			response = r.Post("/age?lang=en", testflight.FORM_ENCODED, "text="+textTarget)

			So(response.StatusCode, ShouldEqual, 200)
			So(len(response.Body), ShouldBeGreaterThan, 0)

			log.Println(response.Body)

			ageReport = age.AgeReport{}
			err = json.Unmarshal(response.RawBody, &ageReport)
			So(err, ShouldBeNil)

			So(ageReport.Error, ShouldBeBlank)
			So(ageReport.Code, ShouldEqual, 200)
			So(ageReport.Age, ShouldEqual, "18_24")
		})

	})
}

func init() {
	age.Flush()
}
