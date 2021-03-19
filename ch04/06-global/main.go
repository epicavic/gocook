package main

import "main/global"

func main() {
	if err := global.UseLog(); err != nil {
		panic(err)
	}
}
