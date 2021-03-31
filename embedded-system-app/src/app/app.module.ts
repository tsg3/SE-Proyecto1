import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';

import { AppComponent } from './app.component';
import { LoginComponent } from './login/login.component';
import { HomeComponent } from './home/home.component';

import { authInterceptorProviders } from './_helpers/auth.interceptor';
import { HouseComponent } from './house/house.component';

@NgModule({
    declarations: [
      AppComponent,
      LoginComponent,
      HomeComponent,
      HouseComponent
    ],
    imports: [
      BrowserModule,
      AppRoutingModule,
      FormsModule,
      HttpClientModule
    ],
    providers: [authInterceptorProviders],
    bootstrap: [AppComponent]
  })
export class AppModule { }