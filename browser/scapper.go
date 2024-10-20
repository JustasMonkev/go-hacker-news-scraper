package browser

import (
	"GoooooShoter/helper"
	"fmt"
	"github.com/playwright-community/playwright-go"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	getter "github.com/hashicorp/go-getter"
)

func GetContext() string {
	var array []string
	var context string
	pw := helper.Check(playwright.Run())

	browser := helper.Check(pw.Chromium.Launch())

	page := helper.Check(browser.NewPage())

	helper.Check(page.Goto("https://news.ycombinator.com"))

	entries := helper.Check(page.Locator(".athing").All())
	for _, entry := range entries {
		link := helper.Check(entry.Locator("td.title > span > a").GetAttribute("href"))
		array = append(array, link)
	}

	for i := 0; i < len(array); i++ {
		if !strings.Contains(array[i], "pdf") {
			helper.Check(page.Goto(array[i], playwright.PageGotoOptions{
				WaitUntil: playwright.WaitUntilStateDomcontentloaded,
			}))

			content := helper.Check(page.InnerText("body"))
			content = strings.ReplaceAll(content, "\n", "") // Remove newlines
			content = strings.ReplaceAll(content, "\t", "") // Remove tabs
			context += content
		} else {
			filePath := "file_" + strconv.Itoa(i) + ".pdf"
			projectDir := helper.Check(os.Getwd())

			destPath := filepath.Join(projectDir, filePath)

			client := &getter.Client{
				Src:  array[i],
				Dst:  destPath,
				Mode: getter.ClientModeFile,
			}

			if err := client.Get(); err != nil {
				fmt.Printf("Error downloading file: %s\n", err)
			}
		}
	}

	// Clean up
	helper.CheckErr(browser.Close())
	helper.CheckErr(pw.Stop())

	return context
}
