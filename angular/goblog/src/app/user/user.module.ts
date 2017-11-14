import { userRoute } from './user.route';
import { RouterModule } from '@angular/router';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LoginComponent } from './login/login.component';
import { RegisterComponent } from './register/register.component';

@NgModule({
  imports: [
    CommonModule,
    RouterModule.forChild(userRoute)
  ],
  declarations: [LoginComponent, RegisterComponent]
})
export class UserModule { }
