package ViewModel

type ArticleCommentPost struct {
	Content   string `json:"content"`
	UserID    int    `json:"user_id" `
	ArticleID int    `json:"article_id" `
	ImgUrl    string `json:"img_url"`
	Name      string `json":name" `
	Account   string `json:"account" `
}
