package telegraph

import "fmt"

// Helper function to easily call EditAccountInfo by an account.
func (a *Account) EditInfo(opts *EditAccountInfoOpts) (*Account, error) {
	return EditAccountInfo(a.AccessToken, opts)
}

// Helper function to easily call GetAccountInfo by an account.
func (a *Account) GetInfo() (*Account, error) {
	return GetAccountInfo(a.AccessToken)
}

// Helper function to easily call RevokeAccessToken by an account.
func (a *Account) RevokeAccessToken() (*Account, error) {
	return RevokeAccessToken(a.AccessToken)
}

// Helper function to easily call CreatePage by an account.
func (a *Account) CreatePage(title string, content string, opts *PageOpts) (*Page, error) {
	return CreatePage(a.AccessToken, title, content, opts)
}

// Helper function to easily call EditPage by an account with previous author_name and author_url.
func (a *Account) EditPage(path string, title string, content string, opts *PageOpts) (*Page, error) {
	fmt.Println("Access Token:", a.AccessToken)
	return EditPage(a.AccessToken, path, title, content, opts)
}

// Helper function to easily call GetPageList by an account.
func (a *Account) GetPageList(opts *PageListOpts) (*PageList, error) {
	return GetPageList(a.AccessToken, opts)
}

// Helper function to easily get page.
func (p *Page) Get(returnContent bool) (*Page, error) {
	return GetPage(p.Path, returnContent)
}

// Helper function to easily call GetViews in a page.
func (p *Page) GetViews(opts *PageViewsOpts) (*PageViews, error) {
	return GetViews(p.Path, opts)
}
