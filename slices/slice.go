package slice

import (
	"bytes"
	"fmt"
)

type path []byte

func AddValuesToSlice(slice []byte) {
	for i := 0; i < len(slice); i++ {
		slice[i] = byte(i)
	}
}

func AddOneToEachElement(slice []byte) {
	for i := range slice {
		slice[i]++
	}
}

//The slice returned will still point to the underlying array of the given slice
func SubtractOneFromNewSlice(slice []byte) []byte {
	return slice[0 : len(slice)-1]
}

func SubtractOneFromLength(slice []byte) {
	slice = slice[0 : len(slice)-1]
}

func PtrSubtractOneFromLength(slicePtr *[]byte) {
	slice := *slicePtr
	*slicePtr = slice[0 : len(slice)-1]
}

func Extend(slice []byte, element byte) ([]byte, error) {
	var err error
	defer rescue("Slice out of bound", &err)
	length := len(slice)
	slice = slice[0 : length+1]
	slice[length] = element
	return slice, err
}

func (p *path) TruncateAtFinalSlash() {
	i := bytes.LastIndex(*p, []byte("/"))
	if i >= 0 {
		*p = (*p)[0:i]
	}
}

func (p *path) ToUpper() {
	toUpper := func(b rune) rune {
		if 'a' <= b && b <= 'z' {
			b = b + 'A' - 'a'
		}
		return b
	}
	*p = bytes.Map(toUpper, *p)
}

//rescue function to recover from panic
func rescue(errorMessage string, err *error) {
	if r := recover(); r != nil {
		*err = fmt.Errorf("%s", errorMessage)
	}
}
