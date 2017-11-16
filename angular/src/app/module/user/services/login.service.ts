import { Observable } from 'rxjs';
import { Injectable } from '@angular/core';
import { Http } from "@angular/http";

@Injectable()
export class LoginService {

  constructor(private http:Http) { }
  login():Observable<any>{
    return this.http.get("https://jsonplaceholder.typicode.com/posts")
    .map((res)=>{
      console.log("service")
      return res.json()
    })
    
  }

}
