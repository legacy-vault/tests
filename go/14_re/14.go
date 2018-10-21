// 14.go.

package main

//import "bytes"
import "fmt"
import "regexp"

func main() {

	var err error
	var re *regexp.Regexp

	re, err = regexp.Compile("p([a-z]+)ch")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(re.MatchString("peach"))
	fmt.Println(re.FindString("peach punch"))
	fmt.Println(re.FindStringIndex("peach punch"))
	fmt.Println(re.FindStringSubmatch("peach punch"))
	fmt.Println(re.FindStringSubmatchIndex("peach punch"))
}
