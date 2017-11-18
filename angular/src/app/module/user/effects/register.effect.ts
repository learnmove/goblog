import { Router } from '@angular/router';
import { MessageService } from 'primeng/components/common/messageservice';
import { RegisterService } from './../services/register.service';
import { Observable } from 'rxjs/Observable';
import * as  Action  from './../../../interface/action';
import { UserHelper } from './../../../helper/user';
import { RegisterSuccess, RegisterFail } from './../action/user.action';

import { UserActionTypes } from './../action/contrast/user.contrast';
import { Injectable } from '@angular/core';
import { Effect, Actions } from '@ngrx/effects'
import { of } from 'rxjs/observable/of';
import  'rxjs';




@Injectable()
export class RegisterEffect {

@Effect()register$:Observable<any>=this.action$
.ofType(UserActionTypes.Register)
.switchMap((action:Action.CustomAction)=>{
    console.log(action)
    return this.registerService.register(action.payload).map((payload)=>
{
     setTimeout(()=>{
     this.messageService.add({severity:'success', summary:'Service Message', detail:'註冊成功'});
 },500)
    return  new RegisterSuccess(payload)
}).catch((err)=>{
    console.log(err)
 return Observable.of(new RegisterFail(err))
}
  )
})


@Effect({dispatch:false}) 
$registerfail:Observable<Action.CustomAction>=this.action$
.ofType(UserActionTypes.RegisterFail)
.do((action:Action.CustomAction)=>{
 
 setTimeout(()=>{
     this.messageService.add({severity:'error', summary:'Service Message', detail:'註冊失敗'});
 },500)
 
})

constructor(private router: Router,private messageService: MessageService,private action$:Actions,private registerService:RegisterService) { 
  }
}
