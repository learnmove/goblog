import { ImgurEffect } from './effects/imgur.effect';
import { EffectsModule } from '@ngrx/effects/effects';
import { ImgurService } from './services/imgur.service';
import { StoreDevtoolsModule } from '@ngrx/store-devtools';
import { StoreModule } from '@ngrx/store';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { PaginationComponent } from './components/pagination/pagination.component';
import { imgurReducer } from "app/module/share/reducers/imgur.reducer";

@NgModule({
  imports: [
    CommonModule,
    StoreModule.forFeature("shareModule",{imgur:imgurReducer}),
    EffectsModule.forRoot([ImgurEffect]),
    StoreDevtoolsModule.instrument(
      {maxAge:5}
    ),

  ],
  
  declarations: [PaginationComponent],
  providers:[ImgurService],
  exports:[PaginationComponent]
  
})
export class ShareModule { }
