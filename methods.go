package telegraph

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

// Use this method to create a new Telegraph account. Most users only need one account, but this can be useful for channel administrators who would like to keep individual author names and profile links for each of their channels.
// On success, returns an Account object with the regular fields and an additional access_token field.
// - shortName (type string): Account name, helps users with several accounts remember which they are currently using. Displayed to the user above the "Edit/Publish" button on Telegra.ph, other users don't see this name.
// - opts (type CreateAccountOpts): All optional parameters.
// https://telegra.ph/api#createAccount
func CreateAccount(shortName string, opts *CreateAccountOpts) (*Account, error) {
	var (
		u = url.Values{}
		a Account
	)
	u.Add("short_name", shortName)

	if opts != nil {
		u.Add("author_name", opts.AuthorName)
		u.Add("author_url", opts.AuthorUrl)
	}

	r, err := Get("createAccount", u)
	if err != nil {
		return nil, err
	}
	return &a, json.Unmarshal(r, &a)
}

// Use this method to update information about a Telegraph account. Pass only the parameters that you want to edit.
// On success, returns an Account object with the default fields.
// - accessToken (type string): Access token of the Telegraph account.
// - opts (type EditAccountInfoOpts): All optional parameters.
// https://telegra.ph/api#editAccountInfo
func EditAccountInfo(accessToken string, opts *EditAccountInfoOpts) (*Account, error) {
	var (
		u = url.Values{}
		a Account
	)
	u.Add("access_token", accessToken)
	if opts != nil {
		if opts.ShortName != "" {
			u.Add("short_name", opts.ShortName)
		}
		if opts.AuthorName != "" {
			u.Add("author_name", opts.AuthorName)
		}
		if opts.AuthorUrl != "" {
			u.Add("author_url", opts.AuthorUrl)
		}
	}

	r, err := Get("editAccountInfo", u)
	if err != nil {
		return nil, err
	}
	return &a, json.Unmarshal(r, &a)
}

// Use this method to get information about a Telegraph account.
// Returns an Account object on success.
// - accessToken (type string): Access token of the Telegraph account.
// https://telegra.ph/api#getAccountInfo
func GetAccountInfo(accessToken string) (*Account, error) {
	var (
		u = url.Values{}
		a Account
	)
	u.Add("access_token", accessToken)
	u.Add("fields", `["short_name", "author_name", "author_url", "auth_url", "page_count"]`)

	r, err := Get("getAccountInfo", u)
	if err != nil {
		return nil, err
	}
	return &a, json.Unmarshal(r, &a)
}

// Use this method to revoke access_token and generate a new one, for example, if the user would like to reset all connected sessions, or you have reasons to believe the token was compromised.
// On success, returns an Account object with new access_token and auth_url fields.
// - accessToken (type string): Access token of the Telegraph account.
// https://telegra.ph/api#revokeAccessToken
func RevokeAccessToken(accessToken string) (*Account, error) {
	var (
		u = url.Values{}
		a Account
	)
	u.Add("access_token", accessToken)

	r, err := Get("revokeAccessToken", u)
	if err != nil {
		return nil, err
	}
	return &a, json.Unmarshal(r, &a)
}

// Use this method to create a new Telegraph page.
// On success, returns a Page object.
// - accessToken (type string): Access token of the Telegraph account.
// - title (type string): Page title.
// - content (type string): Content of the page (Array of Node, up to 64 KB converted into a json string).
// - opts (type PageOpts): All optional parameters.
// https://telegra.ph/api#createPage
func CreatePage(accessToken string, title string, content string, opts *PageOpts) (*Page, error) {
	var (
		u = url.Values{}
		a Page
	)
	u.Add("access_token", accessToken)
	u.Add("title", title)
	cNode, err := ContentFormat(content)
	if err != nil {
		return nil, err
	}
	cNodeB, err := json.Marshal(cNode)
	if err != nil {
		return nil, err
	}
	u.Add("content", string(cNodeB))

	if opts != nil {
		u.Add("author_name", opts.AuthorName)
		u.Add("author_url", opts.AuthorUrl)
		u.Add("return_content", strconv.FormatBool(opts.ReturnContent))
	}

	r, err := Get("createPage", u)
	if err != nil {
		return nil, err
	}
	return &a, json.Unmarshal(r, &a)
}

// Use this method to edit an existing Telegraph page.
// On success, returns a Page object.
// - accessToken (type string): Access token of the Telegraph account.
// - path (type string): Path to the page.
// - title (type string): Page title.
// - content (type string): Content of the page (Array of Node, up to 64 KB converted into a json string).
// - opts (type PageOpts): All optional parameters.
// https://telegra.ph/api#editPage
func EditPage(accessToken string, path string, title string, content string, opts *PageOpts) (*Page, error) {
	var (
		u = url.Values{}
		a Page
	)
	u.Add("access_token", accessToken)
	u.Add("path", path)
	u.Add("title", title)
	cNode, err := ContentFormat(content)
	if err != nil {
		return nil, err
	}
	cNodeB, err := json.Marshal(cNode)
	if err != nil {
		return nil, err
	}
	u.Add("content", string(cNodeB))

	if opts != nil {
		u.Add("author_name", opts.AuthorName)
		u.Add("author_url", opts.AuthorUrl)
		u.Add("return_content", strconv.FormatBool(opts.ReturnContent))
	}

	r, err := Get("editPage", u)
	if err != nil {
		return nil, err
	}
	return &a, json.Unmarshal(r, &a)
}

// Use this method to get a Telegraph page.
// On success, returns a Page object.
// - path (type string): Path to the Telegraph page (in the format Title-12-31, i.e. everything that comes after http://telegra.ph/).
// - returnContent (type bool): If true, content field will be returned in Page object.
// https://telegra.ph/api#getPage
func GetPage(path string, returnContent bool) (*Page, error) {
	var (
		u = url.Values{}
		a Page
	)
	u.Add("path", path)
	u.Add("return_content", strconv.FormatBool(returnContent))

	r, err := Get("getPage", u)
	if err != nil {
		return nil, err
	}
	return &a, json.Unmarshal(r, &a)
}

// Use this method to get a list of pages belonging to a Telegraph account.
// Returns a PageList object, sorted by most recently created pages first.
// - accessToken (type string): Access token of the Telegraph account.
// - opts
// https://telegra.ph/api#getPageList
func GetPageList(accessToken string, opts *PageListOpts) (*PageList, error) {
	var (
		u = url.Values{}
		a PageList
	)
	u.Add("access_token", accessToken)
	if opts != nil {
		if opts.Offset != 0 {
			u.Add("offset", strconv.FormatInt(opts.Offset, 10))
		}
		if opts.Limit != 0 {
			u.Add("limit", strconv.FormatInt(opts.Limit, 10))
		}
	}

	r, err := Get("getPageList", u)
	if err != nil {
		return nil, err
	}
	return &a, json.Unmarshal(r, &a)
}

// Use this method to get the number of views for a Telegraph article.
// Returns a PageViews object on success. By default, the total number of page views will be returned.
// - path (type string): Path to the Telegraph page (in the format Title-12-31, i.e. everything that comes after http://telegra.ph/).
// - opts (type PageViewsOpts): All optional parameters.
// https://telegra.ph/api#getViews
func GetViews(path string, opts *PageViewsOpts) (*PageViews, error) {
	var (
		u = url.Values{}
		a PageViews
	)
	u.Add("path", path)

	if opts != nil {
		u.Add("year", strconv.FormatInt(opts.Year, 10))
		u.Add("month", strconv.FormatInt(opts.Month, 10))
		u.Add("day", strconv.FormatInt(opts.Day, 10))
		u.Add("hour", strconv.FormatInt(opts.Hour, 10))
	}

	r, err := Get("getViews", u)
	if err != nil {
		return nil, err
	}
	return &a, json.Unmarshal(r, &a)
}

// Use this method to upload a file to Telegraph.
// (You can upload some specific file formats like .jpg, .jpeg, .png, .gif, etc only)
// Returns a path to the uploaded file i.e. everything that comes after https://telegra.ph/
// - filePath (type string): location of the file to upload to Telegraph.
// https://telegra.ph/upload
func UploadFile(filePath string) (string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	part, err := writer.CreateFormFile("file", filePath)
	if err != nil {
		return "", err
	}
	if _, err = io.Copy(part, file); err != nil {
		return "", err
	}
	if err = writer.Close(); err != nil {
		return "", err
	}
	request, err := http.NewRequest(http.MethodPost, "https://telegra.ph/upload", body)
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())
	var client http.Client
	httpResponse, err := client.Do(request)
	if err != nil {
		return "", err
	}
	b, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return "", err
	}
	rUpload := make([]Upload, 0)
	if err := json.Unmarshal(b, &rUpload); err != nil {
		m := map[string]string{}
		if err := json.Unmarshal(b, &m); err != nil {
			return "", err
		}
		return "", fmt.Errorf("failed to upload: %s", m["error"])
	}
	return rUpload[0].Path, nil
}

// Use this method to upload a file to Telegraph.
// (You can upload some specific file formats like .jpg, .jpeg, .png, .gif, etc only)
// Returns a path to the uploaded file i.e. everything that comes after https://telegra.ph/
// - filePath (type string): location of the file to upload to Telegraph.
// https://telegra.ph/upload
func UploadFileByBytes(content []byte) (string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormField("file")
	if err != nil {
		return "", err
	}
	if _, err = io.Copy(part, bytes.NewReader(content)); err != nil {
		return "", err
	}
	if err = writer.Close(); err != nil {
		return "", err
	}
	request, err := http.NewRequest(http.MethodPost, "https://telegra.ph/upload", body)
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())
	var client http.Client
	httpResponse, err := client.Do(request)
	if err != nil {
		return "", err
	}
	b, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return "", err
	}
	rUpload := make([]Upload, 0)
	if err := json.Unmarshal(b, &rUpload); err != nil {
		m := map[string]string{}
		if err := json.Unmarshal(b, &m); err != nil {
			return "", err
		}
		return "", fmt.Errorf("failed to upload: %s", m["error"])
	}
	return rUpload[0].Path, nil
}
