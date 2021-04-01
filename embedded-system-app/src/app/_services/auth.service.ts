import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';

// Backend
//
// const AUTH_API = 'http://localhost:4200/api/auth/';
//
// const httpOptions = {
//   headers: new HttpHeaders({ 'Content-Type': 'application/json' })
// };

export class Credential {
  Username: string;
  Password: string;
  constructor(username: string, password: string){
    this.Username = username;
    this.Password = password;
  }
}

export class authResponse {
  Logged: boolean;
  Token: string;
  constructor(logged: boolean, token: string){
    this.Logged = logged;
    this.Token = token;
  }
}

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  constructor(private http: HttpClient) { }

  // Backend
  // 
  // login(credentials: any): Observable<any> {
  //   return this.http.post(AUTH_API + 'signin', {
  //     username: credentials.username,
  //     password: credentials.password
  //   }, httpOptions);
  // }
  login(credentials: Credential): authResponse {
    if (credentials.Username == 'cusadmin' 
        && credentials.Password == 'password'){
      return new authResponse(true, window.btoa('cusadmin:password'));
    }
    else{
      return new authResponse(false, "Incorrect username or password!");
    }
  }
}
