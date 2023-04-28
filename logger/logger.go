package logger

import (
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	Info  *log.Logger
	Error *log.Logger
)

type emptyW struct {
}

func (ew *emptyW) Write(p []byte) (n int, err error) {
	return 0, nil
}

func init() {

	//日志输出文件
	fname := "sys.log" + strconv.FormatInt(time.Now().UnixMilli(), 10)
	file, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Faild to open error logger file:", err)
	}
	//自定义日志格式
	Info = log.New(io.MultiWriter(&emptyW{}), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	//Info = log.New(io.MultiWriter(file, os.Stdout), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(file, os.Stdout), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
