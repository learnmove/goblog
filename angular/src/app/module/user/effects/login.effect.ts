import { MessageService } from 'primeng/components/common/messageservice';
import { Observable } from 'rxjs/Observable';
import * as  Action  from './../../../interface/action';
import { UserHelper } from './../../../helper/user';
import { LoginSuccess, LoginFail } from './../action/user.action';

import { UserActionTypes } from './../action/contrast/user.contrast';
import { LoginService } from './../services/login.service';
import { Injectable } from '@angular/core';
import { Effect, Actions } from '@ngrx/effects'
import { of } from 'rxjs/observable/of';
import  'rxjs';




@Injectable()
export class LoginEffect {

@Effect()login$:Observable<any>=this.action$
.ofType(UserActionTypes.Login)
.switchMap((action:Action.CustomAction)=>{
    console.log("shit")
    return this.loginService.login(action.payload).map((payload)=>
{
    UserHelper.setUser(payload)
    return  new LoginSuccess(payload)
}).catch((err)=>{
    console.log(err)
 return Observable.of(new LoginFail(err))
}
  )
})


@Effect({dispatch:false}) 
$loginfail:Observable<Action.CustomAction>=this.action$
.ofType(UserActionTypes.LoginFail)
.do((action:Action.CustomAction)=>    
{
this.messageService.add({severity:'error', summary:'Service Message', detail:'註冊失敗'});
    
} )

constructor(private action$:Actions,private loginService:LoginService,private messageService:MessageService) { 
  }
}
