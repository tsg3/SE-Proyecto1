import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

const HOUSE_API = 'http://192.168.0.16:4200/api/';

export class ResponseElements {
  Id: number;
  State: string;
  constructor(id: number, state: string){
    this.Id = id;
    this.State = state;
  }
}

export class ResponseImage {
  Data: string;
  constructor(data: string){
    this.Data = data;
  }
}

@Injectable({
  providedIn: 'root'
})
export class HouseService {
  constructor(private http: HttpClient) { }

  switchLight (id: number, state: string): Observable<any> {
    if (state == "0") {
      return this.http.post(HOUSE_API + 'lights/turnOff/' + id.toString(), {});
    } else {
      return this.http.post(HOUSE_API + 'lights/turnOn/' + id.toString(), {});
    }
  }

  setLights (state: string): Observable<any> {
    if (state == "0") {
      return this.http.post(HOUSE_API + 'lights/turnOffAll', {});
    } else {
      return this.http.post(HOUSE_API + 'lights/turnOnAll', {});
    }
  }

  getLights (): Observable<any> {
    return this.http.get(HOUSE_API + 'lights/getAllLights');
  }

  getDoors (): Observable<any> {
    return this.http.get(HOUSE_API + 'doors/getState');
  }

  getImage (): Observable<any> { 
    return this.http.get(HOUSE_API + 'camera/take');
  }
}
