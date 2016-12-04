package server

import (
	"comentarismo-age/age"
	"encoding/json"
	"log"
	"net/http"
)

func ReportAgeHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
	//log.Println(req.Form) // print information on server side.

	lang := req.URL.Query().Get("lang")
	//validate inputs
	if lang == "" {
		w.WriteHeader(http.StatusNotFound)
		jsonBytes, _ := json.Marshal(WebError{Error: "Missing lang"})
		w.Write(jsonBytes)
		return
	}

	gen := req.URL.Query().Get("age")
	if gen == "" {
		w.WriteHeader(http.StatusNotFound)
		jsonBytes, _ := json.Marshal(WebError{Error: "Missing age argument"})
		w.Write(jsonBytes)
		return
	}


	//log.Println("lang , ", lang)
	if lang != "pt" && lang != "en" && lang != "fr" && lang != "es" && lang != "it" && lang != "hr" && lang != "ru" {
		errMsg := "Error: SentimentHandler Language " + lang + " not yet supported, use lang={en|pt|es|it|fr|hr|ru} eg lang=en"
		log.Println(errMsg)
		jsonBytes, _ := json.Marshal(WebError{Error: errMsg})
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonBytes)
		return
	}

	text := req.Form["text"]
	reply := age.AgeReport{}

	//validate inputs
	if len(text) == 0 {
		errMsg := "Error: ReportSpamHandler text not defined, use eg: text=This Is not SPAM!!!"
		log.Println(errMsg)
		reply.Code = 404
		reply.Error = errMsg
		jsonBytes, _ := json.Marshal(reply)
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonBytes)
		return
	}

	log.Println("ReportSpamHandler, -->  ", gen, text[0], lang)
	age.Train(gen, text[0], lang)
	reply.Code = 200

	//marshal comment
	jsonBytes, err := json.Marshal(&reply)
	if err != nil {
		reply.Code = 404
		errMsg := "Error: ReportSpamHandler Marshal"
		log.Println(errMsg)
		reply.Error = errMsg
		jsonBytes, _ := json.Marshal(reply)
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonBytes)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func RevokeAgeHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	lang := req.URL.Query().Get("lang")

	//validate inputs
	if lang == "" {
		w.WriteHeader(http.StatusNotFound)
		jsonBytes, _ := json.Marshal(WebError{Error: "Missing lang"})
		w.Write(jsonBytes)
		return
	}

	gen := req.URL.Query().Get("age")
	if gen == "" {
		w.WriteHeader(http.StatusNotFound)
		jsonBytes, _ := json.Marshal(WebError{Error: "Missing age argument"})
		w.Write(jsonBytes)
		return
	}

	//log.Println("lang , ", lang)
	if lang != "pt" && lang != "en" && lang != "fr" && lang != "es" && lang != "it" && lang != "hr" && lang != "ru" {
		errMsg := "Error: SentimentHandler Language " + lang + " not yet supported, use lang={en|pt|es|it|fr|hr|ru} eg lang=en"
		log.Println(errMsg)
		jsonBytes, _ := json.Marshal(WebError{Error: errMsg})
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonBytes)
		return
	}

	text := req.Form["text"]
	reply := age.AgeReport{}

	//validate inputs
	if len(text) == 0 {
		errMsg := "Error: RevokeSpamHandler text not defined, use eg: text=This Is not SPAM!!!"
		log.Println(errMsg)
		reply.Code = 404
		reply.Error = errMsg
		jsonBytes, _ := json.Marshal(reply)
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonBytes)
		return
	}

	log.Println("RevokeSpamHandler, -->  ", text)

	age.Untrain(gen, text[0], lang)

	reply.Code = 200

	//marshal comment
	jsonBytes, err := json.Marshal(&reply)
	if err != nil {
		reply.Code = 404
		errMsg := "Error: RevokeSpamHandler Marshal"
		log.Println(errMsg)
		reply.Error = errMsg
		jsonBytes, _ := json.Marshal(reply)
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonBytes)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func AgeHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	//log.Println(req.Form)

	lang := req.URL.Query().Get("lang")

	//validate inputs
	if lang == "" {
		w.WriteHeader(http.StatusNotFound)
		jsonBytes, _ := json.Marshal(WebError{Error: "Missing lang"})
		w.Write(jsonBytes)
		return
	}

	//log.Println("lang , ", lang)
	if lang != "pt" && lang != "en" && lang != "fr" && lang != "es" && lang != "it" && lang != "hr" && lang != "ru" {
		errMsg := "Error: AgeHandler Language " + lang + " not yet supported, use lang={en|pt|es|it|fr|hr|ru} eg lang=en"
		log.Println(errMsg)
		jsonBytes, _ := json.Marshal(WebError{Error: errMsg})
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonBytes)
		return
	}

	text := req.Form["text"]
	reply := age.AgeReport{}

	//validate inputs
	if len(text) == 0 {
		reply.Code = 404
		errMsg := "Error: AgeHandler text not defined, use eg: name=John"
		log.Println(errMsg)
		reply.Error = errMsg
		jsonBytes, _ := json.Marshal(reply)
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonBytes)
		return
	}

	class := age.Classify(text[0], lang)
	reply.Code = 200
	reply.Age = class

	//marshal comment
	jsonBytes, err := json.Marshal(&reply)
	if err != nil {
		reply.Code = 404
		errMsg := "Error: AgeHandler Marshal"
		log.Println(errMsg)
		reply.Error = errMsg
		jsonBytes, _ := json.Marshal(reply)
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonBytes)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
