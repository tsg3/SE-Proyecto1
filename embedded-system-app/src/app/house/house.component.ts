import { Component, OnInit } from '@angular/core';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { TokenStorageService } from '@app/_services/token-storage.service';

@Component({
  selector: 'app-house',
  templateUrl: './house.component.html',
  styleUrls: ['./house.component.css']
})
export class HouseComponent implements OnInit {
  isLoggedIn = false;
  imgTaken = ''

  // Sin backend
  public lightsOn = [false, false, false, false, false];

  // Sin backend
  public doorsOpen = [false, false, false, false];

  constructor(private token: TokenStorageService, 
    private modalService: NgbModal) { }

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

  // Sin backend
  takePhoto (content: any): void {
    this.getPhoto();
    this.modalService.open(content, {windowClass: 'dark-modal', size: "lg"});
  }

  // Sin backend
  private getPhoto(): void {
    this.imgTaken = "../../assets/light-on.svg";
  }

  resetPhoto(): void {
    this.imgTaken = '';
  }
}
