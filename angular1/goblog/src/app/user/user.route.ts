import { RegisterComponent } from './register/register.component';
import { LoginComponent } from './login/login.component';
import { Routes } from '@angular/router';
export const userRoute:Routes=[
    {
        path:'',redirectTo:'login',pathMatch:'full'
       
    },
    {
        path:'login',component:LoginComponent
       
    },
    {
         path:'register',component:RegisterComponent
    }
]