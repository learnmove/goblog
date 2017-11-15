import { rootReducer } from './../reducer/index';
import { LoginService } from './../services/user/login.service';
import {  StoreDevtoolsModule } from '@ngrx/store-devtools';
import { userRoute } from './user.route';
import { RouterModule } from '@angular/router';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LoginComponent } from './login/login.component';
import { RegisterComponent } from './register/register.component';
import { StoreModule, ReducerManager } from "@ngrx/store/store";
import { EffectsModule } from '@ngrx/effects';
import { counterReducer } from "../reducer/counter";
import { userReducer } from "../reducer/user";
import { LoginEffectService } from "../effect/user/login.effect.service";

@NgModule({
  imports: [
    CommonModule,
    RouterModule.forChild(userRoute),
    StoreModule.forFeature("usermodule",rootReducer),

 
  ],

  declarations: [LoginComponent, RegisterComponent],
})
export class UserModule { }
