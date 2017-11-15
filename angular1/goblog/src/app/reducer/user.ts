import { Action } from "./../action/action";

export const LOGIN="LOGIN"
export function userReducer(state={user:{name:"ilove"}},action:Action,){
    switch(action.type){
        case LOGIN:
        return {...state,dd:action.payload}
        default:
        return state
    }
}