import { Component, OnInit } from '@angular/core';
import { TokenStorageService } from '@app/_services/token-storage.service';

@Component({
  selector: 'app-house',
  templateUrl: './house.component.html',
  styleUrls: ['./house.component.css']
})
export class HouseComponent implements OnInit {
  isLoggedIn = false;
  // Sin backend
  public lightsOn = [false, false, false, false, false];

  constructor(private token: TokenStorageService) { }

  ngOnInit(): void {
    this.isLoggedIn = !!this.token.getToken();
  }

  // Sin backend
  switchLight (light: number): void {
    this.lightsOn[light] = !this.lightsOn[light];
  }

  // Sin backend
  turnOnLights (): void {
    for (let light = 0; light < 5; light++) {
      this.lightsOn[light] = true; 
    }
  }

  // Sin backend
  turnOffLights (): void {
    for (let light = 0; light < 5; light++) {
      this.lightsOn[light] = false; 
    }
  }
}
