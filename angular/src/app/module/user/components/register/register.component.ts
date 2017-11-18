import { Register } from './../../action/user.action';
import { Store } from '@ngrx/store';
import { Validators } from '@angular/forms';
import { FormGroup } from '@angular/forms/forms';
import { FormBuilder } from '@angular/forms';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'play-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.scss']
})
export class RegisterComponent implements OnInit {
  registerForm:FormGroup

  constructor( private fb:FormBuilder,private store:Store<any>, ) {
    this.createRegisterForm()
   }
   createRegisterForm(){
    this.registerForm=this.fb.group({
        name:['',Validators.required],
        account:['',Validators.compose([Validators.minLength(3),Validators.maxLength(50),Validators.required,Validators.pattern('[a-zA-Z0-9]+[a-zA-Z0-9]+')])],
        password:['',Validators.compose([Validators.minLength(6),Validators.maxLength(100),Validators.required])],
        city:['',Validators.required],
        age: ['',Validators.compose([Validators.min(15),Validators.max(70),Validators.required])],
        phone:  ['',],
        line:['',],
        address: ['',],
    })
   }
  ngOnInit() {
  }
  onRegister(){
    this.store.dispatch(new Register(this.registerForm.value))
  }
}
