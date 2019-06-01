package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string `json:"name"`
	Age  int
}

func (u User) ReflectCallFuncHasArgs(name string, age int) {
	fmt.Println("ReflectCallFuncHasArgs name: ", name, ", age:", age, "and origal User.Name:", u.Name)
}

func (u User) ReflectCallFuncNoArgs() {
	fmt.Println("ReflectCallFuncNoArgs")
}



// func main() {

// 	user := User{1, "Allen.Wu", 25}

// 	DoFiledAndMethod(user)

// 	getValue := reflect.ValueOf(user)

// 	methodValue := getValue.MethodByName("ReflectCallFuncHasArgs")
// 	args := []reflect.Value{reflect.ValueOf("wuduozhi"),reflect.ValueOf(30)}
// 	methodValue.Call(args)

// 	methodValue = getValue.MethodByName("ReflectCallFuncNoArgs")
// 	args = make([]reflect.Value, 0)
// 	methodValue.Call(args)

// }

// 通过接口来获取任意参数，然后一一揭晓
func DoFiledAndMethod(input interface{}) {

	getType := reflect.TypeOf(input)
	fmt.Println("get Type is :", getType.Name())

	getValue := reflect.ValueOf(input)
	fmt.Println("get all Fields is:", getValue)

	// 获取方法字段
	// 1. 先获取interface的reflect.Type，然后通过NumField进行遍历
	// 2. 再通过reflect.Type的Field获取其Field
	// 3. 最后通过Field的Interface()得到对应的value
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()

		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 获取方法
	// 1. 先获取interface的reflect.Type，然后通过.NumMethod进行遍历
	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		fmt.Printf("%s: %v \n", m.Name, m.Type)
		for  j := 0;j < m.Type.NumIn();j++{
			fmt.Println(m.Type.In(j))  
		}
	}
}

