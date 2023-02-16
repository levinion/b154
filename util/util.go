package util

import "os"

func IsDir(filename string) bool {
	stat, err := os.Stat(filename)
	checkErr(err)
	return stat.IsDir()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
