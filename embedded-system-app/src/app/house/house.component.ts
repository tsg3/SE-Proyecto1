import { Component, OnDestroy, OnInit } from '@angular/core';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { TokenStorageService } from '@app/_services/token-storage.service';
import { HouseService } from '@app/_services/house.service';
import { isWhileStatement } from 'typescript';

@Component({
  selector: 'app-house',
  templateUrl: './house.component.html',
  styleUrls: ['./house.component.css']
})
export class HouseComponent implements OnInit, OnDestroy {
  isLoggedIn = false;
  public lightsOn = [false, false, false, false, false];
  public doorsOpen = [false, false, false, false];
  imgTaken = ''
  intervalID: any;

  constructor(private token: TokenStorageService, 
    private modalService: NgbModal,
    private houseService: HouseService) { }

  ngOnInit(): void {
    this.isLoggedIn = !!this.token.getToken();
    let resLights = this.houseService.getLights();
    for (let i = 0; i < 5; i++) {
      this.lightsOn[i] = resLights[i].power;
    }
    this.intervalID = setInterval(() => { 
        this.updateDoors(); 
      }, 2000);
  }

  ngOnDestroy(): void {
    if (this.intervalID) {
      clearInterval(this.intervalID);
    }
  }

  updateDoors (): void {
    let resDoors = this.houseService.getDoors();
    for (let i = 0; i < 4; i++) {
      this.doorsOpen[i] = resDoors[i].opened;
    }
  }

  // Sin backend
  switchLight (light: number): void {
    this.lightsOn[light] = !this.lightsOn[light];
    this.houseService.switchLight(light);
  }

  // Sin backend
  turnOnLights (): void {
    for (let i = 0; i < 5; i++) {
      this.lightsOn[i] = true;
    }
    this.houseService.setLights(true);
  }

  // Sin backend
  turnOffLights (): void {
    for (let i = 0; i < 5; i++) {
      this.lightsOn[i] = false;
    }
    this.houseService.setLights(false);
  }

  // Sin backend
  takePhoto (content: any): void {
    this.getPhoto();
    this.modalService.open(content, {windowClass: 'dark-modal', size: "lg"});
  }

  // Sin backend
  private getPhoto(): void {
    this.imgTaken = this.houseService.getImage();
  }

  resetPhoto(): void {
    this.imgTaken = '';
  }
}
