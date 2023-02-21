import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './login/login.component';
import { ProfileComponent } from './profile/profile.component';
import {RegisterComponent} from "./register/register.component";
import {LandingComponent} from "./landing/landing.component";
import {OnboardingComponent} from "./onboarding/onboarding.component";

const routes: Routes = [
  //more specific routes should be above less specific routes
  { path: '', component: LandingComponent },
  { path: 'login', component: LoginComponent },
  { path: 'profile', component: ProfileComponent },
  { path: 'register', component: RegisterComponent },
  { path: 'onboarding', component: OnboardingComponent },
  { path: '',   redirectTo: '/home', pathMatch: 'full' }, // redirect to `first-component`
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
