package cupfmc

import (
	"strings"
	"time"
)

func GetDelays(text string) []time.Duration {
	var result []time.Duration
	text = strings.ToTitle(text)
	alphabet := GetAlphabet()
	splitted := strings.Split(text, " ")
	for i := 0; i < len(splitted); i++ {
		if i != 0 {
			result = append(result, Word)
		}
		for z := 0; z < len(splitted[i]); z++ {
			res, ok := alphabet[string(splitted[i][z])]
			if ok {
				for y := 0; y < len(res); y++ {
					if y != 0 {
						result = append(result, Delay)
					}
					result = append(result, res[y])
				}
				if z == len(splitted[i])-1 {
				} else {
					result = append(result, Sign)
				}
			}
		}
	}

	return result
}
