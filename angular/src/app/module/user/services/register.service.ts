import { Observable } from 'rxjs';
import { Http,RequestOptions,Headers  } from '@angular/http';
import { Injectable } from '@angular/core';
import { Website } from './../../../config/config';


@Injectable()
export class RegisterService {

  constructor(private http:Http ) { }
  register(data):Observable<any>{
    console.log(data)
    let headers=new Headers({"content-type":"application/json"})
    let options=new RequestOptions({headers:headers})
    return this.http.post(Website+`api/v1/user/register`,data,options)
    .map((res)=>{console.log(res); return res.json()})

  }
}
