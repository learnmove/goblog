import { CommentService } from './services/comment.service';
import { imgurReducer } from 'app/module/share/reducers/imgur.reducer';
import { StoreModule } from '@ngrx/store';
import { ImgurEffect } from './../share/effects/imgur.effect';
import { EffectsModule } from '@ngrx/effects/effects';
import { ArticlePostService } from './services/article-post.service';
import { ReactiveFormsModule,FormsModule } from '@angular/forms';
import { ArticleService } from './services/article.service';
import { ShareModule } from './../share/share.module';
import { ArticleListService } from './services/article-list.service';
import { ArticleListResolver } from './resolver/article-list.resolver';
import { RouterModule } from '@angular/router';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ArticleListComponent } from './components/article-list/article-list.component';
import { articleRoute } from "app/module/article/article.route";
import { ArticleComponent } from './components/article/article.component';
import { ArticlePostComponent } from './components/article-post/article-post.component';


@NgModule({
  imports: [

    CommonModule,
    RouterModule.forChild(articleRoute),
    FormsModule ,
    ShareModule,
    ReactiveFormsModule,

  ],
  providers:[CommentService,ArticlePostService,ArticleListResolver,ArticleListService,ArticleService],
  declarations: [ArticleListComponent, ArticleComponent, ArticlePostComponent]
})
export class ArticleModule { }
