import { Action } from '@ngrx/store';
import * as root from './../action/user.action';
import { UserActionTypes } from './../action/contrast/user.contrast';
import { INITAIL_USER_STATE } from './../store/user.store';
export function userReducer (state=INITAIL_USER_STATE,action:root.Actions){
    switch (action.type){
        case UserActionTypes.Login:
        console.log("login................")
        return state
        case UserActionTypes.LoginSuccess:
            return {...state,userInfo:action.payload}
        case UserActionTypes.LoginFail:
            return state
        case UserActionTypes.Register:
        console.log("register............")
        return state
        case UserActionTypes.RegisterSuccess:
        console.log("registerSuccess")
        return state 
        case UserActionTypes.RegisterFail:
        console.log("registerFail")
        return state

        
        default :
            return state

    }
}