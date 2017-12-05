import { Website } from './../../../config/config';
import { Observable } from 'rxjs/Observable';
import { Http } from '@angular/http';
import { Injectable } from '@angular/core';

@Injectable()
export class ArticleService {


  constructor(private http:Http) { }
  getArticle(id):Observable<any>{
    return this.http.get(`${Website}api/v1/article/${id}`)
    .map(res=>{
      
      return res.json()
    })
  }
}
