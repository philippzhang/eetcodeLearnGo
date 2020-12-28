package main

import (
	"fmt"
	"math"
)

//定义Teacher
type Teacher struct {
	name string
	age  int8
	sex  byte
}
type Address struct {
	province, city string
}
type Person struct {
	name    string
	age     int
	address *Address
}
type Employee struct {
	name, currency string
	salar          float64
}
type Rectangle struct {
	width, height float64
}
type Circle struct {
	radius float64
}

type Human struct {
	name, phone string
	age         int
}
type Student struct {
	Human  //匿名字段
	school string
}
type Employee2 struct {
	Human   //匿名字段
	company string
}

type Phone interface {
	call()
}
type AndroidPhone struct {
}
type IPhone struct {
}

func (a AndroidPhone) call() {
	fmt.Println("我是安卓手机，可以打电话了")
}
func (i IPhone) call() {
	fmt.Println("我是苹果手机，可以打电话了")
}

type ISayHello interface {
	SayHello() string
}
type Duck struct {
	name string
}
type People struct {
	name string
}

func (d Duck) SayHello() string {
	return d.name + "叫：ga ga ga!"
}
func (p People) SayHello() (t string) {
	t = p.name + "说：Hello!"
	return
}

type Income interface {
	calculate() float64
	source() string
}
type A1 struct {
	proj   string
	amount float64
}
type A2 struct {
	proj string
	hour float64
	rate float64
}

func (f A1) calculate() float64 {
	return f.amount
}
func (f A1) source() string {
	return f.proj
}
func (f A2) calculate() float64 {
	return f.hour * f.rate
}
func (f A2) source() string {
	return f.proj
}

//空接口
type A interface {
}

type Cat struct {
	name string
	age  int
}
type Pig struct {
	name string
	sex  int
}

//定义接口
type Shape interface {
	perimeter() float64
	area() float64
}

//三角形
type Triangle struct {
	a, b, c float64
}

func (r Rectangle) perimeter() float64 {
	return (r.height + r.width) * 2
}
func (r Rectangle) area() float64 {
	return r.height * r.width
}
func (t Triangle) perimeter() float64 {
	return t.a + t.b + t.c
}
func (t Triangle) area() float64 {
	//海伦公式
	p := t.perimeter()
	return math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
}
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func (c Circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

func main() {
	//声明方式实例化结构体
	fmt.Println("1.--------------------------")
	var t1 Teacher
	fmt.Println(t1)
	fmt.Printf("t1:%T , %v , %q \n", t1, t1, t1)
	t1.name = "Steven"
	t1.age = 35
	t1.sex = 1
	fmt.Println(t1)
	fmt.Printf("t1:%T , %v , %q \n", t1, t1, t1)
	//遍历简短声明
	t2 := Teacher{}
	t2.name = "David"
	t2.age = 30
	t2.sex = 1
	fmt.Println(t2)
	//声明初始化
	t3 := Teacher{
		name: "Josh",
		age:  28,
		sex:  1,
	}
	fmt.Println(t3)
	t3 = Teacher{name: "Josh2", age: 27, sex: 1}
	fmt.Println(t3)
	t4 := Teacher{"Ruby", 30, 0}
	fmt.Println(t4)

	//用new实例化struct
	t5 := new(Teacher)
	(*t5).name = "David"
	(*t5).age = 30
	(*t5).sex = 1
	fmt.Printf("t5:%T , %v , %q \n", t5, t5, t5)

	t5.name = "David2"
	t5.age = 31
	t5.sex = 1
	fmt.Printf("t5:%T , %v , %q \n", t5, t5, t5)

	fmt.Println("2.--------------------------")
	fmt.Printf("修改前:%T %v %p \n", t4, t4, &t4)
	//struct是值类型
	changeName(t4)

	fmt.Printf("修改后:%T %v %p \n", t4, t4, &t4)

	fmt.Printf("修改前:%T %v %p \n", t5, t5, &t5)
	//struct是值类型
	changeName(*t5)

	fmt.Printf("修改后:%T %v %p \n", t5, t5, &t5)

	fmt.Println("3.--------------------------")

	d1 := Teacher{"豆豆", 40, 1}
	fmt.Printf("d1:%T %v %p \n", d1, d1, &d1)
	//深拷贝
	d2 := d1
	fmt.Printf("d2:%T %v %p \n", d2, d2, &d2)
	d2.name = "毛毛"
	fmt.Printf("d1:%T %v %p \n", d1, d1, &d1)
	fmt.Printf("d2:%T %v %p \n", d2, d2, &d2)

	fmt.Println("4.--------------------------")
	//浅拷贝：直接赋值指针地址
	d3 := &d1
	fmt.Printf("d3:%T %v %p \n", d3, d3, &d3)
	d3.name = "球球"
	d3.age = 20
	fmt.Printf("d1:%T %v %p \n", d1, d1, &d1)
	fmt.Printf("d3:%T %v %p \n", d3, d3, &d3)

	fmt.Println("5.--------------------------")
	//浅拷贝:通过new()实例化
	d4 := new(Teacher)
	d4.name = "多多"
	d4.age = 31
	d4.sex = 0
	d5 := d4
	fmt.Printf("d4:%T %v %p \n", d4, d4, &d4)
	fmt.Printf("d5:%T %v %p \n", d5, d5, &d5)
	d5.name = "贝贝"
	d5.age = 40
	fmt.Printf("d4:%T %v %p \n", d4, d4, &d4)
	fmt.Printf("d5:%T %v %p \n", d5, d5, &d5)

	fmt.Println("6.--------------------------")
	//匿名函数
	res := func(a, b float64) float64 {
		return math.Pow(a, b)
	}(2, 3)

	fmt.Println(res)

	fmt.Println("7.--------------------------")
	//匿名结构体
	addr := struct {
		province, city string
	}{"陕西省", "西安市"}
	fmt.Println(addr)

	cat := struct {
		name, color string
		age         int8
	}{
		name:  "绒毛",
		color: "黑白",
		age:   1,
	}
	fmt.Println(cat)

	e1 := Teacher{"Steven", 30, 1}
	fmt.Println(e1)

	fmt.Println("8.--------------------------")
	//结构体嵌套
	p := Person{}
	p.name = "Steven"
	p.age = 35
	addr1 := Address{}
	addr1.province = "北京市"
	addr1.city = "海地区"
	p.address = &addr1
	fmt.Println(p)
	fmt.Println("性别：", p.name, ", 年龄:", p.age, ",省：", p.address.province, ",市:"+p.address.city)

	emp1 := Employee{"Daniel", "$", 2000}
	emp1.printSalary()

	r1 := Rectangle{10, 4}
	c1 := Circle{1}

	fmt.Println("r1的面积:", r1.Area())
	fmt.Println("c1的面积:", c1.Area())

	r2 := Rectangle{5, 8}
	r3 := r2
	fmt.Printf("r2的地址: %p \n", &r2)
	fmt.Printf("r3的地址: %p \n", &r3)
	r2.setValue()
	fmt.Println(r2)
	fmt.Println(r3)

	r2.setValue2()
	fmt.Println(r2)
	fmt.Println(r3)

	fmt.Println("9.--------------------------")
	//继承
	s1 := Student{Human{"Daniel", "120000000", 13}, "十一中学"}
	s2 := Employee2{Human{"Steven", "11111111", 35}, "111sssss"}

	s1.SayHi() //这里调用 (h *Student) SayHi() 方法重写采用就近原则
	s2.SayHi()

	fmt.Println("10.--------------------------")
	//定义接口类型
	var phone Phone
	phone = new(AndroidPhone)
	fmt.Printf("%T, %v, %p\n", phone, phone, &phone)
	phone.call()
	phone = AndroidPhone{}
	fmt.Printf("%T, %v, %p\n", phone, phone, &phone)
	phone.call()
	phone = new(IPhone) //隐士赋值，地址没有变
	fmt.Printf("%T, %v, %p\n", phone, phone, &phone)
	phone.call()
	phone = IPhone{}
	fmt.Printf("%T, %v, %p\n", phone, phone, &phone)
	phone.call()

	fmt.Println("11.--------------------------")
	var phone2 Phone
	phone2 = AndroidPhone{}
	fmt.Printf("%T, %v, %p\n", phone2, phone2, &phone2)
	phone2.call()
	phone2 = IPhone{}
	fmt.Printf("%T, %v, %p\n", phone2, phone2, &phone2)
	phone2.call()

	fmt.Println("12.--------------------------")
	duck := Duck{"Yaya"}
	people := People{"Steven"}
	fmt.Println(duck.SayHello())
	fmt.Println(people.SayHello())

	//定义接口
	var iI ISayHello

	//不用继承，执行实现接口的所有方法即可
	iI = duck
	fmt.Printf("%T, %v, %p\n", iI, iI, &iI)
	fmt.Println(iI.SayHello())

	iI = people
	fmt.Printf("%T, %v, %p\n", iI, iI, &iI)
	fmt.Println(iI.SayHello())

	fmt.Println("13.--------------------------")
	//多态
	pp1 := A1{"项目1", 5000}
	pp2 := A2{"项目2", 100, 40}

	ic := []Income{pp1, pp2}
	fmt.Println(calculate(ic))

	fmt.Println("14.--------------------------")
	//空接口
	var aa1 A = Cat{"Mimi", 1}
	var aa2 A = Pig{"Xiaohei", 1}
	var aa3 A = "Learn golang with me"
	var aa4 A = 100
	var aa5 A = 3.14
	showInfo(aa1)
	showInfo(aa2)
	showInfo(aa3)
	showInfo(aa4)
	showInfo(aa5)

	//定义map,value是任何数据类型
	map1 := make(map[string]interface{})
	map1["name"] = "Daniel"
	map1["age"] = 13
	map1["height"] = 1.71
	fmt.Println(map1)

	//定义切片，存储是任意数据类型
	slice1 := make([]interface{}, 0, 10)
	slice1 = append(slice1, aa1, aa2, aa3, aa4, aa5)
	fmt.Println(slice1)

	fmt.Println("15.--------------------------")

	//类型转换
	//instance , ok := 接口对象.(实际类型)
	//接口对象.(type)

	var ss Shape
	ss = Rectangle{3, 4}
	getResult(ss)
	showInfo(ss)

	ss = Triangle{3, 4, 5}
	getResult(ss)
	showInfo(ss)

	ss = Circle{1}
	getResult(ss)
	showInfo(ss)

	xx := Triangle{3, 4, 5}
	fmt.Println(xx)

}

func changeName(t Teacher) {
	t.name = "Daniel"
	t.age = 13
	fmt.Printf("函数体内修改:%T %v %p \n", t, t, &t)

}

/**
方法
*/
func (e Employee) printSalary() {
	fmt.Printf("员工姓名%s,薪资：%s%.2f\n", e.name, e.currency, e.salar)
}

//定义Rectangle的方法
func (r Rectangle) Area() float64 {
	return r.width * r.height
}
func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}
func (r Rectangle) setValue() {
	r.height = 10
}
func (r *Rectangle) setValue2() {
	r.height = 20
}

func (h *Human) SayHi() {
	fmt.Printf("Human：%s ,%d %s\n", h.name, h.age, h.phone)
}

func (h *Student) SayHi() {
	fmt.Printf("Student：%s ,%d %s\n", h.name, h.age, h.phone)
}

func calculate(ic []Income) float64 {
	sum := 0.0
	for _, income := range ic {
		fmt.Printf("%s , %.2f\n", income.source(), income.calculate())
		sum += income.calculate()
	}

	return sum
}

func showInfo(a A) {
	fmt.Printf("%T , %v \n", a, a)
}

//接口对象转型方式1
func getType(s Shape) {
	if instance, ok := s.(Rectangle); ok {
		fmt.Printf("矩形：长度%.2f ,宽度%.2f ,", instance.height, instance.width)
	} else if instance, ok := s.(Triangle); ok {
		fmt.Printf("三角形：三边分别为 %.2f, %.2f, %.2f ,", instance.a, instance.b, instance.c)
	} else if instance, ok := s.(Circle); ok {
		fmt.Printf("圆形：半径%.2f ", instance.radius)
	}
}

//接口类型转型方式2
func getType2(s Shape) {
	switch instance := s.(type) {
	case Rectangle:
		fmt.Printf("矩形：长度%.2f ,宽度%.2f ,", instance.height, instance.width)
	case Triangle:
		fmt.Printf("三角形：三边分别为 %.2f, %.2f, %.2f ,", instance.a, instance.b, instance.c)
	case Circle:
		fmt.Printf("圆形：半径%.2f ", instance.radius)
	}
}
func getResult(s Shape) {
	//getType(s)
	getType2(s)
	fmt.Printf("周长：%.2f , 面积: %.2f \n", s.perimeter(), s.area())
}

func (t Triangle) String() string {
	return fmt.Sprintf("Triangle对象,属性分别为: %.2f, %.2f %.2f", t.a, t.b, t.c)
}
