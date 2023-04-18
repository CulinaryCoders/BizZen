import { Component } from '@angular/core';
import {ActivatedRoute, Router} from "@angular/router";

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.scss'],
})
export class NavbarComponent {
  constructor(private router: Router, private route: ActivatedRoute){}

  routeToHome() {
    this.router.navigate(['/']);
  }
  routeToDash() {
    this.router.navigateByUrl('/home', {state: {user: history.state.user}});
  }
  routeToProfile() {
    this.router.navigateByUrl('/profile', {state: {user: history.state.user}});
  }
}
