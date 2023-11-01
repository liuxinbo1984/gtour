package main

import "fmt"

func main() {
	var p person
	fmt.Printf("%+v\n", p)

	// p1 := person{"雪山飞狐", 16}
	// fmt.Println(p1, p1.name, p1.age)

	p2 := person{age: 29, name: "shadow"}
	fmt.Println(p2)

	p3 := person{
		name: "arvin",
		age:  30,
		address: address{
			"shanxi",
			"xi'an",
		},
	}

	// p3.province <=> p3.address.province
	fmt.Println(p3, p3.province, p3.address.city)

	// 接口引用
	// var pi fmt.Stringer = p3
	// fmt.Println(pi.String())
	printStringer(&p3)

	// 类型断言加转换
	var s fmt.Stringer = &p3
	pp := s.(*person)
	fmt.Printf("type: %T, value:%v\n", pp, pp)
}

type person struct {
	name string
	age  int

	address
}

type address struct {
	province string
	city     string
}

// 接收者为指针，指针类型的接收者才实现了接口，调用时也只能使用指针类型；
// 如果接收者为类型，指针类型和实例类型都实现了该接口
func (p *person) String() string {
	return fmt.Sprintf("the name is %s, age is %d", p.name, p.age)
}

func printStringer(param fmt.Stringer) {
	fmt.Println(param.String())
}
