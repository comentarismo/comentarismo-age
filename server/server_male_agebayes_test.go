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

func TestMaleAgeHandler(t *testing.T) {

	testflight.WithServer(server.InitRouting(), func(r *testflight.Requester) {

		Convey("Should Learn female names in english and report age female for a female name", t, func() {
			targetWord := "matt"
			response := r.Post("/report?lang=en&age=35_44", testflight.FORM_ENCODED, "text="+targetWord)
			log.Println(response.Body)
			assert.Equal(t, 200, response.StatusCode)

			targetWord = "matthew"
			response = r.Post("/report?lang=en&age=35_44", testflight.FORM_ENCODED, "text="+targetWord)
			log.Println(response.Body)
			assert.Equal(t, 200, response.StatusCode)

			targetWord = "hank"
			response = r.Post("/report?lang=en&age=35_44", testflight.FORM_ENCODED, "text="+targetWord)
			log.Println(response.Body)
			assert.Equal(t, 200, response.StatusCode)

			targetWord = "mark"
			response = r.Post("/report?lang=en&age=35_44", testflight.FORM_ENCODED, "text="+targetWord)
			log.Println(response.Body)
			assert.Equal(t, 200, response.StatusCode)

			targetWord = "edward"
			response = r.Post("/report?lang=en&age=35_44", testflight.FORM_ENCODED, "text="+targetWord)
			log.Println(response.Body)
			assert.Equal(t, 200, response.StatusCode)

			targetWord = "henry"
			response = r.Post("/report?lang=en&age=35_44", testflight.FORM_ENCODED, "text="+targetWord)
			log.Println(response.Body)
			assert.Equal(t, 200, response.StatusCode)

			targetWord = "charlie"
			response = r.Post("/report?lang=en&age=35_44", testflight.FORM_ENCODED, "text="+targetWord)
			log.Println(response.Body)
			assert.Equal(t, 200, response.StatusCode)

			targetWord = "ben"
			response = r.Post("/report?lang=en&age=35_44", testflight.FORM_ENCODED, "text="+targetWord)
			log.Println(response.Body)
			assert.Equal(t, 200, response.StatusCode)

			//now try with a name
			textTarget := "charlie"
			response = r.Post("/age?lang=en", testflight.FORM_ENCODED, "text="+textTarget)

			So(response.StatusCode, ShouldEqual, 200)
			So(len(response.Body), ShouldBeGreaterThan, 0)

			log.Println(response.Body) //{"code":200,"error":"","spam":false}

			ageReport := age.AgeReport{}
			err := json.Unmarshal(response.RawBody, &ageReport)
			So(err, ShouldBeNil)

			So(ageReport.Error, ShouldBeBlank)
			So(ageReport.Code, ShouldEqual, 200)
			So(ageReport.Age, ShouldEqual, "35_44")

			//now try with another male name
			textTarget = "henry"
			response = r.Post("/age?lang=en", testflight.FORM_ENCODED, "text="+textTarget)

			So(response.StatusCode, ShouldEqual, 200)
			So(len(response.Body), ShouldBeGreaterThan, 0)

			log.Println(response.Body) //{"code":200,"error":"","age":true}

			ageReport = age.AgeReport{}
			err = json.Unmarshal(response.RawBody, &ageReport)
			So(err, ShouldBeNil)

			So(ageReport.Error, ShouldBeBlank)
			So(ageReport.Code, ShouldEqual, 200)
			So(ageReport.Age, ShouldEqual, "35_44")

			//now revoke the name
			textTarget = "henry"
			response = r.Post("/revoke?lang=en&age=35_44", testflight.FORM_ENCODED, "text="+textTarget)

			So(response.StatusCode, ShouldEqual, 200)
			So(len(response.Body), ShouldBeGreaterThan, 0)
			log.Println(response.Body) //{"code":200,"error":"","age":true}

			//now it should not be that age anymore
			textTarget = "henry"
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
