package main

import (
	// "Sangokushi14/api"

	"Sangokushi14/api"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/panjf2000/ants/v2"
)

type coordinate struct {
	x int
	y int
}

func main() {
	// 三国()
	// is个性2()
	// 截图()
	// getbitmap()
	// Bitmap()
	// time.Sleep(time.Second)
	// x, y := robotgo.GetMousePos()
	// fmt.Println("pos:", x, y)

	// 手动个性截图()
	// 手动战法截图()

	// api.GetdataClient()
	// api.ExampleClient()
	// 个性写入(逐个获取对比())
	个性写入(逐个战法获取对比())

}

//MouseClick 鼠标
func MouseClick() {
	robotgo.ScrollMouse(10, "up")
	robotgo.MouseClick("left", true)
	robotgo.MoveMouseSmooth(100, 200, 1.0, 100.0)

}

// Key 键盘
func Key() {
	robotgo.TypeStr("Hello World")
	robotgo.TypeStr("测试")
	// robotgo.TypeString("测试")

	robotgo.TypeStr("山达尔星新星军团, galaxy. こんにちは世界.")
	robotgo.Sleep(1)

	// ustr := uint32(robotgo.CharCodeAt("测试", 0))
	// robotgo.UnicodeType(ustr)

	robotgo.KeyTap("enter")
	// robotgo.TypeString("en")
	robotgo.KeyTap("i", "alt", "command")

	arr := []string{"alt", "command"}
	robotgo.KeyTap("i", arr)

	robotgo.WriteAll("测试")
	text, err := robotgo.ReadAll()
	if err == nil {
		fmt.Println(text)
	}

}

func Bitmap() {
	bitmap := robotgo.CaptureScreen(450, 330, 40, 25)
	// use `defer robotgo.FreeBitmap(bit)` to free the bitmap
	defer robotgo.FreeBitmap(bitmap)
	fmt.Println("...", bitmap)

	fx, fy := robotgo.FindBitmap(bitmap)
	robotgo.MoveMouse(fx, fy)
	fmt.Println("FindBitmap------", fx, fy)

	robotgo.SaveBitmap(bitmap, "test.png")
}

func getbitmap() {
	bitmap := robotgo.OpenBitmap("test.png")
	defer robotgo.FreeBitmap(bitmap)
	fmt.Println("...", bitmap)

	fx, fy := robotgo.FindBitmap(bitmap)
	robotgo.MoveMouse(fx, fy)
	fmt.Println("FindBitmap------", fx, fy)

	robotgo.SaveBitmap(bitmap, "test.png")

}

func 三国() {
	var tab1 coordinate
	tab1.x = 610
	tab1.y = 504
	robotgo.MoveMouse(tab1.x, tab1.y)
	robotgo.MouseClick("right")
	robotgo.MoveMouse(661, 443)
	robotgo.MouseClick("lift")

}

func is个性1() {
	// bit := robotgo.CaptureScreen(515, 277, 332, 90)
	// bit := robotgo.CaptureScreen(515, 277, 225, 32)
	time.Sleep(time.Second * 2)
	bitmap := robotgo.OpenBitmap("gx/3.png")
	defer robotgo.FreeBitmap(bitmap)
	fmt.Println("...", bitmap)

	// fx, fy := robotgo.FindBitmap(bitmap, bit)
	fx, fy := robotgo.FindBitmap(bitmap)
	robotgo.MoveMouse(fx, fy)
	if fx != -1 {
		fmt.Println("FindBitmap------", fx, fy)
	}

	// robotgo.SaveBitmap(bitmap, "test.png")
}

func is个性2() []robotgo.Bitmap {
	time.Sleep(time.Second * 2)
	bitmap := []robotgo.Bitmap{}
	files, _ := filepath.Glob("./gx/*")
	for _, value := range files {
		Bmap := robotgo.OpenBitmap(value)
		bitmap = append(bitmap, robotgo.ToBitmap(Bmap))
	}
	// fmt.Println(bitmap[3])
	// fx, fy := robotgo.FindBitmap(robotgo.ToCBitmap(bitmap[3]))
	// fmt.Println("FindBitmap------", fx, fy)
	fmt.Println(len(bitmap))
	return bitmap
}
func is战法() []robotgo.Bitmap {
	bitmap := []robotgo.Bitmap{}
	files, _ := filepath.Glob("./zf/*")
	for _, value := range files {
		Bmap := robotgo.OpenBitmap(value)
		bitmap = append(bitmap, robotgo.ToBitmap(Bmap))
	}
	fmt.Println(len(bitmap))
	return bitmap
}

func 获取武将个性(bitmaps *[]robotgo.Bitmap) map[int]int {
	个性 := map[int]int{}
	bT := time.Now()
	fx, fy := robotgo.FindBitmap(robotgo.ToCBitmap((*bitmaps)[0]))

	fmt.Println("FindBitmap------", fx, fy)
	var ii int
	if fx == 399 && fy == 278 {
		ii = 0
	}
	if fx == 510 && fy == 278 {
		ii = 1
	}
	if fx == 621 && fy == 278 {
		ii = 2
	}
	if fx == 399 && fy == 305 {
		ii = 3
	}
	if fx == 510 && fy == 305 {
		ii = 4
	}
	if fx == 621 && fy == 305 {
		ii = 5
	}
	if fx == 399 && fy == 332 {
		ii = 6
	}
	if fx == 510 && fy == 332 {
		ii = 7
	}
	if fx == 621 && fy == 332 {
		ii = 8
	}
	for i := 1; i <= 171; i++ {
		fx, fy := robotgo.FindBitmap(robotgo.ToCBitmap((*bitmaps)[i]))
		if fx != -1 {
			fmt.Println("FindBitmap------", fx, fy)
			if fx == 窗口.x+321 && fy == 窗口.y+200 {
				个性[1] = i
				continue
			}
			if fx == 窗口.x+432 && fy == 窗口.y+200 {
				个性[2] = i
				continue
			}
			if fx == 窗口.x+543 && fy == 窗口.y+200 {
				个性[3] = i
				continue
			}
			if fx == 窗口.x+321 && fy == 窗口.y+227 {
				个性[4] = i
				continue
			}
			if fx == 窗口.x+432 && fy == 窗口.y+227 {
				个性[5] = i
				continue
			}
			if fx == 窗口.x+543 && fy == 窗口.y+227 {
				个性[6] = i
				continue
			}
			if fx == 窗口.x+321 && fy == 窗口.y+254 {
				个性[7] = i
				continue
			}
			if fx == 窗口.x+432 && fy == 窗口.y+254 {
				个性[8] = i
				continue
			}
			if fx == 窗口.x+543 && fy == 窗口.y+254 {
				个性[9] = i
				continue
			}
		}
		if len(个性) == ii {
			break
		}

	}
	eT := time.Since(bT)
	fmt.Println("Run time: ", eT)
	return 个性
}

type Request struct {
	intt   int
	Result chan int
}

func 并发获取武将个性1(bitmaps *[]robotgo.Bitmap) map[int]int {

	个性 := map[int]int{}
	bT := time.Now()
	wg := sync.WaitGroup{}
	p, _ := ants.NewPoolWithFunc(50, func(ii interface{}) {
		i := ii.(int)
		fx, fy := robotgo.FindBitmap(robotgo.ToCBitmap((*bitmaps)[i]))
		if fx != -1 {
			fmt.Println("FindBitmap------", fx, fy)
			switch {
			case fx == 窗口.x+321 && fy == 窗口.y+200:
				个性[1] = i
			case fx == 窗口.x+432 && fy == 窗口.y+200:
				个性[2] = i
			case fx == 窗口.x+543 && fy == 窗口.y+200:
				个性[3] = i
			case fx == 窗口.x+321 && fy == 窗口.y+227:
				个性[4] = i
			case fx == 窗口.x+432 && fy == 窗口.y+227:
				个性[5] = i
			case fx == 窗口.x+543 && fy == 窗口.y+227:
				个性[6] = i
			case fx == 窗口.x+321 && fy == 窗口.y+254:
				个性[7] = i
			case fx == 窗口.x+432 && fy == 窗口.y+254:
				个性[8] = i
			case fx == 窗口.x+543 && fy == 窗口.y+254:
				个性[9] = i
			}
		}
		wg.Done()
	})
	defer p.Release()
	for i := 1; i <= 171; i++ {
		wg.Add(1)
		_ = p.Invoke(i)
	}
	wg.Wait()
	eT := time.Since(bT)
	fmt.Println("Run time: ", eT)
	return 个性
}

func 并发获取武将个性2(bitmaps *[]robotgo.Bitmap) map[int]int {
	个性 := map[int]int{}
	bT := time.Now()
	wg := sync.WaitGroup{}
	fx, fy := robotgo.FindBitmap(robotgo.ToCBitmap((*bitmaps)[0]))
	var iii int
	switch {
	case fx == 窗口.x+321 && fy == 窗口.y+200:
		iii = 0
	case fx == 窗口.x+432 && fy == 窗口.y+200:
		iii = 1
	case fx == 窗口.x+543 && fy == 窗口.y+200:
		iii = 2
	case fx == 窗口.x+321 && fy == 窗口.y+227:
		iii = 3
	case fx == 窗口.x+432 && fy == 窗口.y+227:
		iii = 4
	case fx == 窗口.x+543 && fy == 窗口.y+227:
		iii = 5
	case fx == 窗口.x+321 && fy == 窗口.y+254:
		iii = 6
	case fx == 窗口.x+432 && fy == 窗口.y+254:
		iii = 7
	case fx == 窗口.x+543 && fy == 窗口.y+254:
		iii = 8
	}
	p, _ := ants.NewPoolWithFunc(50, func(ii interface{}) {
		est := ii.(*Request)
		i := est.intt
		fx, fy := robotgo.FindBitmap(robotgo.ToCBitmap((*bitmaps)[i]))
		if fx != -1 {
			fmt.Println("FindBitmap------", fx, fy)
			switch {
			case fx == 窗口.x+321 && fy == 窗口.y+200:
				个性[1] = i
			case fx == 窗口.x+432 && fy == 窗口.y+200:
				个性[2] = i
			case fx == 窗口.x+543 && fy == 窗口.y+200:
				个性[3] = i
			case fx == 窗口.x+321 && fy == 窗口.y+227:
				个性[4] = i
			case fx == 窗口.x+432 && fy == 窗口.y+227:
				个性[5] = i
			case fx == 窗口.x+543 && fy == 窗口.y+227:
				个性[6] = i
			case fx == 窗口.x+321 && fy == 窗口.y+254:
				个性[7] = i
			case fx == 窗口.x+432 && fy == 窗口.y+254:
				个性[8] = i
			case fx == 窗口.x+543 && fy == 窗口.y+254:
				个性[9] = i
			}
		}
		est.Result <- len(个性)
		wg.Done()
	})
	defer p.Release()
	request := &Request{
		intt:   0,
		Result: make(chan int),
	}
	for i := 1; i <= 171; i++ {
		wg.Add(1)
		request.intt = i
		_ = p.Invoke(request)
		if iii == <-request.Result {
			fmt.Println(i)
			break
		}
	}
	wg.Wait()
	eT := time.Since(bT)
	fmt.Println("Run time: ", eT)
	return 个性
}

func 并发获取武将战法1(bitmaps *[]robotgo.Bitmap) map[int]int {
	个性 := map[int]int{}
	bT := time.Now()
	wg := sync.WaitGroup{}
	p, _ := ants.NewPoolWithFunc(50, func(ii interface{}) {
		i := ii.(int)
		fx, fy := robotgo.FindBitmap(robotgo.ToCBitmap((*bitmaps)[i]))
		if fx != -1 {
			fmt.Println("FindBitmap------", fx, fy)
			switch {
			case fx == 窗口.x+293 && fy == 窗口.y+200:
				个性[1] = i
			case fx == 窗口.x+390 && fy == 窗口.y+200:
				个性[2] = i
			case fx == 窗口.x+487 && fy == 窗口.y+200:
				个性[3] = i
			case fx == 窗口.x+584 && fy == 窗口.y+200:
				个性[4] = i
			case fx == 窗口.x+293 && fy == 窗口.y+227:
				个性[5] = i
			case fx == 窗口.x+390 && fy == 窗口.y+227:
				个性[6] = i
			case fx == 窗口.x+487 && fy == 窗口.y+227:
				个性[7] = i
			case fx == 窗口.x+584 && fy == 窗口.y+227:
				个性[8] = i
			case fx == 窗口.x+293 && fy == 窗口.y+254:
				个性[9] = i
			case fx == 窗口.x+390 && fy == 窗口.y+254:
				个性[10] = i
			}
		}
		wg.Done()
	})
	defer p.Release()
	for i := 1; i <= 110; i++ {
		wg.Add(1)
		_ = p.Invoke(i)
	}
	wg.Wait()
	eT := time.Since(bT)
	fmt.Println("Run time: ", eT)
	return 个性
}

func 截图() {
	time.Sleep(time.Second * 2)
	bitmap := robotgo.CaptureScreen(633, 285, 40, 20)
	// use `defer robotgo.FreeBitmap(bit)` to free the bitmap
	defer robotgo.FreeBitmap(bitmap)
	fmt.Println("...", bitmap)
	robotgo.SaveBitmap(bitmap, "gx/t.png")
}

func 手动个性截图() {
	for i := 161; i <= 170; i++ {
		robotgo.MoveMouse(600, 255)
		time.Sleep(time.Second / 2)
		robotgo.MouseClick("lift")
		time.Sleep(time.Second / 2)
		robotgo.KeyTap("down")
		time.Sleep(time.Second / 2)
		robotgo.KeyTap("f1")
		time.Sleep(time.Second)
		robotgo.MoveMouse(650, 292)
		time.Sleep(time.Second / 2)
		robotgo.MouseClick("lift")
		time.Sleep(time.Second)
		robotgo.KeyTap("s", "control")
		time.Sleep(time.Second * 1)
		robotgo.TypeStr(strconv.Itoa(i))
		time.Sleep(time.Second)
		// robotgo.MoveMouse(675, 331)
		robotgo.MoveMouse(457, 447)
		time.Sleep(time.Second / 2)
		robotgo.MouseClick("lift")
		time.Sleep(time.Second * 1)

	}

}
func 手动战法截图() {
	for i := 1; i <= 49; i++ {
		robotgo.MoveMouse(750, 270)
		time.Sleep(time.Second / 1000)
		robotgo.MouseClick("lift")
		time.Sleep(time.Second / 500)
		robotgo.KeyTap("down")
		time.Sleep(time.Second / 3)
		robotgo.KeyTap("f1")
		time.Sleep(time.Second)
		robotgo.MoveMouse(670, 305)
		time.Sleep(time.Second / 10)
		robotgo.MouseClick("lift")
		time.Sleep(time.Second / 20)
		robotgo.KeyTap("s", "control")
		time.Sleep(time.Second / 2)
		robotgo.TypeStr(fmt.Sprintf("%03d", i))
		// robotgo.MoveMouse(675, 331)
		robotgo.MoveMouse(457, 447)
		time.Sleep(time.Second / 20)
		robotgo.MouseClick("lift")
		time.Sleep(time.Second / 10)

	}

}

func 逐个获取对比() []map[int]int {
	bitmaps := is个性2()
	个性s := []map[int]int{}
	for i := 601; i <= 1000; i++ {
		robotgo.MoveMouse(窗口.x+470, 窗口.y+420)
		// time.Sleep(time.Second / 2)
		robotgo.MouseClick("right")
		time.Sleep(time.Second / 10)
		robotgo.MoveMouse(窗口.x+490, 窗口.y+360)
		robotgo.MouseClick("lift")
		time.Sleep(time.Second * 1)
		bT := time.Now()
		个性 := 并发获取武将个性1(&bitmaps)
		// 个性 := 并发获取武将个性2(&bitmaps)
		// 个性 := 获取武将个性(&bitmaps)
		eT := time.Since(bT)
		fmt.Println("Run1 time: ", eT)
		个性s = append(个性s, 个性)
		robotgo.KeyTap("esc")
		time.Sleep(time.Second / 10)
		robotgo.KeyTap("down")
		// robotgo.KeyTap("down")
		// time.Sleep(time.Second / 2)
		// robotgo.KeyTap("f1")
		// time.Sleep(time.Second)
		// robotgo.MoveMouse(650, 292)
		// time.Sleep(time.Second / 2)
		// robotgo.MouseClick("lift")
		// time.Sleep(time.Second)
		// robotgo.KeyTap("s", "control")
		// time.Sleep(time.Second * 1)
		// robotgo.TypeStr(strconv.Itoa(i))
		// time.Sleep(time.Second)
		// // robotgo.MoveMouse(675, 331)
		// robotgo.MoveMouse(457, 447)
		// time.Sleep(time.Second / 2)
		// robotgo.MouseClick("lift")
		time.Sleep(time.Second / 10)
	}
	fmt.Println(个性s)
	return 个性s

}
func 逐个战法获取对比() []map[int]int {
	bitmaps := is战法()
	个性s := []map[int]int{}
	for i := 901; i <= 1000; i++ {
		robotgo.MoveMouse(窗口.x+470, 窗口.y+420)
		// time.Sleep(time.Second / 2)
		robotgo.MouseClick("right")
		time.Sleep(time.Second / 10)
		robotgo.MoveMouse(窗口.x+530, 窗口.y+370)
		robotgo.MouseClick("lift")
		time.Sleep(time.Second * 1)
		bT := time.Now()
		个性 := 并发获取武将战法1(&bitmaps)
		eT := time.Since(bT)
		fmt.Println("Run1 time: ", eT)
		个性s = append(个性s, 个性)
		robotgo.KeyTap("esc")
		time.Sleep(time.Second / 10)
		robotgo.KeyTap("down")
		time.Sleep(time.Second / 10)
	}
	fmt.Println(个性s)
	return 个性s
}

func 个性写入(map1 []map[int]int) {
	个性s := []string{}
	client := api.CreatClient()
	str := api.GetdataClient1(client)
	for _, values := range map1 {
		个性 := ""
		for i := 1; i <= len(values); i++ {
			个性 = 个性 + " " + str[values[i]]
		}
		个性s = append(个性s, 个性)
	}
	f, err := os.Create("2.txt")
	if err != nil {
		fmt.Println(err)
		f.Close()
	}

	for _, v := range 个性s {
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")
}

type 窗口坐标 struct {
	x int
	y int
}

var 窗口 窗口坐标

func init() {

	// fpid2, _ := robotgo.FindIds("San14EditorSC")
	fpid2, _ := robotgo.FindIds("San14EditorSC")
	// if err == nil {
	// fmt.Println("pids...", fpid2)
	// }

	// fpid, _ := robotgo.FindNames()
	// fmt.Println("pids...", fpid[0][1])
	for _, v := range fpid2 {
		fpid1, _ := robotgo.FindName(v)
		fmt.Println("pids...", fpid1)
	}
	// mdata := robotgo.GetActive()
	// hwnd := robotgo.FindWindow("San14EditorSC")
	// robotgo.SetActive(hwnd)
	robotgo.ActivePID(fpid2[0])
	// robotgo.SetHandlePid(fpid2[0])
	x, y, w, h := robotgo.GetBounds(fpid2[0])
	fmt.Println("GetBounds is: ", x, y, w, h)
	窗口 = 窗口坐标{
		x: x,
		y: y,
	}
	// robotgo.MoveMouse(x, y)

}
