import { CommentService } from './../../services/comment.service';
import { UserHelper } from './../../../../helper/user';
import { ArticleCommentPost } from './../../model/comment-post.model';
import { ArticlePost } from './../../model/article-post.model';
import { ArticleService } from './../../services/article.service';
import { ActivatedRoute } from '@angular/router';
import { Component, OnInit } from '@angular/core';
import { Location } from '@angular/common';

@Component({
  selector: 'play-article',
  templateUrl: './article.component.html',
  styleUrls: ['./article.component.scss']
})
export class ArticleComponent implements OnInit {
article:any
comment:ArticleCommentPost
content:string
  constructor(private commentService:CommentService,private location:Location,private route:ActivatedRoute,private articleService:ArticleService) {
    this.article=new ArticlePost()
    this.articleService.getArticle(this.route.snapshot.params.id).subscribe(data=>{
  let imgurl=data.data.imgurl.split(",")
      this.article={...data.data,imgurl}
  console.log(this.article)
      
    })
   }

  ngOnInit() {
  }
  goBack(): void {
        this.location.back();
    }  
  postComment(){
    this.comment=new ArticleCommentPost()
    this.comment.user_id=UserHelper.getUser().id
    this.comment.account=UserHelper.getUser().account
    this.comment.name=UserHelper.getUser().name
    this.comment.article_id=parseInt(this.route.snapshot.params.id)
    this.comment.img_url="https://scontent.fkhh1-1.fna.fbcdn.net/v/t34.0-12/24824567_1555536314536912_1079550990_n.jpg?oh=8316b8adbe5d83e16510f6d4cf389d91&oe=5A27EDF1"
    this.comment.content=this.content
    this.commentService.onPostComment(this.comment).subscribe(data=>{
      console.log(data)
      this.article.article_comment.push(this.comment)
      
    })
    this.content=""
    
  }
}
