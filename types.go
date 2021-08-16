package telegraph

// This object represents a Telegraph account.
type Account struct {
	// Account name, helps users with several accounts remember which they are currently using. Displayed to the user above the "Edit/Publish" button on Telegra.ph, other users don't see this name.
	ShortName string `json:"short_name"`
	// Default author name used when creating new articles.
	AuthorName string `json:"author_name"`
	// Profile link, opened when users click on the author's name below the title. Can be any link, not necessarily to a Telegram profile or channel.
	AuthorUrl string `json:"author_url"`
	// Optional. Only returned by the createAccount and revokeAccessToken method. Access token of the Telegraph account.
	AccessToken string `json:"access_token"`
	// Optional. URL to authorize a browser on telegra.ph and connect it to a Telegraph account. This URL is valid for only one use and for 5 minutes only.
	AuthUrl string `json:"auth_url"`
	// Optional. Number of pages belonging to the Telegraph account.
	PageCount int64 `json:"page_count"`
}

// Optional parameters for createAccount.
type CreateAccountOpts struct {
	// Default author name used when creating new articles.
	AuthorName string `json:"author_name"`
	// Optional. URL to authorize a browser on telegra.ph and connect it to a Telegraph account. This URL is valid for only one use and for 5 minutes only.
	AuthorUrl string `json:"author_url"`
}

// Optional parameters for editAccountInfo.
type EditAccountInfoOpts struct {
	// Account name, helps users with several accounts remember which they are currently using. Displayed to the user above the "Edit/Publish" button on Telegra.ph, other users don't see this name.
	ShortName string `json:"short_name"`
	// Default author name used when creating new articles.
	AuthorName string `json:"author_name"`
	// Profile link, opened when users click on the author's name below the title. Can be any link, not necessarily to a Telegram profile or channel.
	AuthorUrl string `json:"author_url"`
}

// This object represents a page on Telegraph.
type Page struct {
	// Path to the page.
	Path string `json:"path"`
	// URL of the page.
	Url string `json:"url"`
	// Title of the page.
	Title string `json:"title"`
	// Description of the page.
	Description string `json:"description"`
	// Optional. Name of the author, displayed below the title.
	AuthorName string `json:"author_name"`
	// Optional. Profile link, opened when users click on the author's name below the title.  Can be any link, not necessarily to a Telegram profile or channel.
	AuthorUrl string `json:"author_url"`
	// Optional. Image URL of the page.
	ImageUrl string `json:"image_url"`
	// Optional. Content of the page.
	Content []Node `json:"content"`
	// Number of page views for the page.
	Views int64 `json:"views"`
	// Optional. Only returned if access_token passed. True, if the target Telegraph account can edit the page.
	CanEdit bool `json:"can_edit"`
}

// Optional parameters for getPage and editPage.
type PageOpts struct {
	// Optional. Name of the author, displayed below the title.
	AuthorName string `json:"author_name"`
	// Optional. Profile link, opened when users click on the author's name below the title.  Can be any link, not necessarily to a Telegram profile or channel.
	AuthorUrl string `json:"author_url"`
	// If true, a content field will be returned in the Page object (see: Content format).
	ReturnContent bool `json:"return_content"`
}

// This object represents a list of Telegraph articles belonging to an account. Most recently created articles first.
type PageList struct {
	// Total number of pages belonging to the target Telegraph account.
	TotalCount int64 `json:"total_count"`
	// Requested pages of the target Telegraph account.
	Pages []Page `json:"pages"`
}

// Optional parameters for getPageList.
type PageListOpts struct {
	// - offset (type int64): Sequential number of the first page to be returned. (default = 0)
	Offset int64 `json:"offset"`
	// - limit (type int64): Limits the number of pages to be retrieved. (default = 50)
	Limit int64 `json:"limit"`
}

// This object represents the number of page views for a Telegraph article.
type PageViews struct {
	// Number of page views for the target page.
	Views int64 `json:"views"`
}

// Optional parameters for getViews.
type PageViewsOpts struct {
	// Required if month is passed. If passed, the number of page views for the requested year will be returned.
	Year int64 `json:"year"`
	// Required if day is passed. If passed, the number of page views for the requested month will be returned.
	Month int64 `json:"month"`
	// Required if hour is passed. If passed, the number of page views for the requested day will be returned.
	Day int64 `json:"day"`
	// If passed, the number of page views for the requested hour will be returned.
	Hour int64 `json:"hour"`
}

// Node is abstract object represents a DOM Node. It can be a String which represents a DOM text node or a
// NodeElement object.
type Node interface{}

// NodeElement represents a DOM element node.
type NodeElement struct {
	// Name of the DOM element. Available tags: a, aside, b, blockquote, br, code, em, figcaption, figure,
	// h3, h4, hr, i, iframe, img, li, ol, p, pre, s, strong, u, ul, video.
	Tag string `json:"tag"`

	// Attributes of the DOM element. Key of object represents name of attribute, value represents value
	// of attribute. Available attributes: href, src.
	Attrs map[string]string `json:"attrs,omitempty"`

	// List of child nodes for the DOM element.
	Children []Node `json:"children,omitempty"`
}
