package temp

import (
	"fmt"
	"strings"
)

// give a name to a method, if the name with suffix .png return its name
// if its name without .png suffix, then append .png suffix for it and return
func nameChecker(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func nameChecker1(suffix string, name string) string {
	if !strings.HasSuffix(name, suffix) {
		return name + suffix
	}
	return name
}

func main() {

	fmt.Println(nameChecker1(".png", "a.png"))
	fmt.Println(nameChecker1(".png", "a"))

	//closure: inner func with its context as a whole/entire body
	//one can feel closure brings facility, the context(suffix) will be kept for next call,  diff to method , method will destory its context but closure will not
	ft := nameChecker(".png")
	fmt.Println(ft("1.png"))
	fmt.Println(ft("235"))
}
