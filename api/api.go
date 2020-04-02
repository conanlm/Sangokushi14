package api

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

//Rename 重命名
func Rename() {
	files, _ := filepath.Glob("./gx/*")
	for _, value := range files {
		file := strings.Split(value[3:len(value)], ".")
		// fmt.Println(file)
		i, _ := strconv.Atoi(file[0])
		fmt.Println(i)
		// // 重命名文件
		err1 := os.Rename(value, "./gx/"+strconv.Itoa(i-1)+".png")
		if err1 != nil {
			panic(err1)
		} else {
			// println(`文件重命名成功`)
		}
	}
}

func T统一文件名位数() {
	files, _ := filepath.Glob("./gx/*")
	for _, value := range files {
		file := strings.Split(value[3:len(value)], ".")
		// fmt.Println(file)
		// i, _ := strconv.Atoi(file[0])
		// fmt.Println(len(file[0]))
		if len(file[0]) == 1 {
			err1 := os.Rename(value, "./gx/00"+file[0]+".png")
			if err1 != nil {
				panic(err1)
			}
		}
		if len(file[0]) == 2 {
			err1 := os.Rename(value, "./gx/0"+file[0]+".png")
			if err1 != nil {
				panic(err1)
			}
		}
		// // 重命名文件
		// err1 := os.Rename(value, "./gx/"+strconv.Itoa(i-1)+".png")
		// if err1 != nil {
		// 	panic(err1)
		// }
	}
}
