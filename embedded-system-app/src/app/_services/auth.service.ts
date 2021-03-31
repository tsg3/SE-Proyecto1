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

export class auth_res {
  response: boolean;
  token: string;
  constructor(bool: boolean, tok: string){
    this.response = bool;
    this.token = tok;
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
  login(credentials: any): auth_res {
    if (credentials.username == 'cusadmin' 
        && credentials.password == 'password'){
      return new auth_res(true, window.btoa('cusadmin:password'));
    }
    else{
      return new auth_res(false, "Incorrect username or password!");
    }
  }
}
