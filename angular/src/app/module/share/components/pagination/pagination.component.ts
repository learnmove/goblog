import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'play-pagination',
  templateUrl: './pagination.component.html',
  styleUrls: ['./pagination.component.scss'],
})
export class PaginationComponent implements OnInit {
  @Input('pageInfo') pageInfo:any={current_page:1,last_page:10}
  @Output('clickPage')clickPage:EventEmitter<any>=new EventEmitter()
  pages:any
  start_page:any=1
  end_page:any=1
  constructor() {
    console.log("con")
    
  this.pageInfo={current_page:1,last_page:10}
    
   }

  ngOnInit() {
    console.log("init")
    
  this.pageInfo={current_page:1,last_page:10}
    
  }
  ngOnChanges(){

    console.log("change")
this.pageChange()

}
pageClick(event){
this.clickPage.emit(event)
}
pageChange(){
  console.log(this.pageInfo)
  if(this.pageInfo.current_page>=5){
     this.start_page=this.pageInfo.current_page-2
    if(this.pageInfo.last_page<= this.pageInfo.current_page+4){
      this.end_page=this.pageInfo.last_page
    }else{
      this.end_page=this.pageInfo.current_page+2
    }
  }else{
    this.start_page=1
    if(this.pageInfo.last_page>5){
    this.end_page=5
    }else{
    this.end_page=this.pageInfo.last_page
    }

  }
    let pages=new Array()
    for(let i= this.start_page;i<=this.end_page;i++){
      pages.push(i)
    }
      this.pages=pages
    


  }
}