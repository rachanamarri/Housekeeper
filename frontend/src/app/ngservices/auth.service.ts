import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
const AUTH_API = 'http://localhost:8080/';
const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'application/json' })
};

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  constructor(private http: HttpClient) { }

  login(name: string, password: string): Observable<any> {
    return this.http.post(AUTH_API + 'seeker_login', {
      Email: name,
      password
    }, httpOptions);
  }

  loginAsProvider(username: string, password: string): Observable<any> {
    return this.http.post(AUTH_API + 'provider_login', {
      Email: username,
      password
    }, httpOptions);
  }

  register(username: string, email: string, password: string, address: string): Observable<any> {
    return this.http.post(AUTH_API + 'seeker_registration', {
      name: username ,
      email,
      password,
      address,
    }, httpOptions);}

    registerAsProvider(username: string, email: string, password: string, address: string): Observable<any> {
      return this.http.post(AUTH_API + 'provider_registration', {
        name: username,
        email,
        password,
        address,
      }, httpOptions);
  }
}
