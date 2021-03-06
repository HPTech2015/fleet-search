package dirman

import (
	"io/ioutil"
	"os"
	"regexp"
)

type ListDir struct {
}

/*
	List all contents of the target path.
*/
func (listD *ListDir) LsAll(targetPath string) ([]os.FileInfo, error) {
	var lsError error

	results, err := ioutil.ReadDir(targetPath)
	if err != nil {
		return results, err
	}

	return results, lsError
}

/*
	List all contents of the target path, with name that contains the test string.
*/
func (listD *ListDir) LsContains(targetPath, testStr string) ([]os.FileInfo, error) {
	re := ".*" + testStr + ".*"

	var lsError error
	results := make([]os.FileInfo, 0, 5)

	filePaths, err := ioutil.ReadDir(targetPath)
	if err != nil {
		return results, err
	}

	for _, file := range filePaths {
		matched, _ := regexp.MatchString(re, file.Name())
		if matched {
			results = append(results, file)
		}
	}

	return results, lsError
}

/*
	List all contents of the target path, with name that starts with the test string.
*/
func (listD *ListDir) LsStartsWith(targetPath, testStr string) ([]os.FileInfo, error) {
	re := "^" + testStr
	var lsError error
	results := make([]os.FileInfo, 0, 5)

	filePaths, err := ioutil.ReadDir(targetPath)
	if err != nil {
		return results, err
	}

	for _, file := range filePaths {
		matched, _ := regexp.MatchString(re, file.Name())
		if matched {
			results = append(results, file)
		}
	}

	return results, lsError
}

/*
	List all contents of the target path, with name that ends with the test string.
*/
func (listD *ListDir) LsEndsWith(targetPath, testStr string) ([]os.FileInfo, error) {
	re := testStr + "$"
	var lsError error
	results := make([]os.FileInfo, 0, 5)

	filePaths, err := ioutil.ReadDir(targetPath)
	if err != nil {
		return results, err
	}

	for _, file := range filePaths {
		matched, _ := regexp.MatchString(re, file.Name())
		if matched {
			results = append(results, file)
		}
	}

	return results, lsError
}
