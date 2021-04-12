import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class HouseService {
  constructor(private http: HttpClient) { }

  async switchLight (id: number, state: string): Promise<any> {
    if (state == "0") {
      return await this.http.post('/api/lights/turnOff/' + id.toString(), {}).toPromise();
    } else {
      return await this.http.post('/api/lights/turnOn/' + id.toString(), {}).toPromise();
    }
  }

  async setLights (state: string): Promise<any> {
    if (state == "0") {
      return await this.http.post('/api/lights/turnOffAll', {}).toPromise();
    } else {
      return await this.http.post('/api/lights/turnOnAll', {}).toPromise();
    }
  }

  async getLights (): Promise<any> {
    return await this.http.get('/api/lights/getAllLights').toPromise();
  }

  async getDoors (): Promise<any> {
    return await this.http.get('/api/doors/getState').toPromise();
  }

  async getImage (): Promise<any> { 
    return await this.http.get('/api/camera/take').toPromise();
  }
}
