import { LoginEffect } from './module/user/effects/login.effect';
import { EffectsModule } from '@ngrx/effects';
import {StoreModule} from '@ngrx/store';
import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { StoreDevtoolsModule } from '@ngrx/store-devtools';
import { ButtonModule} from "primeng/primeng"

@NgModule({
  declarations: [
    AppComponent,
  ],
  imports: [
    ButtonModule,
    BrowserModule,
    FormsModule,
    HttpModule,
    AppRoutingModule,
    StoreModule.forRoot({}),
     EffectsModule.forRoot([]),
	StoreDevtoolsModule.instrument({
      maxAge: 5
    })
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
