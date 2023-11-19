package utils

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func Capitalize(s string) string {
	capitalized := ""
	for i, c := range s {
		sym := string(c)
		if i == 0 {
			sym = strings.ToUpper(string(c))
		}
		capitalized = capitalized + sym
	}
	return capitalized
}

func DeleteSliceElement(slice []int64, index int) []int64 {
	return append(slice[:index], slice[index+1:]...)
}

func NumToColor(i int) string {
	switch i {
	case 1:
		return "⭐"
	case 2:
		return "⭐⭐"
	case 3:
		return "⭐⭐⭐"
	default:
		return ""
	}
}

func StructToBson(v interface{}) (doc *bson.M, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}
