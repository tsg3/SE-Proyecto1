import { Component, OnInit } from '@angular/core';
import { UserService } from '@app/_services/user-model.service';
import { TokenStorageService } from '@app/_services/token-storage.service'

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
  isLoggedIn = false;
  content: string;

  constructor(private userService: UserService, private tokenStorageService: TokenStorageService) {
    this.content = '';
  }

  // Backend
  // 
  // ngOnInit(): void {
  //   this.userService.getPublicContent().subscribe(
  //     data => {
  //       this.content = data;
  //     },
  //     err => {
  //       this.content = JSON.parse(err.error).message;
  //     }
  //   );
  // }
  ngOnInit(): void {
    this.isLoggedIn = !!this.tokenStorageService.getToken();
  }
}
