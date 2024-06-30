package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("Integer")
	default:
		fmt.Println("Unknown", t)
	}
}

func TestTypeAndValue(t *testing.T) {
	var f int64 = 10
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
}

func TestBasicType(t *testing.T) {
	var f float64 = 12
	CheckType(f)
}

type Customer struct {
	CookieID string
	Name     string
	Age      int
}

func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 3: "three"}
	t.Log("a==b", reflect.DeepEqual(a, b))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 3, 1}

	t.Log("s1 == s2?", reflect.DeepEqual(s1, s2))
	t.Log("s1 == s3?", reflect.DeepEqual(s1, s3))

	c1 := Customer{"1", "Mike", 40}
	c2 := Customer{"1", "Mike", 40}
	fmt.Println(c1 == c2)
	fmt.Println(reflect.DeepEqual(c1, c2))
}

type Employee struct {
	EmployeeID string
	Name       string `format:"normal"`
	Age        int
}

func (e *Employee) UpdateAge(newAge int) {
	e.Age = newAge
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{"1", "Mike", 30}
	// 按名字获取成员
	t.Logf("Name: value(%[1]v), Type(%[1]T)", reflect.ValueOf(*e).FieldByName("Name"))
	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("Failed to get 'Name' field.")
	} else {
		t.Log("Tag:format", nameField.Tag.Get("format"))
	}

	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(50)})
	t.Log("Update Age:", e)
}

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello(name string) {
	fmt.Printf("Hello, %s, My name is %s", name, u.Name)
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("xxx")
		return
	}

	v := reflect.ValueOf(o)
	fmt.Println("Fields:")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}
}

func TestUserInfo(t *testing.T) {
	u := User{1, "OK", 12}
	Info(u)
}

func Set(o interface{}) {
	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("xxx")
		return
	} else {
		v = v.Elem()
	}

	f := v.FieldByName("Name")
	if !f.IsValid() {
		fmt.Println("BAD")
		return
	}
	if f.Kind() == reflect.String {
		f.SetString("Hello")
	}
}
func TestUserSet(t *testing.T) {
	u := User{1, "OK", 12}
	Set(&u)
	fmt.Println(u)
}

func TestUserMethod(t *testing.T) {
	u := User{1, "OK", 12}

	v := reflect.ValueOf(u)
	mv := v.MethodByName("Hello")
	args := []reflect.Value{reflect.ValueOf("joe")}
	mv.Call(args)
}
