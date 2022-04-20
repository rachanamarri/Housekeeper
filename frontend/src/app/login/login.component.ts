import { Component, OnInit } from '@angular/core';
import { AuthService } from '../ngservices/auth.service';
import { SharedService } from '../ngservices/shared-service';
import { TokenStorageService } from '../ngservices/token-storage.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
})
export class LoginComponent implements OnInit {
  form: any = {
    username: null,
    password: null
  };
  isLoggedIn = false;
  isLoginFailed = false;
  errorMessage = '';
  roles: string[] = [];  
  isProvider = true;

  get sharedLoggedInfo() {
    return this.sharedService.getOption().isLogged;
  }

    

  constructor(private authService: AuthService, private tokenStorage: TokenStorageService, private sharedService: SharedService) { 
  
  }
  

  ngOnInit() {
    // this.isLoggedIn = false;
    // if (this.tokenStorage.getToken()) {
      // this.isLoggedIn = true;
      // this.roles = this.tokenStorage.getUser().roles;
      // this.sharedService.sharedData.subscribe((isLoggedIn:boolean) => this.sharedLoggedInfo = isLoggedIn);
    // }
    // console.log('in init', this.isLoggedIn);
}


test1() {
  const { username, password } = this.form;
  console.log("entered test for login as seeker");
  this.isProvider= true;
  this.authService.login(username, password).subscribe({
    next: data => {
      this.tokenStorage.saveToken(data.accessToken);
      this.tokenStorage.saveUser(data);
      this.isLoginFailed = false;
      this.isLoggedIn = true;
      this.roles = this.tokenStorage.getUser().roles;
      //this.reloadPage();
    },
    error: err => {
      this.errorMessage = err.error.message;
      this.isLoginFailed = true;
    }
  });
  // this.ackLogIn();
  // this.sharedService.sharedData.subscribe((isLoggedIn:boolean) => this.sharedLoggedInfo = isLoggedIn);
  // console.log('done ');
  // console.log(this.sharedLoggedInfo);

  this.sharedService.setOption('isLogged', true);
}

test2() {
  console.log("entered test for login as provider");
  this.isProvider= false;
  //this.onSubmit();
  const { username, password } = this.form;
  this.authService.loginAsProvider(username, password).subscribe({
    next: data => {
      this.tokenStorage.saveToken(data.accessToken);
      this.tokenStorage.saveUser(data);
      this.isLoginFailed = false;
      this.isLoggedIn = true;
      this.roles = this.tokenStorage.getUser().roles;
      //this.reloadPage();
    },
    error: err => {
      this.errorMessage = err.error.message;
      this.isLoginFailed = true;
    }
  });
}
}
// onSubmit(): void {

//   const { username, password } = this.form;
//   if(this.isProvider==false){
//   this.authService.login(username, password).subscribe({
//     next: data => {
//       this.tokenStorage.saveToken(data.accessToken);
//       this.tokenStorage.saveUser(data);
//       this.isLoginFailed = false;
//       this.isLoggedIn = true;
//       this.roles = this.tokenStorage.getUser().roles;
//       this.reloadPage();
//     },
//     error: err => {
//       this.errorMessage = err.error.message;
//       this.isLoginFailed = true;
//     }
//   });
// }
// else{
//   this.authService.loginAsProvider(username, password).subscribe({
//     next: data => {
//       this.tokenStorage.saveToken(data.accessToken);
//       this.tokenStorage.saveUser(data);
//       this.isLoginFailed = false;
//       this.isLoggedIn = true;
//       this.roles = this.tokenStorage.getUser().roles;
//       this.reloadPage();
//     },
//     error: err => {
//       this.errorMessage = err.error.message;
//       this.isLoginFailed = true;
//     }
//   });
// }
// }
// reloadPage(): void {
//   window.location.reload();
// }
// }
