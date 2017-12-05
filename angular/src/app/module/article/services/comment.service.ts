import { Website } from './../../../config/config';
import { RequestOptions } from '@angular/http';
import { Http,Headers } from '@angular/http';
import { Injectable } from '@angular/core';

@Injectable()
export class CommentService {

  constructor(private http:Http) { }
  onPostComment(data){
    console.log(data)
   let headers=new Headers({"content-type":"application/json"})
    let options=new RequestOptions({headers:headers})
    return this.http.post(`${Website}api/v1/article/${data.article_id}/comment`,data)
    .map(res=>{
      return res.json()
    })
  }

}
