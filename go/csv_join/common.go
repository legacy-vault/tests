// common.go.

package csv_join

import "os"

const TmpFileNamePostfix = ".tmp"

type IdOffset struct {
	Id     uint
	Offset uint
}

func FileExists(path string) (bool, error) {

	var err error

	_, err = os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	// Unknown Error is considered as a bad File.
	return false, err
}
