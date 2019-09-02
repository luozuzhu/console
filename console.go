package console

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)
type color struct {
	error map[string]string
	info map[string]string
	warn map[string]string
	log map[string]string
}
func getColor() color{
	var c color = color{
		map[string]string{
			"start":"\033[31m",
			"end":"\033[0m",
		},
		map[string]string{
			"start":"\033[35m",
			"end":"\033[0m",
		},
		map[string]string{
			"start":"\033[33m",
			"end":"\033[0m",
		},
		map[string]string{
			"start":"\033[34m",
			"end":"\033[0m",
		},
	}
	return c
}
func Log(s ...interface{}) {
	_,file,line,_ := runtime.Caller(1)
	color := getColor()
	log(s,color.log,file,line)
}
func Info(s ...interface{}){
	_,file,line,_ := runtime.Caller(1)
	color := getColor()
	log(s,color.info,file,line)
}
func Warn(s ...interface{}){
	_,file,line,_ := runtime.Caller(1)
	color := getColor()
	log(s,color.warn,file,line)
}
func Error(s ...interface{}){
	_,file,line,_ := runtime.Caller(1)
	color := getColor()
	log(s,color.error,file,line)
	tmpFile,_ := os.OpenFile("./error.log",os.O_CREATE,0666)
	_ = tmpFile.Close()
	appendFile,_ := os.OpenFile("./error.log",os.O_APPEND,0666)
	errorStr := getNow() + "\t"
	errorStr += fmt.Sprint(s,"")
	errorStr += file  + ":" + strconv.Itoa(line)
	_,_ = appendFile.WriteString(errorStr)
	_ = appendFile.Close()
}
func log(source []interface{},color map[string]string,file string,line int){
	now := getNow()
	fmt.Print(now,"\t")
	fmt.Print(color["start"])
	for _,val := range source {
		fmt.Print(val," ")
	}
	fmt.Print("\t")
	fmt.Print(color["end"])
	fmt.Printf("%s:%d\n",file,line)
	fmt.Println("")
}
//2e460c06415411e9bc7e000c2974f54a
func Mark(s ...interface{}) {
	mark := s[len(s) - 1]
	_,file,line,_ := runtime.Caller(1)
	color := getColor()
	s2 := strings.Repeat("-",100)
	s3 := strings.Repeat("-",102)
	fmt.Print(s2 + mark.(string) + "\tstart" + s2)
	fmt.Println("")
	log(s[0:len(s) - 1],color.log,file,line)
	fmt.Print(s2 + mark.(string) + "\tend" + s3)
	fmt.Println("")
}

func getNow()(t string){
	now := time.Now()
	y,m,d := now.Date()
	h,M,s := now.Clock()
	t = time.Date(y, m, d, h, M, s, 0, time.UTC).Local().String()
	return
}