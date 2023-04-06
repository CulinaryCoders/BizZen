import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { CommonModule } from '@angular/common';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import { CalendarModule, DateAdapter } from 'angular-calendar';
import { adapterFactory } from 'angular-calendar/date-adapters/date-fns';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginComponent } from './login/login.component';
import { ProfileComponent } from './profile/profile.component';
import { RegisterComponent } from './register/register.component';
import { LandingComponent } from './landing/landing.component';
import { OnboardingComponent } from './onboarding/onboarding.component';
import { BusinessOnboardingComponent } from './business-onboarding/business-onboarding.component';
import { CreateServiceComponent } from './create-service/create-service.component';
import { FindClassesComponent } from './find-classes/find-classes.component';
import { ScrollingModule } from '@angular/cdk/scrolling';
import { ServicePageComponent } from './service-page/service-page.component';
import { BusinesesDashboardComponent } from './busineses-dashboard/busineses-dashboard.component';
import { ServiceCalendarComponent } from './service-calendar/service-calendar.component';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    ProfileComponent,
    RegisterComponent,
    LandingComponent,
    OnboardingComponent,
    BusinessOnboardingComponent,
    CreateServiceComponent,
    FindClassesComponent,
    ServicePageComponent,
    BusinesesDashboardComponent,
    ServiceCalendarComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    AppRoutingModule,
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    ScrollingModule,
    CalendarModule.forRoot({
      provide: DateAdapter,
      useFactory: adapterFactory,
    })
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
