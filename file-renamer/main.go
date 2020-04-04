package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	dir := flag.String("dir", "./sample", "directory to search")
	flag.Parse()
	count := 1
	err := filepath.Walk(*dir, func(path string, info os.FileInfo, err error) error {
		if isMatch(info.Name(), "n_008.txt") {
			newPath := rename(path)
			p := fmt.Sprintf("%s/%s-%d.txt", newPath, "whatup", count)
			c := exec.Command("mv", path, p)
			c.Run()
			count++
		}
		return err
	})

	if err != nil {
		panic(err)
	}
	fmt.Println("done")
}

func isMatch(name string, matchTo string) bool {
	re := regexp.MustCompile(`^[a-z]{1}_[0-9]{3}.txt$`)
	return re.MatchString(strings.ToLower(name))
}

func rename(path string) string {
	arr := strings.SplitN(path, "/", -1)
	newArr := arr[:len(arr)-1]
	return strings.Join(newArr, "/")
}
