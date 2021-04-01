import { Injectable } from '@angular/core';

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
  constructor() { }

  switchLight (id: number): void {
    // Se envia al backend
  }

  setLights (state: string): void {
    // Se envia al backend
  }

  getLights (): Array<ResponseElements> {
    // Se recibe del backend 'res'
    let res = [{Id: 0, State: "0"},
              {Id: 1, State: "0"},
              {Id: 2, State: "0"},
              {Id: 3, State: "0"},
              {Id: 4, State: "0"}];
    return res;
  }

  getDoors (): Array<ResponseElements> {
    // Se recibe del backend 'res'
    let res = [{Id: 0, State: "0"},
      {Id: 1, State: "0"},
      {Id: 2, State: "0"},
      {Id: 3, State: "0"}];
    return res;
  }

  getImage (): ResponseImage { 
    // Se recibe del backend 'res'
    let res = {Data: "../../assets/light-on.svg"};
    return res;
  }
}
