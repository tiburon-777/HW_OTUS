package hw03_frequency_analysis //nolint:golint,stylecheck
import (
	"regexp"
	"strings"
)

func Top10(str string) []string {
	template := regexp.MustCompile(`[A-Za-zА-Яа-я\-]+`)
	tmpArr := template.FindAllString(str, -1)
	tmpMap := calculate(tmpArr)
	var result []string
	for i := 0; i < 12; i++ {
		count := 0
		word := ""
		for tmpWord, tmpCount := range tmpMap {
			if tmpCount > count {
				count = tmpCount
				word = tmpWord
			}
		}
		if word != "" {
			result = append(result, word)
		}
		delete(tmpMap, word)
	}
	return result
}

func calculate(arr []string) map[string]int {
	result := make(map[string]int)
	for _, v := range arr {
		if v != "" && v != "-" {
			_, ok := result[strings.ToLower(v)]
			if ok {
				result[strings.ToLower(v)]++
			} else {
				result[strings.ToLower(v)] = 1
			}
		}
	}
	return result
}
