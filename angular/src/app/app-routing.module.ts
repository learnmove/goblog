import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

const routes: Routes = [
  { path: '', redirectTo: 'blog', pathMatch: 'full' },
  {
    path: 'blog',
    loadChildren: './blog-page/blog-page.module#BlogPageModule'
  },
  {
    path: 'todos',
    loadChildren: './todos-page/todos-page.module#TodosPageModule'
  },
  {
    path:'user',
    loadChildren:'./module/user/user.module#UserModule'
  },
  {
    path:'article',
    loadChildren:'./module/article/article.module#ArticleModule'
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
