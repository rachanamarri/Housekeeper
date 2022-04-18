import { Component, OnInit } from '@angular/core';
import { AuthService } from '../ngservices/auth.service';
@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css'],
  
})
export class RegisterComponent implements OnInit {
  form: any = {
    name: null,
    email: null,
    password: null
  };
  isSuccessful = false;
  isSignUpFailed = false;
  errorMessage = '';
  isProvider = true;

  constructor(private authService: AuthService) { 
    
    }

  ngOnInit(): void {
  }
   

  test1() {
    console.log("entered here");
    this.isProvider= true;
    this.onSubmit();
  }

  test2() {
    console.log("entered here");
    this.isProvider= false;
    this.onSubmit();
  }


  onSubmit(): void {
    debugger;
    const { username, email, password, address } = this.form;
    


    if(this.isProvider==false) {
      // hit api
      this.authService.register(username, email, password, address).subscribe({
        next: data => {
          console.log(data);
          this.isSuccessful = true;
          this.isSignUpFailed = false;
        },
        error: err => {
          this.errorMessage = err.error.message;
          this.isSignUpFailed = true;
        }
      });
    }  
    
    else {
      this.authService.registerAsProvider(username, email, password, address).subscribe({
        next: data => {
          console.log(data);
          this.isSuccessful = true;
          this.isSignUpFailed = false;
        },
        error: err => {
          this.errorMessage = err.error.message;
          this.isSignUpFailed = true;
        }
      });
    }
    
  }
  
}
