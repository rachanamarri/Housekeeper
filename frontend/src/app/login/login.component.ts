import { Component, OnInit } from '@angular/core';
import { FormControl, Validators, FormGroup, FormBuilder } from '@angular/forms';


@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
    loginForm: FormGroup | undefined;
    submitted: boolean = false;
    

  constructor(private formBuilder: FormBuilder) { }
  

  ngOnInit() {
    this.loginForm = this.formBuilder.group({
      username: ['', Validators.required],
      password: ['', Validators.required]

  });
  
  
}
onSubmit() {
  this.submitted = true;

  alert ("Logged in")

}
}
