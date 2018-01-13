# Introduction

An implementation of the standard WordPress API methods is provided,Use the [https://github.com/kolo/xmlrpc](https://github.com/kolo/xmlrpc) library as client.

# Todo List
- [x] [wp.newPost](#wp.newPost)
- [ ] wp.getPost
- [ ] wp.getPosts
- [ ] wp.editPost
- [ ] wp.deletePost

# Usage
- ### wp.newPost
    ```go
    c, err := xmlrpc.NewClient(`https://example.com/xmlrpc.php`, xmlrpc.UserInfo{
		`your username`,
		`your password`,
	})
	if err != nil {
		log.Fatalln(err)
	}
	p := wordpress.NewPost(`content`, `title`, []string{`tag1`, `tag2`}, []string{`cate1`, `cate2`})
	blogID, err := c.Call(p)
	if err != nil {
		log.Println(err)
	}
	log.Println(blogID)
    ```
