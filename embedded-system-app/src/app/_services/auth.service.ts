import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

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

  login(credentials: any): Observable<any> {
    return this.http.post("/api/login/", credentials);
  }
}
