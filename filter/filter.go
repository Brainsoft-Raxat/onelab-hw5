package filter

import (
	"reflect"
	"strings"
	"unicode"
)

func Filter(i interface{}) {
	v := reflect.ValueOf(i).Elem()

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		switch f.Kind() {
		case reflect.String:
			if f.CanSet() {
				f.SetString(RemoveCyrillic(f.String()))
			}
		case reflect.Struct:
			Filter(f.Addr().Interface())

		// In case if field is pointer
		case reflect.Ptr:
			switch f.Elem().Kind() {
			case reflect.Struct:
				Filter(f.Interface())
			case reflect.String:
				f.Elem().SetString(RemoveCyrillic(f.Elem().String()))
			}
		}
	}
	return
}

func RemoveCyrillic(s string) (res string) {
	runes := []rune(s)
	var stringBuilder strings.Builder
	for _, r := range runes {
		if !unicode.Is(unicode.Cyrillic, r) {
			stringBuilder.WriteRune(r)
		}
	}

	return stringBuilder.String()
}
