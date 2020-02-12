package main

import (
	"bufio"
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

var dirname = "./"

func main()  {
	dirname := flag.String("dn", "./", "需要检查重复文件的路径，如 A:\\CloudMusic，" +
		"路径最后面建议不要带 \\ 或 /")
	deleteFile := flag.String("de", "", "需要删除重复文件的清单")
	flag.Parse()

	if *deleteFile != "" {
		batchDeleteFile(*deleteFile)
		os.Exit(0)
	}

	//当前时间戳
	timeUnix := time.Now().Unix()

	//输出文件路径
	outPutPath := *dirname+"/"+strconv.FormatInt(timeUnix, 10)+"_delete_list.txt"

	files, _ := ioutil.ReadDir(*dirname)

	var musicMd5Temp map[string]string /*创建集合 */
	musicMd5Temp = make(map[string]string)
	var fileNum = 0
	var repeatNum = 0

	for _,file := range files {
		if file.IsDir() {
			continue
		} else {
			fileNum += 1

			fileName := *dirname+"/"+file.Name()
			md5Value, err := Md5SumFile(fileName)

			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			result := fmt.Sprintf("%x", md5Value)//文件MD5值计算结果

			if getMapKey(result, musicMd5Temp) != "" {
				repeatPath := getMapKey(result, musicMd5Temp)
				fmt.Printf("  文件名：%s\n重复路径：%s\n\n", fileName, repeatPath)

				writeFile(outPutPath, repeatPath+"\n", "APPEND")

				repeatNum += 1
			}

			//在地图中插入
			musicMd5Temp[fileName] = result
			fmt.Println(result+"\t"+fileName)
		}
	}

	if repeatNum != 0{
		fmt.Printf("扫描完成，共扫描 %d 文件，有 %d 文件重复，耗时：%ds\n重复文件清单路径(没有重复文件不会建立该文件)：" +
			"%s\n请使用 -de 参数，指定重复名单路径进行删除重复文件",
			fileNum, repeatNum, time.Now().Unix()-timeUnix, outPutPath)
	}else{
		fmt.Printf("扫描完成，共扫描 %d 文件，有 %d 文件重复，耗时：%ds",
			fileNum, repeatNum, time.Now().Unix()-timeUnix)
	}

}

func batchDeleteFile(fileList string) {
	fi, err := os.Open(fileList)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}

		if string(a) != "" {
			err := os.Remove(string(a))
			if err != nil {
				fmt.Printf("[删除失败]%s\t原因：%s\n", string(a), err)
			}else{
				fmt.Printf("[删除成功]%s\n", string(a))
			}

		}
	}
}

/**
 * 写入文件
 */
func writeFile(filename, writeString, writeType string) {
	var f *os.File
	var err1 error

	//如果文件存在
	if checkFileIsExist(filename) {
		if writeType == "APPEND" {
			//打开文件，以追加的方式写入
			f, err1 = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0666)
		} else {
			//打开文件，以覆盖的方式写入
			f, err1 = os.OpenFile(filename, os.O_WRONLY, 0666)
		}
	} else {
		//文件不存在，创建新的文件
		f, err1 = os.Create(filename)
	}

	check(err1)
	//n, err1 := io.WriteString(f, writeString) //写入文件(字符串)
	_, err1 = io.WriteString(f, writeString) //写入文件(字符串)
	check(err1)
	//fmt.Printf("写入 %d 个字节n", n)
}


func getMapKey(md5 string, maps map[string]string) string {
	for fileName := range maps {
		//如果该文件的MD5值 == 待查找的 MD5
		if maps[fileName] == md5 {
			//该文件重复
			return fileName
		}
	}

	//输出空，该文件不重复
	return ""
}

//计算文件 MD5 值
func Md5SumFile(file string) (value [md5.Size]byte, err error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}

	value = md5.Sum(data)
	return
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
