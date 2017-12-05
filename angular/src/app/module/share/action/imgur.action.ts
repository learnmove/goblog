import { ImgurActionTypes } from './contrast/imgur.contrast';
import { CustomAction } from './../../../interface/action';
export class  UploadAction implements CustomAction{
    readonly type:string=ImgurActionTypes.Upload;
    constructor(public payload:any){

    }
}
export class  UploadSuccessAction implements CustomAction{
    readonly type:string=ImgurActionTypes.UploadSuccess;
    constructor(public payload:any){
        
    }
}

export class  UploadFailAction implements CustomAction{
    readonly type:string=ImgurActionTypes.UploadFail;
    constructor(public payload:any){
        
    }
}
export class  ImgurClearState implements CustomAction{
    readonly type:string=ImgurActionTypes.ClearState;
    constructor(public payload:any){
        
    }
}
export type Actions=
UploadAction|UploadSuccessAction|UploadFailAction|ImgurClearState