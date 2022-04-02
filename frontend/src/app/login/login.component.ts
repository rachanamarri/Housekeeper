import { Component, OnInit } from '@angular/core';
import { AuthService } from '../ngservices/auth.service';
import { TokenStorageService } from '../ngservices/token-storage.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  form: any = {
    name: null,
    password: null
  };
  isLoggedIn = false;
  isLoginFailed = false;
  errorMessage = '';
  roles: string[] = [];  
    

  constructor(private authService: AuthService, private tokenStorage: TokenStorageService) { }
  

  ngOnInit() {
    if (this.tokenStorage.getToken()) {
      this.isLoggedIn = true;
      this.roles = this.tokenStorage.getUser().roles;
    }
}
onSubmit(): void {
  debugger;
  const { username, password } = this.form;
  this.authService.login(username, password).subscribe({
    next: data => {
      this.tokenStorage.saveToken(data.accessToken);
      this.tokenStorage.saveUser(data);
      this.isLoginFailed = false;
      this.isLoggedIn = true;
      this.roles = username;
    },
    error: err => {
      this.errorMessage = err.error.message;
      this.isLoginFailed = true;
    }
  });
}
reloadPage(): void {
  window.location.reload();
}
}
