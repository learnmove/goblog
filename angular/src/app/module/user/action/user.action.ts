import { CustomAction } from './../../../interface/action';

import { UserActionTypes } from './contrast/user.contrast';
export class Login implements CustomAction {
    readonly type:string=UserActionTypes.Login;
    constructor(public payload:any){
        
    }
}
export class LoginSuccess implements CustomAction {
    readonly type:string=UserActionTypes.LoginSuccess;
    constructor(public payload:any){
        
    }
}
export class LoginFail implements CustomAction{
    readonly type:string=UserActionTypes.LoginFail;
    constructor(public payload:any){
        
    }
}
export class Register implements CustomAction{
    readonly type:string=UserActionTypes.Register;
    constructor(public payload:any){
        
    }
}
export class RegisterSuccess implements CustomAction{
    readonly type:string=UserActionTypes.RegisterSuccess;
    constructor(public payload:any){
        
    }
}
export class RegisterFail implements CustomAction{
    readonly type:string=UserActionTypes.RegisterFail;
    constructor(public payload:any){
        
    }
}
export type Actions=
Login|LoginSuccess|LoginFail|Register|RegisterSuccess|RegisterFail