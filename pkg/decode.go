package pkg

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sync"

	"github.com/levinion/b154/coding"
)

func DecodeFile(filename string, num int, show bool) {
	content, err := os.ReadFile(filename)
	checkErr(err)
	result := coding.DecodeToBytes(content)
	for i := 0; i < num-1; i++ {
		result = coding.DecodeToBytes(result)
	}
	if show {
		fmt.Println(string(result))
	}
	os.WriteFile(filename, result, os.ModePerm)
}

func DecodeDir(rootDir string, num int, show bool) {
	var wg sync.WaitGroup
	filepath.Walk(rootDir, func(path string, info fs.FileInfo, err error) error {
		checkErr(err)
		if info.IsDir() {
			return nil
		}
		wg.Add(1)
		go func() {
			DecodeFile(path, num, show)
			wg.Done()
		}()
		return nil
	})
	wg.Wait()
}
