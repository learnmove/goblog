import { Store } from '@ngrx/store';
import { ArticlePostService } from './../../services/article-post.service';
import { UserHelper } from './../../../../helper/user';
import { ArticlePost } from './../../model/article-post.model';
import { Validators } from '@angular/forms';
import { FormBuilder } from '@angular/forms';
import { FormGroup } from '@angular/forms/forms';
import { ArticleListService } from './../../services/article-list.service';
import { ActivatedRoute } from '@angular/router';
import { Router } from '@angular/router';
import { Component, OnInit } from '@angular/core';
import { UploadAction,ImgurClearState } from "app/module/share/action/imgur.action";

@Component({
  selector: 'play-article-post',
  templateUrl: './article-post.component.html',
  styleUrls: ['./article-post.component.scss']
})
export class ArticlePostComponent implements OnInit {
category:any
articleForm:FormGroup
articlePost:ArticlePost
constructor(private store:Store<any>,private articlePostService:ArticlePostService,private fb:FormBuilder,private router:Router,private route:ActivatedRoute ,private acsv :ArticleListService) { 
    this.category=this.route.snapshot.data['category'].data
    this.createArticleForm()
  }
  ngOnInit() {
  }
  createArticleForm(){
    this.articleForm=this.fb.group({
    title:['',Validators.required],
    city_id:['',Validators.required],
    article_category_id:['',Validators.required],
    pet_category_id:['',Validators.required],
    content:['',Validators.required],
    img_url:[''],
    })
  }
  fileUpload(e){
    for(let i =0;i<e.target.files.length;i++){
      let f =new FileReader()
      f.onload=function(e:any){
        let base64encode=e.target.result
        this.store.dispatch(new UploadAction(base64encode.replace(/data:image\/(jpeg|gif|png);base64,/,"")))
      }.bind(this)
      f.readAsDataURL(e.target.files[i])
     
      
    }

  }
  onPost(){
    
    this.articlePost=this.articleForm.value
        this.category.category.find((item)=>{
          if(item.id==this.articleForm.value.article_category_id)
          {
            this.articlePost.article_category_name=item.name
          }
          return item.id==this.articleForm.value.article_category_id
        })      
        this.category.cities.find((item)=>{
          if(item.id==this.articleForm.value.city_id)
          {
            this.articlePost.city_name=item.name
          }
          return item.id==this.articleForm.value.city_id
        })   
        this.category.pets.find((item)=>{
          if(item.id==this.articleForm.value.pet_category_id)
          {
            this.articlePost.pet_category_name=item.name
            
          }
          return item.id==this.articleForm.value.pet_category_id
        }) 
        this.articlePost.article_category_id=parseInt(this.articlePost.article_category_id)
        this.articlePost.city_id=parseInt(this.articlePost.city_id)
        this.articlePost.pet_category_id=parseInt(this.articlePost.pet_category_id)
        
        this.articlePost.user_id=UserHelper.getUser().id
        this.articlePost.name=UserHelper.getUser().name
        this.articlePost.account=UserHelper.getUser().account
         let a=""
       this.store.select(state=>state.shareModule.imgur.images).subscribe((d)=>{
        a=d.toString()
       })
       this.articlePost.img_url=a
        this.articlePostService.postArticle(this.articlePost).subscribe(data=>console.log(data))
    this.store.dispatch(new ImgurClearState({}))
        
  }
}
