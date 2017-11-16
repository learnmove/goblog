import { Action } from '@ngrx/store';

import { UserActionTypes } from './contrast/user.contrast';
export class Login implements Action {
    readonly type:string=UserActionTypes.Login;
    constructor(public payload:any=null){
        
    }
}
export class LoginSuccess implements Action {
    readonly type:string=UserActionTypes.LoginSuccess;
    constructor(public payload:any=null){
        
    }
}
