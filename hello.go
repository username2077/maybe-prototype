package main

import "fmt"

func main() {
	//fmt print to cmd
	fmt.Println("Hello World!")

	//fmt Sprintf
	var num = 123
	var date = "2021-9-25"
	var url = "num = %d date = %s"
	var target_url = fmt.Sprintf(url, num, date)
	fmt.Println(target_url)

	//var v_name v_type = value
	var a string = "worst day"
	fmt.Println(a)
	var b int = 12345678987654321
	fmt.Println(b)

	//no value
	var c int//0
	fmt.Println(c)
	var d bool//false
	fmt.Println(d)

	//more clean way var
	//try to use this way just in simple function
	intVal := true
	fmt.Println(intVal)

	//muti var
	//globar
	var (
		e int
		f bool
	)
	//other
	var g, h int = 3, 4
	fmt.Println(e, f, g, h)

	//address
	fmt.Println(&a, &b, &c, &d, &e, &f, &intVal)
}