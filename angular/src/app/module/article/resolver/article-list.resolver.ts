import { Website } from './../../../config/config';
import { Observable } from 'rxjs/Observable';
import  'rxjs';
import { Http } from '@angular/http';
import { Injectable } from '@angular/core';
import { Resolve } from "@angular/router/router";
@Injectable()
export class ArticleListResolver implements Resolve<any>{
    constructor(private http:Http){

    }
    resolve():Observable<any>{
        return this.http.get(`${Website}api/v1/category`)
        .map(res=>res.json())
    }

}