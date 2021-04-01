import { Component, OnDestroy, OnInit } from '@angular/core';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { TokenStorageService } from '@app/_services/token-storage.service';
import { HouseService } from '@app/_services/house.service';

@Component({
  selector: 'app-house',
  templateUrl: './house.component.html',
  styleUrls: ['./house.component.css']
})
export class HouseComponent implements OnInit, OnDestroy {
  isLoggedIn = false;
  public lightsOn = ["0", "0", "0", "0", "0"];
  public doorsOpen = ["0", "0", "0", "0"];
  imgBase64 = ""
  intervalID: any;

  constructor(private token: TokenStorageService, 
    private modalService: NgbModal,
    private houseService: HouseService) { }

  ngOnInit(): void {
    this.isLoggedIn = !!this.token.getToken();
    let resLights = this.houseService.getLights();
    for (let i = 0; i < 5; i++) {
      this.lightsOn[i] = resLights[i].State;
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
      this.doorsOpen[i] = resDoors[i].State;
    }
  }

  // Sin backend
  switchLight (id: number): void {
    this.lightsOn[id] = this.lightsOn[id] == "1" ? "0" : "1";
    this.houseService.switchLight(id);
  }

  // Sin backend
  turnOnLights (): void {
    for (let i = 0; i < 5; i++) {
      this.lightsOn[i] = "1";
    }
    this.houseService.setLights("1");
  }

  // Sin backend
  turnOffLights (): void {
    for (let i = 0; i < 5; i++) {
      this.lightsOn[i] = "0";
    }
    this.houseService.setLights("0");
  }

  takePhoto (content: any): void {
    this.getPhoto(content);
  }

  private getPhoto(content: any): void {
    this.houseService.getImage().subscribe(
      data => {
        this.imgBase64 = data.Data;
        this.modalService.open(content, {windowClass: 'dark-modal', size: "md"});
      }, err => {
        console.log(err.error);
      }
    );
  }

  resetPhoto(): void {
    this.imgBase64 = '';
  }
}
