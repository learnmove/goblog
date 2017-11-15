import { RouterModule } from '@angular/router';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { PostListComponent } from './post-list/post-list.component';
import {postChild} from './post.route'
@NgModule({
  imports: [
    CommonModule,
    RouterModule.forChild(postChild)
  ],
  declarations: [PostListComponent]
})
export class PostModule { }
