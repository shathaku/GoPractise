package main

import "fmt"

func main() {
	var a int = 0;
	var b int = 1;
	var temp int = 0;

	for i:=1;i<=10;i++ {
		if(i==1) {
			fmt.Println(a);
			continue
		}
		if(i==2) {
			fmt.Println(b);
			continue;
		}
		temp = a+b;
		a = b;
		b = temp;
		fmt.Println(temp)
	}
}