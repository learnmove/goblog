import { ImgurActionTypes } from './../action/contrast/imgur.contrast';
import { INITIAL_IMGUR_STATE } from './../store/imgur.store';
import * as root from './../action/imgur.action'
export function imgurReducer(state=INITIAL_IMGUR_STATE,action:root.Actions){
    switch (action.type){
        case ImgurActionTypes.Upload:
        console.log("upload....")
             return state
        case ImgurActionTypes.UploadSuccess:
              
              return {...state,images:[...state.images,action.payload]}
        case ImgurActionTypes.UploadFail:
            console.log("upload.  UploadFail")
            return state
        case ImgurActionTypes.ClearState:
            return {...state,images:[]}
        default:
            return state
    }


}