import { Action } from '@ngrx/store';
import * as root from './../action/user.action';
import { UserActionTypes } from './../action/contrast/user.contrast';
import { INITAIL_USER_STATE } from './../store/user.store';
export function userReducer (state=INITAIL_USER_STATE,action:root.Actions){
    switch (action.type){
        case UserActionTypes.Login:
        console.log("reducer")
        return {shit:'iam'}
        case UserActionTypes.LoginSuccess:
        return {article:action.payload}
        default :
            return state

    }
}