import { Website } from './../../../config/config';
import { Observable } from 'rxjs/Observable';
import { Http } from '@angular/http';
import { Injectable } from '@angular/core';

@Injectable()
export class ArticleListService {

  constructor(private http:Http) { }
  getArticleList(queryString):Observable<any>{
    return this.http.get(`${Website}api/v1/article?${queryString}`)
    .map(res=>{
      return res.json()
    })
  }

}
