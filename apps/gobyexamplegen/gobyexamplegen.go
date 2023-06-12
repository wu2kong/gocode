package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

/*
操作流程
1. 读取目录
2. 根据目录读取文件 go+sh 文件
3. 生成doc md文档
4. 写入到指定位置

注意事项
1、判断是否已经存在
*/

func main() {
	exampleDir := "/Users/lihao/github/gobyexample/examples/"
	outDocDir := "/Users/lihao/gitme/learngo/docs/gobyexample/"

	// 构建文件名相关的映射数据
	f, _ := os.Open("./filenamemap.txt")
	defer f.Close()

	reader := bufio.NewReader(f)
	type mappingItem struct {
		Title string
		Index int
	}
	fileNameMapping := make(map[string]mappingItem)
	var idx int
	for {
		idx++
		row, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		rowContent := string(row)
		segs := strings.Split(rowContent, "->")

		var title string
		if len(segs) > 1 {
			title = segs[1]
		} else {
			title = segs[0]
		}

		rSpace := strings.ReplaceAll(segs[0], " ", "-")
		rOther := strings.ReplaceAll(rSpace, "/", "-")
		key := strings.ToLower(rOther)

		fileNameMapping[key] = mappingItem{Title: title, Index: idx}
	}

	// 循环读取目录列表
	outdirInfo, _ := ioutil.ReadDir(outDocDir)
	existFileDict := make(map[string]bool)
	for _, info := range outdirInfo {
		dirName := info.Name()
		// print(dirName)
		// fmt.Println()
		existFileDict[dirName] = true
	}

	// 循环读取目录列表
	exampleDirInfo, _ := ioutil.ReadDir(exampleDir)
	for _, info := range exampleDirInfo {
		dirName := info.Name()
		inGoFilename := fmt.Sprintf("%s%s/%s.go", exampleDir, dirName, dirName)
		inShellFilename := fmt.Sprintf("%s%s/%s.sh", exampleDir, dirName, dirName)

		// 查找已经存在的
		mappingItem, isMappingExist := fileNameMapping[dirName]
		var title string
		var idx int
		if isMappingExist {
			title = mappingItem.Title
			idx = mappingItem.Index
		}

		// 输出信息
		outMarkdownFilename := fmt.Sprintf("%s%0.2d-%s.md", outDocDir, idx, dirName)
		isExist := matchExist(dirName+".md", existFileDict)
		if isExist {
			println("ignored: ", outMarkdownFilename)
			continue
		}

		goContent, _ := os.ReadFile(inGoFilename)
		shellContent, _ := os.ReadFile(inShellFilename)

		outMarkdownContent := fmt.Sprintf("---\n\n---\n\n# %s\n\n```go\n%s```\n\n```bash\n%s```\n\n", title,
			goContent, shellContent)

		err := os.WriteFile(outMarkdownFilename, []byte(outMarkdownContent), 0644)
		if err != nil {
			print(string(err.Error()))
			break
		}
		// println(inGoFilename)
		// println(inShellFilename)
		// println(outMarkdownFilename)
		println("success: ", outMarkdownFilename)
		// break
	}

	println("all done.")
}

func matchExist(filename string, mapping map[string]bool) bool {
	isExist := false
	for k := range mapping {
		// 如果存在，就忽略；否则继续
		if strings.HasSuffix(k, filename) {
			isExist = true
		}
	}

	return isExist
}
