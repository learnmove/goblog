import { ArticlePostComponent } from './components/article-post/article-post.component';
import { ArticleComponent } from './components/article/article.component';
import { ArticleListResolver } from './resolver/article-list.resolver';
import { ArticleListComponent } from './components/article-list/article-list.component';
import { Routes } from '@angular/router';
export const articleRoute:Routes=[
    {
    path:"",
    component:ArticleListComponent,
    resolve:{category:ArticleListResolver}
},
{
    path:"post",
    component:ArticlePostComponent,
    resolve:{category:ArticleListResolver}
    
},
{
    path:":id",
    component:ArticleComponent,
}
]