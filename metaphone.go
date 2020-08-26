package phonetics

import (
	"errors"
	"log"
	"regexp"
	"strings"
)

func EncodeMetaphone(word string, lang ...string) string {
	_lang := "en"
	if len(lang) > 0 {
		_lang = lang[0]
	} else {
		var re = regexp.MustCompile("^[A-Za-z0-1]+$")
		if re.MatchString(word) {
			_lang = "en"
		} else {
			var re = regexp.MustCompile("^[А-Яа-я0-1ЇІЄїіє]+$")
			if re.MatchString(word) {
				_lang = "ru"
			}
		}
	}
	log.Println(_lang, word)
	switch _lang {
	case "ua", "ru":
		return encodeMetaphoneRu(word)
	case "en":
		return encodeMetaphoneEn(word)
	default:
		panic(errors.New("Lang (" + _lang + ") not support"))
	}
}

func EncodeMetaphoneWords(word string, lang ...string) []string {
	_lang := "en"
	if len(lang) > 0 {
		_lang = lang[0]
	} else {
		var re = regexp.MustCompile("^[A-Za-z0-1 ]+$")
		if re.MatchString(word) {
			_lang = "en"
		} else {
			var re = regexp.MustCompile("^[А-Яа-я0-1ЇІЄїіє ]+$")
			if re.MatchString(word) {
				_lang = "ru"
			}
		}
	}

	words := strings.Split(word, " ")
	var wordsReturn []string

	switch _lang {
	case "ua", "ru":
		for _, value := range words {
			wordsReturn = append(wordsReturn, encodeMetaphoneRu(value))
		}
	case "en":
		for _, value := range words {
			wordsReturn = append(wordsReturn, encodeMetaphoneEn(value))
		}
	default:
		panic(errors.New("Lang (" + _lang + ") not support"))
	}
	return wordsReturn
}
