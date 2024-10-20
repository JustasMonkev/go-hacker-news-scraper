package main

import (
	"GoooooShoter/browser"
	"GoooooShoter/helper"
	"fmt"
	"os"
	"time"
)

func main() {
	context := browser.GetContext()

	now := time.Now()

	file := helper.Check(os.Create(now.Format("2006-01-02") + "-demo.txt"))

	defer file.Close()

	helper.Check(fmt.Fprint(file, context))
}
