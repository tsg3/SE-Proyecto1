import { Component, OnInit } from '@angular/core';
import { AuthService, Credential } from '@app/_services/auth.service';
import { TokenStorageService } from '@app/_services/token-storage.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  form: Credential;
  isLoggedIn = false;
  isLoginFailed = false;
  errorMessage = '';

  constructor(private authService: AuthService, private tokenStorage: TokenStorageService) {
    this.form = new Credential("", "");
  }

  ngOnInit(): void {
    if (this.tokenStorage.getToken()) {
      this.isLoggedIn = true;
    }
  }

  // Backend
  // 
  // onSubmit(): void {
  //   this.authService.login(this.form).subscribe(
  //     data => {
  //       this.tokenStorage.saveToken(data.accessToken);

  //       this.isLoginFailed = false;
  //       this.isLoggedIn = true;
  //       this.reloadPage();
  //     },
  //     err => {
  //       this.errorMessage = err.error.message;
  //       this.isLoginFailed = true;
  //       this.isLoggedIn = false;
  //     }
  //   );
  // }
  onSubmit(): void {
    let authentication = this.authService.login(this.form);
    if (authentication.Logged){
      this.tokenStorage.saveToken(authentication.Token);
      this.isLoginFailed = false;
      this.isLoggedIn = true;
      this.reloadPage();
    }
    else{
      this.errorMessage = authentication.Token;
      this.isLoginFailed = true;
      this.isLoggedIn = false;
    }
  }

  reloadPage(): void {
    window.location.reload();
  }

}
