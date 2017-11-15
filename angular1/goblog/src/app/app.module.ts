import { LoginService } from './services/user/login.service';
import { LoginEffectService } from './effect/user/login.effect.service';
import { EffectsModule } from '@ngrx/effects';
import { rootReducer, reducerToken } from './reducer/index';
import { DataService } from './services/data.service';
import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import {FormsModule} from '@angular/forms';
import { AppComponent } from './app.component';
import {HttpModule} from '@angular/http';
import {RouterModule,Routes} from '@angular/router';
import {StoreModule,Store} from '@ngrx/store';
import {StoreDevtoolsModule} from "@ngrx/store-devtools";
import { counterReducer } from "./reducer/counter";
const appRoutes: Routes = [
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
    RouterModule.forRoot(appRoutes),
    StoreModule.forRoot(rootReducer),
  ],
  providers: [DataService],
  bootstrap: [AppComponent]
})
export class AppModule { }
