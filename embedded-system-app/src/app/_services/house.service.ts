import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

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
      return this.http.post('/api/lights/turnOff/' + id.toString(), {});
    } else {
      return this.http.post('/api/lights/turnOn/' + id.toString(), {});
    }
  }

  setLights (state: string): Observable<any> {
    if (state == "0") {
      return this.http.post('/api/lights/turnOffAll', {});
    } else {
      return this.http.post('/api/lights/turnOnAll', {});
    }
  }

  getLights (): Observable<any> {
    return this.http.get('/api/lights/getAllLights');
  }

  getDoors (): Observable<any> {
    return this.http.get('/api/doors/getState');
  }

  getImage (): Observable<any> { 
    return this.http.get('/api/camera/take');
  }
}
