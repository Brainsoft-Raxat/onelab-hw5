package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveCyrillic(t *testing.T) {
	type testCase struct {
		name string
		text string
		want string
	}
	testCases := []testCase{
		{
			"testing for empty string",
			"",
			"",
		},
		{
			"testing for ordinary string with cyrillic letters",
			"Хочишь Французки булочка.Wanna some french buns?",
			"  .Wanna some french buns?",
		},
		{
			"testing for string that contains only cyrillic letters",
			"ВсемПриветВсемПриветВсемСобрашивмсяРахмет",
			"",
		},
		{
			"testing for string that contains only english letters",
			"Hello World!",
			"Hello World!",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, RemoveCyrillic(tc.text))
		})
	}
}

func TestFilter(t *testing.T) {
	type InnerStruct struct {
		FirstField string
	}

	type TestStruct struct {
		FirstField  int
		SecondField string
		ThirdField  *string
		FourthField struct {
			FirstField  int
			SecondField string
			ThirdField  *string
			FourthField struct {
				FirstField string
			}
		}
		FifthField *InnerStruct
	}

	innerStruct := InnerStruct{"blue - синий"}
	rInnerStruct := InnerStruct{"blue - "}

	str1 := "cool - круто"
	rstr1 := RemoveCyrillic(str1)

	str2 := "red - красный"
	rstr2 := RemoveCyrillic(str2)

	testCase := struct {
		Name       string
		testStruct TestStruct
		want       TestStruct
	}{
		"testing for nested struct",
		TestStruct{
			FirstField:  1,
			SecondField: "Hello - Привет",
			ThirdField:  &str1,
			FourthField: struct {
				FirstField  int
				SecondField string
				ThirdField  *string
				FourthField struct{ FirstField string }
			}{FirstField: 2, SecondField: "blahblah - блабла", ThirdField: &str2, FourthField: struct{ FirstField string }{FirstField: "white - белый"}},
			FifthField: &innerStruct,
		},

		TestStruct{
			FirstField:  1,
			SecondField: "Hello - ",
			ThirdField:  &rstr1,
			FourthField: struct {
				FirstField  int
				SecondField string
				ThirdField  *string
				FourthField struct{ FirstField string }
			}{FirstField: 2, SecondField: "blahblah - ", ThirdField: &rstr2, FourthField: struct{ FirstField string }{FirstField: "white - "}},
			FifthField: &rInnerStruct,
		},
	}

	t.Run(testCase.Name, func(t *testing.T) {
		Filter(&testCase.testStruct)
		assert.Equal(t, testCase.want, testCase.testStruct)
	})

}
