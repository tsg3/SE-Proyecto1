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

  async ngOnInit() {
    this.isLoggedIn = !!this.token.getToken();
    await this.houseService.getLights().then(
      data => {
        for (let i = 0; i < 5; i++) {
          this.lightsOn[i] = data[i].State;
        }
      }, err => {
        console.log(err);
      }
    );
    this.endInterval();
    this.startInterval();
  }

  ngOnDestroy(): void {
    this.endInterval();
  }

  startInterval(): void {
    this.intervalID = setInterval(() => { 
      this.updateDoors(); 
    }, 2000);
  }

  endInterval(): void {
    if (this.intervalID) {
      clearInterval(this.intervalID);
    }
  }

  async updateDoors () {
    await this.houseService.getDoors().then(
      data => {
        for (let i = 0; i < 4; i++) {
          this.doorsOpen[i] = data[i].State;
        }
      }, err => {
        console.log(err);
      }
    );
  }

  async switchLight (id: number) {
    this.endInterval();
    let finalState = this.lightsOn[id] == "1" ? "0" : "1";
    await this.houseService.switchLight(id, finalState).then(
      data => {
        this.lightsOn[data.Id] = data.State;
      }, err => {
        console.log(err);
      }
    );
    this.startInterval();
  }

  async turnOnLights () {
    this.endInterval();
    await this.houseService.setLights("1").then(
      data => {
        if (data.Id = -1 && data.State == "ALLON") {
          for (let i = 0; i < 5; i++) {
            this.lightsOn[i] = "1";
          }
        }
      }, err => {
        console.log(err);
      }
    );
    this.startInterval();
  }

  async turnOffLights () {
    this.endInterval();
    await this.houseService.setLights("0").then(
      data => {
        if (data.Id = -1 && data.State == "ALLOFF") {
          for (let i = 0; i < 5; i++) {
            this.lightsOn[i] = "0";
          }
        }
      }, err => {
        console.log(err);
      }
    );
    this.startInterval();
  }

  takePhoto (content: any): void {
    this.endInterval();
    this.getPhoto(content);
    this.startInterval();
  }

  private async getPhoto(content: any) {
    await this.houseService.getImage().then(
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
