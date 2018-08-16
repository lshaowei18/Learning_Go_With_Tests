package slice

import (
	"reflect"
	"testing"
)

//Example 1 of pointer being passed into the function
func TestAddValuesToSlice(t *testing.T) {
	got := make([]byte, 5)
	AddValuesToSlice(got)
	want := []byte{0, 1, 2, 3, 4}
	assertValues(t, got, want)
}

//Example 2
func TestAddOneToEachElement(t *testing.T) {
	got := make([]byte, 5)
	AddOneToEachElement(got)
	want := []byte{1, 1, 1, 1, 1}
	assertValues(t, got, want)
}

//Showing underlying array changes when slice is edited
func TestUnderlyingArray(t *testing.T) {
	var got [6]byte
	slice := got[0:5]
	AddOneToEachElement(slice)
	want := [6]byte{1, 1, 1, 1, 1, 0}
	assertValues(t, got, want)
}

//Testing if new slice created will affect old TestUnderlyingArray
func TestNewSliceAffectsArray(t *testing.T) {
	var got [6]byte
	slice := got[0:5]
	secondSlice := SubtractOneFromNewSlice(slice)
	AddOneToEachElement(secondSlice)
	want := [6]byte{1, 1, 1, 1, 0, 0}
	assertValues(t, got, want)
}

//Test if you can edit the slice length by passing it into a function
//Can't, have to pass in a pointer.
//contents of slice argument can be modified by a function, but its header cannot.
func TestSubtractOneFromLength(t *testing.T) {
	got := make([]byte, 5)
	SubtractOneFromLength(got)
	want := []byte{0, 0, 0, 0, 0}
	assertValues(t, got, want)
}

func TestPtrSubtractOneFromLength(t *testing.T) {
	got := make([]byte, 5)
	PtrSubtractOneFromLength(&got)
	want := []byte{0, 0, 0, 0}
	assertValues(t, got, want)
}

//Example of passed by values
func TestSubtractOneFromNewSlice(t *testing.T) {
	slice := []byte{1, 2, 3, 4, 5}
	got := SubtractOneFromNewSlice(slice)
	want := []byte{1, 2, 3, 4}

	assertValues(t, got, want)
}

//Increasing slice capacity
func TestExtend(t *testing.T) {
	slice := make([]byte, 2)
	_, got := Extend(slice, 1)
	if got == nil {
		t.Errorf("Expected error due to out of bound, got %v", got)
	}
}

func TestTruncateAtFinalSlash(t *testing.T) {
	got := path("/usr/bin/tso")
	got.TruncateAtFinalSlash()
	want := path("/usr/bin")

	assertValues(t, got, want)
}

func TestToUpper(t *testing.T) {
	t.Run("Change string to upper", func(t *testing.T) {
		got := path("/usr/bin/tso")
		got.ToUpper()
		want := path("/USR/BIN/TSO")
		assertValues(t, got, want)
	})
}

//Helper function
func assertValues(t *testing.T, got, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
