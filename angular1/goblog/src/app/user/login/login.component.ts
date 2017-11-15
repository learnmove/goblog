import { LoginAction } from './../../action/user/user.action';
import { INCREMENT, DECREMENT } from './../../reducer/counter';
import { State } from './../../store/state/state.store';
import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { ActivatedRoute } from "@angular/router";
import { Store } from "@ngrx/store/store";
import { Observable } from 'rxjs/Observable';


@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
  encapsulation: ViewEncapsulation.None
})
export class LoginComponent implements OnInit {
public id : number
public counter:Observable<number>
public user:Observable<Object>
  constructor(
    private routepar:ActivatedRoute,
    private store:Store<State>
    ) { 
      console.log(this.store)
      this.counter=this.store.select('counter')
      this.user=this.store.select('user')
      this.user.subscribe(v=>{console.log(v)})
  }

  ngOnInit() {
    // this.routepar.params.subscribe((par)=>{
    //   this.id=par["id"]
    //   console.log(this.id)
    
}
 increment(){
    this.store.dispatch({type:INCREMENT})
    this.store.dispatch(new LoginAction("shit") )
  }
  decrement(){
    this.store.dispatch({type:DECREMENT})
     this.counter.subscribe((val)=>{ console.log(val)})

  }

 
}
