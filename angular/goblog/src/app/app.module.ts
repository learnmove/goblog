import { DataService } from './services/data.service';
import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import {FormsModule} from '@angular/forms';
import { AppComponent } from './app.component';
import {HttpModule} from '@angular/http';
import {RouterModule,Routes} from '@angular/router';
const appRoutes:Routes=[
 { path:'',redirectTo:'user',pathMatch:'full'},
 {path:'user',loadChildren:'./user/user.module#UserModule'},
 {path:'post',loadChildren:'./post/post.module#PostModule'}
 
]
@NgModule({
  declarations: [
    AppComponent,
  ],
  imports: [
    FormsModule,
    BrowserModule,
    HttpModule,
    RouterModule.forRoot(appRoutes)
  ],
  providers: [DataService],
  bootstrap: [AppComponent]
})
export class AppModule { }
