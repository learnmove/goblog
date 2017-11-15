import { Action } from "./../../action/action";
import { LoginService } from './../../services/user/login.service';

import { LOGIN } from './../../reducer/user';
import { Observable } from 'rxjs/Observable';
import { Injectable } from '@angular/core';
import { Effect, Actions } from "@ngrx/effects";

@Injectable()
export class LoginEffectService {

@Effect() 
login$:Observable<any>=this.action$
.ofType(LOGIN)
.map((action:Action)=>{
  console.log(action)
  return this.loginService.login(action.payload)
})


  constructor(private action$:Actions,private loginService:LoginService) { 

  }

}
