package age_test

import (
	"testing"
	"comentarismo-age/age"
	"log"
)

//* condition: (now().year - 18) to (now().year - 24) ->eg: 1998 to 1992 -->map: bayes:18_24
//* condition: (now().year - 25) to (now().year - 34) ->eg: 1991 to 1982 -->map: bayes:25_34
//* condition: (now().year - 35) to (now().year - 44) ->eg: 1981 to 1972 -->map: bayes:35_44
//* condition: (now().year - 45) to (now().year - 54) ->eg: 1971 to 1962 -->map: bayes:45_54
//* condition: (now().year - 55) to (now().year - 64) ->eg: 1961 to 1952 -->map: bayes:55_64

func TestClassifyAgeEn(t *testing.T) {
	log.Println("Will start server on learning mode, default to English. ")

	lang := "en"
	age.Train("18_24", "rachael",lang)
	age.Train("25_34", "hannah",lang)
	age.Train("25_34", "norah",lang)
	age.Train("18_24", "claire",lang)
	age.Train("25_34", "lauren",lang)
	age.Train("18_24", "marsha",lang)

	//FEMALE

	targetWord := "rachael"
	class := age.Classify(targetWord,lang)
	if class != "18_24" {
		t.Errorf("Classify failed, word (%s) should be 18_24, result: %s", targetWord, class)
	}

	targetWord = "hannah"
	class = age.Classify(targetWord,lang)
	if class != "25_34" {
		t.Errorf("Classify failed, word (%s) should be 25_34, result: %s", targetWord, class)
	}

	targetWord = "norah"
	class = age.Classify(targetWord,lang)
	if class != "25_34" {
		t.Errorf("Classify failed, word (%s) should be 25_34, result: %s", targetWord, class)
	}

	targetWord = "claire"
	class = age.Classify(targetWord,lang)
	if class != "18_24" {
		t.Errorf("Classify failed, word (%s) should be 18_24, result: %s", targetWord, class)
	}

	targetWord = "lauren"
	class = age.Classify(targetWord,lang)
	if class != "25_34" {
		t.Errorf("Classify failed, word (%s) should be 25_34, result: %s", targetWord, class)
	}

	targetWord = "marsha"
	class = age.Classify(targetWord,lang)
	if class != "18_24" {
		t.Errorf("Classify failed, word (%s) should be bad, result: %s", targetWord, class)
	}

	//MALE

	age.Train("18_24", "matt",lang)
	age.Train("25_34", "matthew",lang)
	age.Train("35_44", "hank",lang)
	age.Train("45_54", "mark",lang)
	age.Train("18_24", "edward",lang)
	age.Train("18_24", "henry",lang)
	age.Train("25_34", "charlie",lang)
	age.Train("35_44", "ben",lang)

	targetWord = "matt"
	class = age.Classify(targetWord,lang)
	if class != "18_24" {
		t.Errorf("Classify failed, word (%s) should be 18_24, result: %s", targetWord, class)
	}

	targetWord = "matthew"
	class = age.Classify(targetWord,lang)
	if class != "25_34" {
		t.Errorf("Classify failed, word (%s) should be 25_34, result: %s", targetWord, class)
	}
	targetWord = "hank"
	class = age.Classify(targetWord,lang)
	if class != "35_44" {
		t.Errorf("Classify failed, word (%s) should be 35_44, result: %s", targetWord, class)
	}
	targetWord = "mark"
	class = age.Classify(targetWord,lang)
	if class != "45_54" {
		t.Errorf("Classify failed, word (%s) should be 45_54, result: %s", targetWord, class)
	}
	targetWord = "edward"
	class = age.Classify(targetWord,lang)
	if class != "18_24" {
		t.Errorf("Classify failed, word (%s) should be 18_24, result: %s", targetWord, class)
	}
	targetWord = "henry"
	class = age.Classify(targetWord,lang)
	if class != "18_24" {
		t.Errorf("Classify failed, word (%s) should be 18_24, result: %s", targetWord, class)
	}
	targetWord = "charlie"
	class = age.Classify(targetWord,lang)
	if class != "25_34" {
		t.Errorf("Classify failed, word (%s) should be 25_34, result: %s", targetWord, class)
	}
	targetWord = "ben"
	class = age.Classify(targetWord,lang)
	if class != "35_44" {
		t.Errorf("Classify failed, word (%s) should be 35_44, result: %s", targetWord, class)
	}

}

func init() {
	age.Flush()
}
