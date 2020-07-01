package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetItemsAndUrls(filepath string, secret []byte) ([]string, []string, error) {
	var items []string
	var urls []string

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("%s\n", err)
		return items, urls, err
	} else {
		reader := bufio.NewReader(file)
		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				break
			}
			if string(line[0:3]) != "###" {
				if len(secret) > 0 {
					bytes, err := Decrypt(string(line), secret)
					line = []byte(bytes)
					Check(err)
				}
				line_content := strings.Split(string(line), " - ")
				if len(line_content) == 2 {
					items = append(items, line_content[0])
					urls = append(urls, line_content[1])
				}
			}
		}
		file.Close()
		return items, urls, nil
	}
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func CreateFile(filepath string) {
	file, err := os.OpenFile(filepath, os.O_RDONLY|os.O_CREATE, 0666)
	Check(err)

	defer file.Close()
}

func InitBookmarksFile(filepath string) {
	if !FileExists(filepath) {
		file, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		Check(err)
		_, err = file.WriteString("### Browsy CLI Bookmarks config file ###\n\n")
		Check(err)

		file.Close()
	}
}
