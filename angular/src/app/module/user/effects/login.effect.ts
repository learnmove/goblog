import { LoginSuccess } from './../action/user.action';
import { Action } from '@ngrx/store';

import { UserActionTypes } from './../action/contrast/user.contrast';
import { LoginService } from './../services/login.service';
import { Observable } from 'rxjs';
import { Injectable } from '@angular/core';
import { Effect, Actions } from '@ngrx/effects'
import 'rxjs/add/operator/map'


@Injectable()
export class LoginEffect {
constructor(private action$:Actions,private loginService:LoginService) { 
  }
@Effect()login$:Observable<any>=this.action$
.ofType(UserActionTypes.Login)
.switchMap((action)=>{
    console.log("effect")
    return this.loginService.login()
})
.map(payload=>new LoginSuccess(payload))
    


}
