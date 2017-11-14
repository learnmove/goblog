import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { ActivatedRoute } from "@angular/router";

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
  encapsulation: ViewEncapsulation.None
})
export class LoginComponent implements OnInit {
public id : number
  constructor(private routepar:ActivatedRoute) { 
  }

  ngOnInit() {
    this.routepar.params.subscribe((par)=>{
      this.id=par["id"]
      console.log(this.id)
})
  }

}
