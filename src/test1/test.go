package main

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

func main() {

	/**
	测试数组
	*/
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 3, 4}
	fmt.Println("arr1 %v", arr1)
	//数据是按值传递
	changeArray(arr1)
	fmt.Println("arr1 %v", arr1)
	fmt.Println("arr2 %v", arr2)

	slices1 := arr1[:]
	fmt.Println("slices1 %v", slices1)
	//切片是按引用传递
	changeSlices(slices1)
	fmt.Println("slices1 %v", slices1)

	/**
	测试切片
	*/
	fmt.Println("1.--------------------------")
	numbers := make([]int, 0, 20)
	printSlices("numbers:", numbers)
	//添加元素
	numbers = append(numbers, 1)
	printSlices("numbers:", numbers)
	//添加多个元素
	numbers = append(numbers, 2, 3, 4, 5, 6, 7)
	printSlices("numbers:", numbers)
	fmt.Println("2.--------------------------")
	//追加一个切片,如果添加元素超过cap，自动扩容一倍
	s1 := []int{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000, 1100, 1200, 1300, 1400}
	numbers = append(numbers, s1...)
	printSlices("numbers:", numbers)
	fmt.Println("3.--------------------------")
	//切片删除元素
	//删除第一个元素
	numbers = numbers[1:]
	printSlices("numbers:", numbers)
	//删除最后一个元素
	numbers = numbers[:len(numbers)-1]
	printSlices("numbers:", numbers)
	//删除中间的一个元素
	a := int(len(numbers) / 2)
	fmt.Println("中间数：", a)
	numbers = append(numbers[:a], numbers[a+1:]...)
	printSlices("numbers:", numbers)
	fmt.Println("4.--------------------------")
	//创建切片numbers1是之前切片的两倍容量
	numbers1 := make([]int, len(numbers), cap(numbers)*2)
	//复制numbers的内容到numbers1
	count := copy(numbers1, numbers)
	fmt.Println("复制个数：", count)
	printSlices("numbers1:", numbers1)
	numbers[len(numbers)-1] = 99
	numbers1[0] = 100

	//两者不存在联系
	printSlices("numbers:", numbers)
	printSlices("numbers1:", numbers1)

	fmt.Println("5.--------------------------")
	var sa []string
	printSliceMsg(sa)

	for i := 0; i < 15; i++ {
		//当使用append追加元素到切片时，如果容量不够，Go就会创建一个新的切片变量来存储元素
		sa = append(sa, strconv.Itoa(i))
		printSliceMsg(sa)
	}

	printSliceMsg(sa)

	/**
	测试map
	*/
	//创建map
	fmt.Println("6.--------------------------")
	countryMap := make(map[string]string)
	countryMap["China"] = "Beijing"
	countryMap["Japan"] = "Tokyo"
	countryMap["India"] = "New Delhi"
	countryMap["France"] = "Paris"
	countryMap["Italy"] = "Rome"

	//遍历map
	for k, v := range countryMap {
		fmt.Println("国家", k, "首都", v)
	}
	fmt.Println("7.--------------------------")
	for _, v := range countryMap {
		fmt.Println("国家", "首都", v)
	}
	fmt.Println("8.--------------------------")
	for k := range countryMap {
		fmt.Println("国家", k, "首都", countryMap[k])
	}
	fmt.Println("9.--------------------------")

	if value, ok := countryMap["England"]; ok {
		fmt.Println("首都：", value)
	} else {
		fmt.Println("首都未检索到")
	}

	fmt.Println("10.--------------------------")
	if _, ok := countryMap["Italy"]; ok {
		delete(countryMap, "Italy")
	}

	fmt.Println("删除后:", countryMap)

	fmt.Println("11.--------------------------")
	//map是按引用传参
	changeMap(countryMap)

	fmt.Println("修改后:", countryMap)

	fmt.Println("12.--------------------------")

	//清空map
	countryMap = make(map[string]string)

	fmt.Println("清空后:", countryMap)

	fmt.Println("13.--------------------------")
	s := "我爱GO语言"
	len := 0
	//遍历字符串
	for i, ch := range s {
		fmt.Printf("%d:%c", i, ch)
		len++
	}
	fmt.Println("\n字符串长度", len)

	fmt.Println("14.--------------------------")
	//遍历所有字节
	for i, ch := range []byte(s) {
		fmt.Printf("%d:%x", i, ch)
	}
	fmt.Println()

	fmt.Println("15.--------------------------")

	//遍历所有字符,如果有汉字，遍历时推荐用rune
	count = 0
	for i, ch := range []rune(s) {
		fmt.Printf("%d:%c", i, ch)
		count++
	}
	fmt.Println("\n字符串长度", count)
	fmt.Println("字符串长度", utf8.RuneCountInString(s))

	fmt.Println("16.--------------------------")
	/**
	随机数
	*/
	//指定种子
	rand.Seed(time.Now().UnixNano())
	//随机一个非负整数
	fmt.Println(rand.Int())
	//0到10的随机数
	fmt.Println(rand.Intn(10))
	//0.0-1.0的随机数
	fmt.Println(rand.Float64())
	//5-10之间的随机数
	fmt.Println(rand.Intn(10-5+1) + 5)

	/**
	字符串搜索
	*/
	fmt.Println("17.--------------------------")
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Index("seafood", "foo"))
	fmt.Println(strings.Index("GO语言学习", "言"))
	fmt.Println(strings.IndexRune("GO语言学习", '言'))
	fmt.Println(strings.IndexAny("GO语言学习", "abc语习"))

	/**
	分隔字符串
	*/
	fmt.Println("18.--------------------------")
	fmt.Println(strings.Fields(" abc 123 ABC xyz  XYZ \t bbb"))
	fmt.Println(strings.Split("a,b,c", ","))
	fmt.Println(strings.FieldsFunc("a,b,c", func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}))

	/**
	分割字符串
	*/
	fmt.Println("19.--------------------------")
	fmt.Println(strings.Trim(" steven wang ", " "))
	fmt.Println(strings.TrimFunc("a,b,c", func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}))
	fmt.Println(strings.TrimSpace(" steven wang "))

	/**
	比较字符串
	*/
	fmt.Println("20.--------------------------")
	fmt.Println(strings.Compare("abc", "bcd"))
	fmt.Println("abc" < "bcd")
	//忽略大小写
	fmt.Println(strings.EqualFold("Go", "go"))
	//指定重复次数
	fmt.Println("g" + strings.Repeat("o", 2) + "le")
	//替换
	fmt.Println(strings.Replace("王老大 王老二 王老三", "王", "张", 2))
	fmt.Println(strings.Replace("王老大 王老二 王老三", "王", "张", -1))
	//连接
	s2 := []string{"abc", "ABC", "123"}
	fmt.Println(strings.Join(s2, ","))

	/**
	Parse类函数
	*/
	fmt.Println("21.--------------------------")
	//字符串转整型
	aa, _ := strconv.Atoi("100")
	fmt.Println(aa)
	//二进制转int64
	num, _ := strconv.ParseInt("010001", 2, 64)
	fmt.Println(num)
	pi := "3.1415926"
	num1, _ := strconv.ParseFloat(pi, 64)
	fmt.Printf("%T,%v\n", num1, num1*2)

	flag, _ := strconv.ParseBool("false")
	fmt.Printf("%T , %v\n", flag, flag)

	/**
	Format
	*/
	fmt.Println("21.--------------------------")
	//整型转字符串
	fmt.Println(strconv.Itoa(199))
	//转16进制后，再转字符串
	fmt.Println(strconv.FormatInt(-19968, 16))

	/**
	regexp正则
	*/
	fmt.Println("22.--------------------------")
	//Match检查正则表达式是否与字节切片匹配
	flag1, _ := regexp.Match("^\\d{6,7}$", []byte("0123456789"))
	fmt.Println(flag1)
	//MatchString检查字符串是否匹配
	flag2, _ := regexp.MatchString("^\\d{6,7}$", "0123456789")
	fmt.Println(flag2)
	//Compile编译为正则表达式对象
	RegExp, _ := regexp.Compile("^\\d{6}\\D{2}$")
	//MustCompile编译为正则表达式对象，但不返回error
	RegExp2 := regexp.MustCompile("^\\d{6}\\D{2}$")
	//Match判断正则表达式是否匹配
	flag = RegExp2.Match([]byte("一丁"))
	fmt.Println(flag)
	//MatchString判断字符串是否匹配
	flag = RegExp.MatchString("123456ab")
	fmt.Println(flag)
	//ReplaceAll替换
	text := "字符串 123 正则 456 abc"
	RegExp3 := regexp.MustCompile("[\\d\\s]+")
	result := string(RegExp3.ReplaceAll([]byte(text), []byte("-")))
	fmt.Println(result)
	text = "第一部分#第二部分#第三部分#第四部分"
	MyRegexp := regexp.MustCompile("#+")
	arr := MyRegexp.Split(text, 5)
	fmt.Println(arr)

	/**
	time
	*/
	fmt.Println("23.--------------------------")
	time1 := time.Now()
	fmt.Println(time1)
	fmt.Println(time.Local)
	fmt.Println(time.UTC)
	//字符串转时间，第一个参数是参考格式，第二个是需要转的字符串
	t, _ := time.Parse("2000-16-01 19:40:40", "2020-11-18 17:40:40")
	fmt.Println(t)
	//时间转字符串，参数是格式
	fmt.Println(time.Now().Format("2000-16-01 19:40:40"))
	//时间转字符串
	fmt.Println(time.Now().String())
	//时间转时间戳，到1970年的秒
	fmt.Println(time.Now().Unix())
	//时间转时间戳，到1970年的纳秒
	fmt.Println(time.Now().UnixNano())
	//时间相等
	fmt.Println(time1.Equal(time.Now()))
	//时间早
	fmt.Println(time1.Before(time.Now()))
	//时间晚
	fmt.Println(time1.After(time.Now()))
	//年、月、日
	fmt.Println(time.Now().Year())
	fmt.Println(time.Now().Month())
	fmt.Println(time.Now().Day())
	//时间差
	fmt.Println(time.Now().Sub(time1))
	fmt.Println(time.Now().Sub(time1).Seconds())

	d, _ := time.ParseDuration("1h30m")
	//增加时间
	fmt.Println(time.Now().Add(d))
	//减少时间
	fmt.Println(time.Now().Add(-d))
	//增加1天
	fmt.Println(time.Now().AddDate(0, 0, 1))

	/**
	math
	*/
	fmt.Println("24.--------------------------")
	//判断是否为空
	fmt.Println(math.IsNaN(3.14))
	//取顶
	fmt.Println(math.Ceil(1.000001))
	//取底
	fmt.Println(math.Floor(1.999991))

	username := ""
	age := 0
	//输入
	//fmt.Scanln(&username,&age)
	//输出
	fmt.Println(username, age)

}

func printSlices(name string, x []int) {
	fmt.Print(name, "\t")
	fmt.Printf("addr:%p \t len=%d \t cap=%d \t slice=%v\n", x, len(x), cap(x), x)
}

func changeArray(a [4]int) {
	a[1] = 250
}

func changeSlices(a []int) {
	a[1] = 250
}
func printSliceMsg(x []string) {
	fmt.Printf("addr:%p \t len=%d \t cap=%d \t slice=%v\n", x, len(x), cap(x), x)
}
func changeMap(m map[string]string) {
	m["USA"] = "Washington"
}
