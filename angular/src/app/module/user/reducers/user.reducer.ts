import { UserActionTypes } from './../action/contrast/user.contrast';
import { Action } from './../../../interface/action.interface';
import { INITAIL_USER_STATE } from './../store/user.store';
export function userReducer (state=INITAIL_USER_STATE,action:Action){
    switch (action.type){
        case UserActionTypes.Login:
        return {shit:'iam'}
        default :
            return state
    }
}