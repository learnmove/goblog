import { ImgurService } from './../services/imgur.service';
import { ImgurActionTypes } from './../action/contrast/imgur.contrast';
import { Observable } from 'rxjs/Observable';
import * as  Action  from './../../../interface/action';

import { Injectable } from '@angular/core';
import { Effect, Actions } from '@ngrx/effects'
import { of } from 'rxjs/observable/of';
import  'rxjs';
import { UploadSuccessAction,UploadFailAction } from "app/module/share/action/imgur.action";




@Injectable()
export class ImgurEffect {

@Effect()upload$:Observable<any>=this.action$
.ofType(ImgurActionTypes.Upload)
.mergeMap((action:Action.CustomAction)=>{
      return this.imgurService.uploadImage(action.payload).map((payload)=>{
         return new UploadSuccessAction(payload.data.link)
      })

}).catch((err)=>{
    console.log(err)
 return Observable.of(new UploadFailAction(err))
}
  )
constructor(private action$:Actions,private imgurService:ImgurService) { 
  }
}
