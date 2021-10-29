package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	registry "golang.org/x/sys/windows/registry"
	// registry "github.com/golang/sys/windows/registry"
)

func main() {
	//查询插入的u盘个数
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Services\USBSTOR\Enum`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	n, _, err := k.GetIntegerValue("Count")
	if err != nil {
		log.Fatal(err)
	}
	if n < 1 {
		fmt.Println("没有检测到u盘！")
		os.Exit(-1)
	}

	//查询u盘序列号
	var sn, pvid, vid, pid string
	information, _, err := k.GetStringValue(strconv.Itoa(0))
	strn := strconv.FormatUint(n, 10) //n是uint64类型，先转成string
	nInt, _ := strconv.Atoi(strn)     //再转成int类型
	if n > 1 {
		fmt.Printf("\n检测到多个u盘，按插入顺序输出u盘信息\n\n")
	}
	for i := 0; i < nInt; i++ {
		information, _, err = k.GetStringValue(strconv.Itoa(i))
		if err != nil {
			log.Fatal(err)
		}
		sn = strings.Split(information, "\\")[2]
		pvid = strings.Split(information, "\\")[1]
		vid = strings.Split(pvid, "&")[0]
		vid = strings.Split(vid, "_")[1]
		pid = strings.Split(pvid, "&")[1]
		pid = strings.Split(pid, "_")[1]
		fmt.Println("当前u盘sn码:", sn)
		fmt.Println("当前u盘vid: ", vid)
		fmt.Println("当前u盘pid: ", pid)
		fmt.Println("")
	}
	time.Sleep(5 * time.Minute)
}
