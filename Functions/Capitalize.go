package reloaded
import (
	"strings"
)

func Capitalize(word string) string {
	res := strings.Title(strings.ToLower(word))
	i := strings.Index(word, "_")
	if i == -1 {
		return res
	} else {
		res = strings.Replace(res, string(res[i+1]), strings.ToUpper(string(res[i+1])), i+1)
		return res
	}
}