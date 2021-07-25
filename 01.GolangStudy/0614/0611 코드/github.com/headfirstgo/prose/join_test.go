package prose

import (
	"testing"
)

type testData struct {
	list []string
	want string
}

func TestJointWithCommas(t *testing.T) {
	tests := []testData{
		testData{list: []string{}, want: ""},
		testData{list: []string{"apple"}, want: "apple"},
		testData{list: []string{"apple", "orange"}, want: "apple and orange"},
		testData{list: []string{"apple", "orange", "pear"}, want: "apple, orange, and pear"},
	}
	for _, test := range tests {
		got := JoinWithCommas(test.list)
		if got != test.want {
			t.Errorf("JoinWithCommas(%#v)= \"%s\", want \"%s\"", test.list, got, test.want)
		}
	}
}

// func errorString(list []string, got string, want string) string {
// 	return fmt.Sprintf("JoinWithCommas(%#v)= \"%s\", want \"%s\"", list, got, want)
// }

// func TestOneElements1(t *testing.T) {
// 	list := []string{"apple"}
// 	want := "apple"
// 	got := JoinWithCommas(list)
// 	if got != want {
// 		t.Error(errorString(list, got, want))
// 	}
// }

// func TestTwoElements1(t *testing.T) {
// 	list := []string{"apple", "orange"}
// 	want := "apple and orange"  // 원하는 반환값
// 	got := JoinWithCommas(list) //실제 반환값
// 	if got != want {
// 		t.Error(errorString(list, got, want))
// 		//t.Errorf("JoinWithCommas(%#v)= \"%s\", want \"%s\"", list, got, want)
// 	}
// }

// func TestThreeElements1(t *testing.T) {
// 	list := []string{"apple", "orange", "pear"}
// 	want := "apple, orange, and pear" // 원하는 반환값
// 	got := JoinWithCommas(list)       //실제 반환값
// 	if got != want {
// 		t.Error(errorString(list, got, want))
// 		//t.Errorf("JoinWithCommas(%#v)= \"%s\", want \"%s\"", list, got, want)
// 	}
// }
