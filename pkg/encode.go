package pkg

import (
	"io/fs"
	"os"
	"path/filepath"
	"sync"

	"github.com/levinion/b154/coding"
)

func EncodeFile(filename string, num int) {
	content, err := os.ReadFile(filename)
	checkErr(err)
	result := coding.EncodeToBytes(content)
	for i := 0; i < num-1; i++ {
		result = coding.EncodeToBytes(result)
	}
	os.WriteFile(filename, result, os.ModePerm)
}

func EncodeDir(rootDir string, num int) {
	var wg sync.WaitGroup
	filepath.Walk(rootDir, func(path string, info fs.FileInfo, err error) error {
		checkErr(err)
		if info.IsDir() {
			return nil
		}
		wg.Add(1)
		go func() {
			EncodeFile(path, num)
			wg.Done()
		}()
		return nil
	})
	wg.Wait()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
