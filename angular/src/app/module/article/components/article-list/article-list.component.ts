import { ArticleListService } from './../../services/article-list.service';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router, NavigationExtras } from "@angular/router";

@Component({
  selector: 'play-article-list',
  templateUrl: './article-list.component.html',
  styleUrls: ['./article-list.component.scss']
})
export class ArticleListComponent implements OnInit {
  category:any
  queryString:any
  CityValue=0
  CatValue=0
  PetValue=0
  articleList:any
  page:any=1
  pageInfo:any
  constructor(private router:Router,private route:ActivatedRoute ,private acsv :ArticleListService) { 
    this.category=this.route.snapshot.data['category'].data

  }

  ngOnInit() {
    this.route.queryParams.subscribe(v=>{
      if(v['page']!=undefined){
        this.page=v['page']
      }
      if(v['selecCat']!=undefined){
        this.CatValue=v['selecCat']
      }
      if(v['selecCity']!=undefined){
        this.CityValue=v['selecCity']
      }
      if(v['selecPet']!=undefined){
        this.PetValue=v['selecPet']
      }
    })
    this.fetchData()
  }

  setCityCategory(value){
     let qp=Object.assign({},this.route.snapshot.queryParams)
    qp['selecCity']=value
    delete qp['page']
    this.router.navigate(['article'],{queryParams:qp})
    this.CityValue=value
    this.page=1
    
    this.fetchData()
  }
  setCategory(value){
        let qp=Object.assign({},this.route.snapshot.queryParams)
    qp['selecCat']=value
    delete qp['page']
    
    this.router.navigate(['article'],{queryParams:qp})
    this.CatValue=value
    this.page=1
    
    this.fetchData()
  }
  setPetCategory(value){
    let qp=Object.assign({},this.route.snapshot.queryParams)
    qp['selecPet']=value
    delete qp['page']
    
    this.router.navigate(['article'],{queryParams:qp})
      
    this.PetValue=value
    this.page=1
    
    this.fetchData()
  }
  clickPage(page){
    this.page=page
    let qp=Object.assign({},this.route.snapshot.queryParams)
    qp['page']=page
    this.router.navigate(['article'],{queryParams:qp})
      
    this.fetchData()
    
  }
  fetchData(){
    this.queryString=`city_id=${this.CityValue}&article_category_id=${this.CatValue}&pet_category_id=${this.PetValue}&page=${this.page}`
    this.acsv.getArticleList(this.queryString).subscribe(({data})=>{
      this.articleList=data.data
      this.pageInfo=data.pageinfo
      }
    )
  }
}
