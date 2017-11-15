import { StoreDevtoolsModule } from '@ngrx/store-devtools';
import { StoreModule } from '@ngrx/store';
import { RouterModule } from '@angular/router';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LoginComponent } from './components/login/login.component';
import { RegisterComponent } from './components/register/register.component';
import { userRoute } from "app/module/user/user.routes";
import { userReducer } from "app/module/user/reducers/user.reducer";

@NgModule({
  imports: [
    CommonModule,
    RouterModule.forChild(userRoute),
    StoreModule.forFeature("userModule",{user:userReducer}),
    StoreDevtoolsModule.instrument({
      maxAge: 5
    })
    
  ],
  declarations: [LoginComponent, RegisterComponent]
})
export class UserModule { }
