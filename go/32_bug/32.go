// 32.go.

package main

type Destination struct {
	Creator func(int) error
	Size    int
}
type UserConfig struct {
	Name         string
	Destinations []*Destination
}
type User struct {
	Config *UserConfig
}
type Record struct {
	Person *User
	ID     int
}
type Records [2]*Record

func CreatorOfSomething(x int) error {
	return nil
}

func main() {

	var aCreator func(int) error
	var aDestination *Destination
	var destinations []*Destination
	var rec *Record
	var recs Records
	var userCfg *UserConfig
	var user *User

	aCreator = CreatorOfSomething
	aDestination = &Destination{
		Creator: aCreator,
		Size:    10,
	}
	destinations = []*Destination{aDestination}
	userCfg = &UserConfig{
		Name:         "John",
		Destinations: destinations,
	}
	user = &User{
		Config: userCfg,
	}
	rec = &Record{
		Person: user,
		ID:     123,
	}
	recs[0] = rec
}
