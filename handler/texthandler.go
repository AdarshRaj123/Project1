package handler

import (
	"Project1/utils"
	"net/http"
	"sort"
	"strings"
)

type Word struct {
	Word  string
	Count int
}

type WordCountList []Word

func (p WordCountList) Len() int           { return len(p) }
func (p WordCountList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p WordCountList) Less(i, j int) bool { return p[i].Count > p[j].Count }

func CountWords(w http.ResponseWriter, r *http.Request) {

	body := struct {
		Text string `json :"text"`
	}{}

	if parseErr := utils.ParseBody(r.Body, &body); parseErr != nil {
		utils.RespondError(w, http.StatusBadRequest, parseErr, "failed to parse request body")
		return
	}
	text := body.Text

	fields := strings.FieldsFunc(text, func(r rune) bool {
		return !('a' <= r && r <= 'z' || 'A' <= r && r <= 'Z')
	})
	words := make(map[string]int)
	for _, field := range fields {
		words[strings.ToLower(field)]++
	}
	p := make(WordCountList, len(words))

	i := 0
	for k, v := range words {
		p[i] = Word{k, v}
		i++
	}

	sort.Sort(p)

	count := len(words)
	if count > 10 {
		count = 10
	}
	utils.RespondJSON(w, http.StatusOK, p[0:count])
}
