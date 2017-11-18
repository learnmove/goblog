import { Website } from './../../../config/config';
import { Observable } from 'rxjs';
import { Injectable } from '@angular/core';
import { Http, RequestOptions,Headers } from "@angular/http";

@Injectable()
export class LoginService {

  constructor(private http:Http) { }
  login(data):Observable<any>{
    console.log(data)
    let headers=new Headers({"content-type":"application/json"})
    let options=new RequestOptions({headers:headers})
    return this.http.post(Website+"api/v1/user/login",data,options)
    .map((res)=>{
      console.log(res.json())
      return res.json()
    })
    
  }

}
