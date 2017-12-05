import { Website } from './../../../config/config';
import { Observable } from 'rxjs';
import { Injectable } from '@angular/core';
import { Http, RequestOptions,Headers } from "@angular/http";

@Injectable()
export class ImgurService {

  constructor(private http:Http) { }
  uploadImage(data):Observable<any>{
    let headers=new Headers({"content-type":"application/json","Authorization":"Client-ID 081f020139e558d"})
    let options=new RequestOptions({headers:headers})
    return this.http.post("https://api.imgur.com/3/image",{"image":data},options)
    .map((res)=>{
      console.log(res.json())
      return res.json()
    })
    
  }

}
