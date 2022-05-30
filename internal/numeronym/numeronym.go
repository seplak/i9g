package numeronym

import (
	"fmt"
	"strings"
)

type Numeronym struct {
	Text string
}

func (n Numeronym) Convert() string {
	if strings.Contains(n.Text, " ") {
		split := strings.Split(n.Text, " ")
		var sentence []string
		for _, word := range split {
			sentence = append(sentence, createNumeronym(word))
		}
		return strings.Join(sentence, " ")
	}

	return createNumeronym(n.Text)
}

func createNumeronym(word string) string {
	// i don't make the rules but this seems reasonable
	if len(word) < 4 {
		return word
	}
	split := strings.Split(word, "")
	numWord := fmt.Sprintf("%s%d%s", split[0], len(split[1:len(split)-1]), split[len(split)-1])
	return numWord
}

