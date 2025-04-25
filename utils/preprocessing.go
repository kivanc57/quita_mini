package utils

import (
	"regexp"
	"strings"
	"github.com/aaaton/golem/v4"
	"github.com/aaaton/golem/v4/dicts/en"
)

func GetFreqMap(corpus string, lemmatize bool) map[string]int {
	freqMap := make(map[string]int)

	lemmatizer, _ := golem.New(en.New())
	re := regexp.MustCompile(`[^a-z]+`)


	tokens := strings.Fields(corpus)
	for _, token := range tokens {
		// lowercase + replace non-letter
		token = strings.ToLower(token)
		token = re.ReplaceAllString(token, "")
	
		if token == "" {
			continue
		}
	
		if lemmatize {
			token = lemmatizer.Lemma(token)
		}
	
		freqMap[token]++
	}
	return freqMap
}
