package telegraph

import (
	"net/http"
)

// GetTelegraphClient returns a new TelegraphClient using the specified options.
func GetTelegraphClient(options *ClientOpt) *TelegraphClient {
	if options == nil {
		options = &ClientOpt{}
	}
	if options.HttpClient == nil {
		options.HttpClient = http.DefaultClient
	}
	if options.ApiUrl == "" {
		options.ApiUrl = "https://api.telegra.ph/"
	}
	return &TelegraphClient{
		HttpClient: options.HttpClient,
	}
}

// EditInfo is a helper method to easily call EditAccountInfo by an account.
func (a *Account) EditInfo(client *TelegraphClient, opts *EditAccountInfoOpts) (*Account, error) {
	return client.EditAccountInfo(a.AccessToken, opts)
}

// GetInfo is a helper method to easily call GetAccountInfo by an account.
func (a *Account) GetInfo(client *TelegraphClient) (*Account, error) {
	return client.GetAccountInfo(a.AccessToken)
}

// RevokeAccessToken is a helper method to easily call RevokeAccessToken by an account.
func (a *Account) RevokeAccessToken(client *TelegraphClient) (*Account, error) {
	return client.RevokeAccessToken(a.AccessToken)
}

// CreatePage is a helper method to easily call CreatePage by an account.
func (a *Account) CreatePage(client *TelegraphClient, title, content string, opts *PageOpts) (*Page, error) {
	return client.CreatePage(a.AccessToken, title, content, opts)
}

// EditPage is a helper method to easily call EditPage by an account with previous author_name and author_url.
func (a *Account) EditPage(client *TelegraphClient, path, title, content string, opts *PageOpts) (*Page, error) {
	return client.EditPage(a.AccessToken, path, title, content, opts)
}

// GetPageList is a helper method to easily call GetPageList by an account.
func (a *Account) GetPageList(client *TelegraphClient, opts *PageListOpts) (*PageList, error) {
	return client.GetPageList(a.AccessToken, opts)
}

// Get is a helper method to easily get page.
func (p *Page) Get(client *TelegraphClient, returnContent bool) (*Page, error) {
	return client.GetPage(p.Path, returnContent)
}

// GetViews helper method to easily call GetViews in a page.
func (p *Page) GetViews(client *TelegraphClient, opts *PageViewsOpts) (*PageViews, error) {
	return client.GetViews(p.Path, opts)
}
