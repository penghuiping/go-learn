package _io

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFile(filename string) {
	//读取文件
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("出错啦,", err)
		return
	}
	buf := make([]byte, 100)
	reader := bufio.NewReader(file)
	for {
		size, error := reader.Read(buf)
		if error == io.EOF {
			break
		} else {
			fmt.Println(string(buf[:size]))
		}
	}
	//有点像java中的finally
	defer file.Close()
}

func readFileAll(filename string) {
	//整个文件内容一起读到一个字符串
	buf, _ := ioutil.ReadFile(filename)
	fmt.Println(string(buf))
}

func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("出错啦,", err)
		return
	}
	//写文件
	writer := bufio.NewWriter(file)
	writer.WriteString("hello world\n")
	writer.WriteString("hello world\n")
	writer.WriteString("hello world\n")
	writer.Flush()
	defer file.Close()
}

func copyFile(srcFilename string, distFilename string) {
	srcfile, err := os.OpenFile(srcFilename, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("出错啦 srcfile,", err)
		return
	}
	defer srcfile.Close()

	distfile, err := os.OpenFile(distFilename, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("出错啦 distfile,", err)
		return
	}
	defer distfile.Close()
	_, err1 := io.Copy(distfile, srcfile)

	if err1 != nil {
		fmt.Println("出错啦", err1)
	}
}

func FileIOExample() {
	basePath := "/Users/penghuiping/Programming/repository/go-learn"
	writeFile(basePath + "/1.txt")
	readFile(basePath + "/1.txt")
	copyFile(basePath+"/1.txt", basePath+"/2.txt")

}
