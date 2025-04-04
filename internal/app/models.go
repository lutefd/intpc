package app

type InsomniaExport struct {
	Type         string               `yaml:"type" json:"type"`
	Name         string               `yaml:"name" json:"name"`
	Meta         InsomniaMeta         `yaml:"meta" json:"meta"`
	Collection   []InsomniaResource   `yaml:"collection" json:"collection"`
	CookieJar    InsomniaCookieJar    `yaml:"cookieJar,omitempty" json:"cookieJar,omitempty"`
	Environments InsomniaEnvironments `yaml:"environments,omitempty" json:"environments,omitempty"`
}

type InsomniaMeta struct {
	ID        string `yaml:"id" json:"id"`
	Created   int64  `yaml:"created" json:"created"`
	Modified  int64  `yaml:"modified" json:"modified"`
	IsPrivate bool   `yaml:"isPrivate,omitempty" json:"isPrivate,omitempty"`
	SortKey   int64  `yaml:"sortKey,omitempty" json:"sortKey,omitempty"`
}

type InsomniaResource struct {
	Name       string               `yaml:"name" json:"name"`
	Meta       InsomniaMeta         `yaml:"meta" json:"meta"`
	URL        string               `yaml:"url,omitempty" json:"url,omitempty"`
	Method     string               `yaml:"method,omitempty" json:"method,omitempty"`
	Body       *InsomniaRequestBody `yaml:"body,omitempty" json:"body,omitempty"`
	Parameters []InsomniaParameter  `yaml:"parameters,omitempty" json:"parameters,omitempty"`
	Headers    []InsomniaHeader     `yaml:"headers,omitempty" json:"headers,omitempty"`
	Settings   InsomniaSettings     `yaml:"settings,omitempty" json:"settings,omitempty"`
	Children   []InsomniaResource   `yaml:"children,omitempty" json:"children,omitempty"`
}

type InsomniaRequestBody struct {
	Text     string `yaml:"text,omitempty" json:"text,omitempty"`
	MimeType string `yaml:"mimeType,omitempty" json:"mimeType,omitempty"`
}

type InsomniaHeader struct {
	Name     string `yaml:"name" json:"name"`
	Value    string `yaml:"value" json:"value"`
	Disabled bool   `yaml:"disabled,omitempty" json:"disabled,omitempty"`
}

type InsomniaParameter struct {
	Name     string `yaml:"name" json:"name"`
	Value    string `yaml:"value" json:"value"`
	Disabled bool   `yaml:"disabled,omitempty" json:"disabled,omitempty"`
}

type InsomniaSettings struct {
	RenderRequestBody bool            `yaml:"renderRequestBody,omitempty" json:"renderRequestBody,omitempty"`
	EncodeURL         bool            `yaml:"encodeUrl,omitempty" json:"encodeUrl,omitempty"`
	FollowRedirects   string          `yaml:"followRedirects,omitempty" json:"followRedirects,omitempty"`
	Cookies           InsomniaCookies `yaml:"cookies,omitempty" json:"cookies,omitempty"`
	RebuildPath       bool            `yaml:"rebuildPath,omitempty" json:"rebuildPath,omitempty"`
}

type InsomniaCookies struct {
	Send  bool `yaml:"send" json:"send"`
	Store bool `yaml:"store" json:"store"`
}

type InsomniaCookieJar struct {
	Name string       `yaml:"name" json:"name"`
	Meta InsomniaMeta `yaml:"meta" json:"meta"`
}

type InsomniaEnvironments struct {
	Name      string       `yaml:"name" json:"name"`
	Meta      InsomniaMeta `yaml:"meta" json:"meta"`
	IsPrivate bool         `yaml:"isPrivate,omitempty" json:"isPrivate,omitempty"`
}

type PostmanCollection struct {
	Info      PostmanInfo       `json:"info"`
	Items     []PostmanItem     `json:"item"`
	Variables []PostmanVariable `json:"variable,omitempty"`
	Auth      *PostmanAuth      `json:"auth,omitempty"`
	Events    []PostmanEvent    `json:"event,omitempty"`
}

type PostmanInfo struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Schema      string `json:"schema"`
	PostmanID   string `json:"_postman_id"`
}

type PostmanItem struct {
	Name        string            `json:"name"`
	Description string            `json:"description,omitempty"`
	Items       []PostmanItem     `json:"item,omitempty"`
	Request     *PostmanRequest   `json:"request,omitempty"`
	Response    []PostmanResponse `json:"response,omitempty"`
}

type PostmanRequest struct {
	Method      string          `json:"method"`
	URL         PostmanURL      `json:"url"`
	Headers     []PostmanHeader `json:"header,omitempty"`
	Body        *PostmanBody    `json:"body,omitempty"`
	Description string          `json:"description,omitempty"`
	Auth        *PostmanAuth    `json:"auth,omitempty"`
}

type PostmanURL struct {
	Raw      string         `json:"raw"`
	Protocol string         `json:"protocol,omitempty"`
	Host     []string       `json:"host,omitempty"`
	Path     []string       `json:"path,omitempty"`
	Query    []PostmanQuery `json:"query,omitempty"`
}

type PostmanQuery struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	Disabled bool   `json:"disabled,omitempty"`
}

type PostmanHeader struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	Disabled bool   `json:"disabled,omitempty"`
}

type PostmanBody struct {
	Mode       string                 `json:"mode,omitempty"`
	Raw        string                 `json:"raw,omitempty"`
	FormData   []PostmanFormData      `json:"formdata,omitempty"`
	URLEncoded []PostmanURLEncoded    `json:"urlencoded,omitempty"`
	Options    map[string]interface{} `json:"options,omitempty"`
}

type PostmanFormData struct {
	Key      string `json:"key"`
	Value    string `json:"value,omitempty"`
	Type     string `json:"type"`
	Src      string `json:"src,omitempty"`
	Disabled bool   `json:"disabled,omitempty"`
}

type PostmanURLEncoded struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	Disabled bool   `json:"disabled,omitempty"`
}

type PostmanVariable struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
}

type PostmanAuth struct {
	Type   string             `json:"type"`
	Basic  []PostmanAuthParam `json:"basic,omitempty"`
	Bearer []PostmanAuthParam `json:"bearer,omitempty"`
	Digest []PostmanAuthParam `json:"digest,omitempty"`
	OAuth1 []PostmanAuthParam `json:"oauth1,omitempty"`
	OAuth2 []PostmanAuthParam `json:"oauth2,omitempty"`
}

type PostmanAuthParam struct {
	Key   string `json:"key"`
	Value string `json:"value,omitempty"`
	Type  string `json:"type,omitempty"`
}

type PostmanResponse struct {
	Name            string          `json:"name"`
	OriginalRequest PostmanRequest  `json:"originalRequest"`
	Status          string          `json:"status"`
	Code            int             `json:"code"`
	Headers         []PostmanHeader `json:"header"`
	Body            string          `json:"body,omitempty"`
}

type PostmanEvent struct {
	Listen string        `json:"listen"`
	Script PostmanScript `json:"script"`
}

type PostmanScript struct {
	Type string   `json:"type"`
	Exec []string `json:"exec"`
}
