import { Login } from './../../action/user.action';
import { UserState } from './../../store/user.store';
import { Store } from '@ngrx/store';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'play-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  constructor(private store:Store<any>) { 
  }

  ngOnInit() {
    console.log(this.store.select((state)=>state.userModule.user))
    
  }
  test(){
    this.store.dispatch(new Login())
  }
}
