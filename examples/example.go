package main

import (
	"fmt"

	"github.com/anonyindian/telegraph-go"
)

func main() {
	//Use this method to create account
	a, err := telegraph.CreateAccount("tgraph-go", &telegraph.CreateAccountOpts{
		AuthorName: "Telegraph Go Package",
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// The Telegraph API uses a DOM-based format to represent the content of the page.
	// https://telegra.ph/api#Content-format
	_, err = a.CreatePage("Sample Page 1", `[{"tag":"h3", "children":["A Sample Page #1"]}, {"tag":"p", "children":["Hello world! This telegraph page is created using telegraph-go package."]}]`, &telegraph.PageOpts{
		AuthorName: "User1",
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = a.CreatePage("Sample Page 2", `[{"tag":"h3", "children":["A Sample Page #2"]}, {"tag":"p", "children":["Hello world! This telegraph page is created using telegraph-go package."]}]`, &telegraph.PageOpts{
		AuthorName: "User1",
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	// Get a list of pages in your current account with this method
	plist, _ := a.GetPageList(nil)
	for _, page := range plist.Pages {
		// you can print all pages with the help of loop
		fmt.Println(page.Url)
	}

	// Get total pages count in this way
	pcount := plist.TotalCount
	fmt.Println(pcount)
}
