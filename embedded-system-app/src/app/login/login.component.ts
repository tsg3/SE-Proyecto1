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

  onSubmit(): void {
    this.authService.login(this.form).subscribe(
      data => {
        this.tokenStorage.saveToken(data.Token);
        this.isLoginFailed = false;
        this.isLoggedIn = true;
        this.reloadPage();
      },
      err => {
        console.log(err);
        this.errorMessage = err.error;
        this.isLoginFailed = true;
        this.isLoggedIn = false;
      }
    );
  }

  reloadPage(): void {
    window.location.reload();
  }

}
