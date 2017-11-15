import { Http } from '@angular/http';
import { Injectable } from '@angular/core';

@Injectable()
export class LoginService {

  constructor(private http:Http) { }
  login(payload:any){
    return this.http.get(`https://jsonplaceholder.typicode.com/posts`).map((res)=>
   {
     console.log(res.json())
    res.json() 
  }
    )
    
  }
}
