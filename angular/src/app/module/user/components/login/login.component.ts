import { UserHelper } from './../../../../helper/user';
import { Login } from './../../action/user.action';
import { UserState } from './../../store/user.store';
import { Store } from '@ngrx/store';
import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, FormBuilder, Validators } from "@angular/forms";
import {Message} from 'primeng/primeng'
import {MessageService} from 'primeng/components/common/messageservice';

@Component({
  selector: 'play-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  mgs:Message[]=[]
  loginForm:FormGroup
  constructor(private messageService: MessageService,private store:Store<any>,private fb:FormBuilder) { 
    this.createLoginForm()
  }
  // loginForm=new FormGroup({
  //   account:new FormControl()
  // })

  ngOnInit() {
    // console.log(this.store.select((state)=>state.userModule.user))
    
  }
  createLoginForm(){
    this.loginForm=this.fb.group({
      account:['',Validators.required],
      password:['',Validators.required]
    })
  }

  onLogin(){
  console.log(UserHelper.getUser())
  this.store.dispatch(new Login(this.loginForm.value))

  }
}
