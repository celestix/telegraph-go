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
	_, err = a.CreatePage("Sample", `<h3>Sample Page #1</h3> <p>Hello world! This telegraph page is created using telegraph-go package.</p><br><a href="https://github.com/anonyindian/telegraph-go">Click here to open package</a>`, &telegraph.PageOpts{
		AuthorName: "User1",
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = a.CreatePage("Sample", `<h3>Sample Page #2</h3> <p>Hello world! This telegraph page is created using telegraph-go package.</p>`, &telegraph.PageOpts{
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
