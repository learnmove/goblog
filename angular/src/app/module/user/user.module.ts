import { RegisterService } from './services/register.service';
import { LoginService } from './services/login.service';
import { LoginEffect } from './effects/login.effect';
import { RegisterEffect } from './effects/register.effect';
import { StoreDevtoolsModule } from '@ngrx/store-devtools';
import { StoreModule } from '@ngrx/store';
import { RouterModule } from '@angular/router';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LoginComponent } from './components/login/login.component';
import { RegisterComponent } from './components/register/register.component';
import { userRoute } from "app/module/user/user.routes";
import { userReducer } from "app/module/user/reducers/user.reducer";
import { EffectsModule } from "@ngrx/effects";
import { ButtonModule} from "primeng/primeng"
import { ReactiveFormsModule } from "@angular/forms";
import {MessagesModule,GrowlModule} from 'primeng/primeng';
import { MessageService } from "primeng/components/common/messageservice";

@NgModule({
  imports: [
    ButtonModule,
    GrowlModule,
    MessagesModule,
    CommonModule,
    
    ReactiveFormsModule,
    RouterModule.forChild(userRoute),
    StoreModule.forFeature("userModule",{user:userReducer}),
    EffectsModule.forFeature([LoginEffect,RegisterEffect]),
    StoreDevtoolsModule.instrument({
      maxAge: 5
    }),
    
  ],
  declarations: [LoginComponent, RegisterComponent],
  providers:[LoginService,MessageService,RegisterService]
})
export class UserModule { }
