import { UserActionTypes } from './contrast/user.contrast';
import { Action } from './../../../interface/action.interface';
export class Login implements Action{
    readonly type:string=UserActionTypes.Login;
    constructor(public payload:any){
        
    }
}