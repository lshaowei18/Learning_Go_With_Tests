package slice

import (
	"bytes"
	"fmt"
	"reflect"
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

func Extend(slice []byte, element byte) (s []byte, err error) {
	defer rescue("Slice out of bound", &err)
	length := len(slice)
	s = slice[0 : length+1]
	s[length] = element

	return
}

func DoubleSliceCapacity(s interface{}) error {
	valPtr := reflect.ValueOf(s)
	if valPtr.Kind() != reflect.Ptr {
		return fmt.Errorf("Interface not a pointer.\n")
	}
	val := valPtr.Elem()
	if val.Kind() != reflect.Slice {
		return fmt.Errorf("Error in DoubleSliceCapacity, interface given not a slice")
	}
	dst := reflect.MakeSlice(val.Type(), val.Len(), val.Len())
	val.Set(reflect.AppendSlice(dst, val))
	return nil
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
