package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"xxx.com/loghandle/excel"

	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	filePath := "access.log"
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		glog.Error(err)
		return
	}
	titles := map[string]string{"A1": "Ip", "B1": "Time", "C1": "Path", "D1": "Status"}
	excelFile := excel.InitExcel("Sheet1", titles, true)
	buf := bufio.NewReader(f)
	j := 2
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			glog.Error(err)
			return
		}

		lineStr := strings.TrimSpace(string(line))

		re := `^([\d]{1,3}\.[\d]{1,3}\.[\d]{1,3}\.[\d]{1,3}) - - \[(.*)\] "([^\s]+) ([^\s]+) ([^\s]+?)" ([\d]{3}) ([\d]{1,9}) "([^"]*?)" "([^"]*?)"`
		reg := regexp.MustCompile(re)

		parseInfo := reg.FindStringSubmatch(lineStr)
		fmt.Println(parseInfo)

		// 匹配不到正常的格式，那么这条访问记录有问题
		// 跳过
		if len(parseInfo) == 0 {
			fmt.Println("解析异常")
			continue
		}

		t, _ := time.Parse("02/Jan/2006:15:04:05 -0700", parseInfo[2])
		status, _ := strconv.Atoi(parseInfo[6])
		//size, _ := strconv.Atoi(parseInfo[7])
		StrJ := strconv.Itoa(j)
		contents := make(map[string]string)
		contents["A"+StrJ] = parseInfo[1]
		contents["B"+StrJ] = t.Format("2006-01-02 15:04:05")
		contents["C"+StrJ] = parseInfo[4]
		contents["D"+StrJ] = strconv.Itoa(status)

		excel.ExportExcel(excelFile, "Sheet1", contents)
		j++

	}

	excel.SaveExcel(excelFile)
}
