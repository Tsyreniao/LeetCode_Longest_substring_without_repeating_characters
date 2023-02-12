package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Values struct {
	Str       string
	Answer    string
	LenAnswer int
	IsExist   bool
}

var val Values

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	fmt.Println("\"" + val.Str + "\" \"" + val.Answer + "\"")

	tmpl.ExecuteTemplate(w, "index", val)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	val.Str = r.FormValue("Str")
	val.IsExist = true
	val.Answer = lengthOfLongestSubstring(val.Str)
	val.LenAnswer = len(val.Answer)

	http.Redirect(w, r, "/", http.StatusFound) //status 302
}

func lengthOfLongestSubstring(s string) string {
	var answer string
	var newAnswer string
	var finalAnswer string
	var i = 1
	var j = 0

	if len(s) > 0 {
		answer = string(s[0])
		finalAnswer = answer
	}
	for i < len(s) {
		answer += string(s[i])
		j = 0
		newAnswer = ""
		for j < (len(answer) - 1) {
			if answer[j] == s[i] {
				newAnswer = ""
			} else {
				newAnswer += string(answer[j])
			}
			j++
		}
		newAnswer += string(answer[j])
		answer = newAnswer
		if len(answer) > len(finalAnswer) {
			finalAnswer = answer
		}
		i++
	}
	return finalAnswer
}

func main() {
	fmt.Println("Listening on port :3000...")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/save", saveHandler)

	http.ListenAndServe(":3000", nil)
}
