package numeronym

import (
	"testing"
)

func TestNumeronym(t *testing.T) {
	numeronymWordTests := []struct{
		name      string
		numeronym Numeronym
		becomes   string
	}{
		{name: "kubernetes", numeronym: Numeronym{"Kubernetes"}, becomes: "K8s"},
		{name: "observability", numeronym: Numeronym{"observability"}, becomes: "o11y"},
		{name: "a short word", numeronym: Numeronym{"why"}, becomes: "why"},
		{name: "apostrophes count", numeronym: Numeronym{"wouldn't"}, becomes: "w6t"},
		{name: "words separated by spaces", numeronym: Numeronym{"i do observability with Kubernetes"}, becomes: "i do o11y w2h K8s"},
	}
	
	for _, tt := range numeronymWordTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.numeronym.Convert()
			if got != tt.becomes {
				assertStringEquals(got, tt.becomes, t)
			}
		})
	}
}


func assertStringEquals(got, want string, t *testing.T) {
	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}
