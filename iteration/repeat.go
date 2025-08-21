package iteration

import "strings"

func Repeat(str string, freq int) string {
	var res strings.Builder
	for range freq {
		res.WriteString(str)
	}

	return res.String()
}

func main() {

}
