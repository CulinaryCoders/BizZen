import { Component } from '@angular/core';
import {Router} from '@angular/router';


@Component({
  selector: 'app-landing',
  templateUrl: './landing.component.html',
  styleUrls: ['./landing.component.scss']
})
export class LandingComponent {
  constructor(private router:Router) {}
  routeToRegister() {
    this.router.navigate(['/register']);
  }
  routeToLogin() {
    this.router.navigate(['/login']);
  }
}
