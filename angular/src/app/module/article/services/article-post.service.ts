import { Website } from './../../../config/config';
import { RequestOptions,Headers,Http } from '@angular/http';
import { Observable } from 'rxjs/Observable';
import { Injectable } from '@angular/core';

@Injectable()
export class ArticlePostService {

  constructor(private http:Http) { }
  postArticle(body):Observable<any>{
    console.log(body)
    let headers=new Headers({"content-type":"application/json"})
    let options=new RequestOptions({headers:headers})
    return this.http.post(`${Website}api/v1/article`,body)
    .map(res=>{
      return res.json()
    })
  }

}
