package phonetics

import (
	"regexp"
	"strings"
)

func encodeMetaphoneRu(word string) string {
	if word == "" {
		return ""
	}
	const alf = "ОЕАИУЭЮЯПСТРКЛМНБВГДЖЗЙФХЦЧШЩЫЕЇІЄ"
	word = strings.ToUpper(word)
	endString := map[string]string{
		"ОВСКИЙ$":               "@",
		"ЕВСКИЙ$":               "#",
		"ОВСКАЯ$":               "$",
		"ЕВСКАЯ$":               "%",
		"ИЕВА$|ЕЕВА$|ОВА$|ЕВА$": "9",
		"ИЕВ$|ЕЕВ$|ОВ$|ЕВ$":     "4",
		"НКО$":                  "3",
		"ИНА$":                  "1",
		"АЯ$":                   "6",
		"ИЙ$|ЫЙ$":               "7",
		"ЫХ$|ИХ$":               "5",
		"ИН$":                   "8",
		"ИК$|ЕК$":               "2",
		"УК$|ЮК$":               "0"}
	repString := map[string]string{
		"ТС|ДС":       "Ц",
		"[ЇІЕЁЭ]":     "И",
		"ЙО|ИО|ЙЕ|ИЕ": "И",
		"[ОЫЯ]":       "А",
		"[Ю]":         "У",
		"[Є]":         "Е"}

	dullReplacement := map[string]string{
		"Б": "П",
		"З": "С",
		"Д": "Т",
		"В": "Ф",
		"Г": "К"}
	result := ""
	wordRuneNew := []rune(" " + word)
	/*****Убераем лишнии символы *******/
	for _, key := range wordRuneNew {
		if strings.Index(alf, string(key)) > -1 {
			result += string(key)
		}
	}
	if len([]rune(result)) == 1 {
		return result
	}
	wordRuneNew = []rune(" " + result)
	result = ""
	// Исключение повторяющихся символов
	for i, str := range wordRuneNew {
		if i > 0 {
			if string(wordRuneNew[i-1:i]) != string(str) {
				result += string(str)
			}
		}
	}
	//Сжатие окончаний
	for key, value := range endString {
		var re = regexp.MustCompile(key)
		result = re.ReplaceAllString(result, value)
	}
	//Замена
	for key, value := range repString {
		var re = regexp.MustCompile(key)
		result = re.ReplaceAllString(result, value)
	}
	//Оглушение согласных в слабой позиции (перед не сонорными согласными и на конце слов)
	for key, value := range dullReplacement {
		var re = regexp.MustCompile(key + "([ПСТКБВГДЖЗФХЦЧШЩ])|" + key + "()$")
		result = re.ReplaceAllString(result, value+"$1")
	}
	wordRuneNew = []rune(" " + result)
	result = ""
	// Исключение повторяющихся символов
	for i, str := range wordRuneNew {
		if i > 0 {
			if string(wordRuneNew[i-1:i]) != string(str) {
				result += string(str)
			}
		}
	}
	return result
}
