package bot

import (
	"log"
	"regexp"
	s "strings"
)

// checkWords find in text a key words
func checkWords(text string) bool {

	var keyWords = "раз два"

	clear, err := regexp.Compile("[^а-яА-Я]+")
	if err != nil {
		log.Fatal(err)
	}

	words := s.Fields(keyWords)
	clearText := s.Fields(clearWords(text, clear))

	for _, keyWord := range words {
		for _, target := range clearText {
			if s.EqualFold(s.ToLower(keyWord), s.ToLower(target)) {
				return true
			}
		}
	}
	return false
}

// clear text
func clearWords(text string, reg *regexp.Regexp) string {
	text = reg.ReplaceAllString(text, " ")
	return text
}
