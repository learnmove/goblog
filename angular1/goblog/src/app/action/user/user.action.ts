import { LOGIN } from "../../reducer/user";
import { Action } from "../action";
export class LoginAction implements Action {
    readonly type:string=LOGIN
    constructor(public payload:any=null){}
}