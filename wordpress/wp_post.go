package wordpress

import (
	"time"
)

type Post struct {
	BlogID int
	PostContent
}

type PostContent struct {
	PostType      string `xmlrpc:"post_type"`
	PostStatus    string `xmlrpc:"post_status"`
	PostTitle     string `xmlrpc:"post_title"`
	PostAuthor    int    `xmlrpc:"post_author"`
	PostExcerpt   string `xmlrpc:"post_excerpt"`
	PostContent   string `xmlrpc:"post_content"`
	PostDate      string `xmlrpc:"post_date"`
	PostFormat    string `xmlrpc:"post_format"`
	PostName      string `xmlrpc:"post_name"`
	PostPassword  string `xmlrpc:"post_password"`
	CommentStatus string `xmlrpc:"comment_status"`
	PingStatus    string `xmlrpc:"ping_status"`
	Sticky        int    `xmlrpc:"sticky"`
	PostThumbnail int    `xmlrpc:"post_thumbnail"`
	PostParent    int    `xmlrpc:"post_parent"`
	// Terms         Terms      `xmlrpc:"terms"`
	TermsNames TermsNames `xmlrpc:"terms_names"`
	Enclosure  Enclosure  `xmlrpc:"enclosure"`
}

type Terms struct {
	TermID         string `xmlrpc:"term_id"`
	TermGroup      string `xmlrpc:"term_group"`
	Taxonomy       string `xmlrpc:"taxonomy"`
	TermTaxonomyID int    `xmlrpc:"term_taxonomy_id"`
	Name           string `xmlrpc:"name"`
	Slug           string `xmlrpc:"slug"`
	Description    string `xmlrpc:"description"`
	Parent         string `xmlrpc:"parent"`
	Count          int    `xmlrpc:"count"`
}

type TermsNames struct {
	PostCategory []string `xmlrpc:"category"`
	TagsInput    []string `xmlrpc:"post_tag"`
}

type Enclosure struct {
	Url    string `xmlrpc:"url"`
	Length int    `xmlrpc:"length"`
	Type   string `xmlrpc:"type"`
}

func (p Post) GetMethord() string {
	return `wp.newPost`
}

func (p Post) GetArgs(user string, pwd string) interface{} {
	args := make([]interface{}, 4)
	args = append(args, p.BlogID, user, pwd, p.PostContent)
	return args
}

func NewPost(content string, title string, tags []string, cate []string) (p Post) {
	p.PostContent = PostContent{
		PostType:    `post`,
		PostStatus:  `publish`,
		PostTitle:   title,
		PostContent: content,
		PostDate:    time.Now().Format(`2006-01-02 15:04:05`),
		TermsNames: TermsNames{
			PostCategory: cate,
			TagsInput:    tags,
		},
	}
	return p
}

// NewSpecificPost Customize various values by yourself
func NewSpecificPost(content PostContent) (p Post) {
	p.PostContent = content
	return p
}
