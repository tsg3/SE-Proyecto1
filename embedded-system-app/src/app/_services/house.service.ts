import { Injectable } from '@angular/core';

export class Light {
  light: number;
  power: boolean;
  constructor(num: number, bool: boolean){
    this.light = num;
    this.power = bool;
  }
}

export class Door {
  door: number;
  opened: boolean;
  constructor(num: number, bool: boolean){
    this.door = num;
    this.opened = bool;
  }
}

@Injectable({
  providedIn: 'root'
})
export class HouseService {
  constructor() { }

  switchLight (light: number): void {
    // Se envia al backend
  }

  setLights (power: boolean): void {
    // Se envia al backend
  }

  getLights (): Array<Light> {
    // Se recibe del backend 'res'
    let res = [{light: 0, power: false},
              {light: 1, power: false},
              {light: 2, power: false},
              {light: 3, power: false},
              {light: 4, power: false}];
    return res;
  }

  getDoors (): Array<Door> {
    // Se recibe del backend 'res'
    let res = [{door: 0, opened: false},
      {door: 1, opened: false},
      {door: 2, opened: false},
      {door: 3, opened: false}];
    return res;
  }

  getImage (): string { 
    // Se recibe del backend 'res'
    let res = "../../assets/light-on.svg"
    return res;
  }
}
