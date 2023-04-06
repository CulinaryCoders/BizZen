import { Component } from '@angular/core';
import {Router} from "@angular/router";
import {formatDate} from "@angular/common";

@Component({
  selector: 'app-busineses-dashboard',
  templateUrl: './busineses-dashboard.component.html',
  styleUrls: ['./busineses-dashboard.component.scss']
})
export class BusinesesDashboardComponent {
  constructor(private router: Router) {};
  formatDate(day: Date) {
    return formatDate(day, "MMM dd, yyyy", 'en')
  }

  formatTime(time: Date) {
    return formatDate(time, "HH:mm", 'en')
  }

  routeToHome() {
    this.router.navigate(['/home']);
  }

  services = [
    {
      name: "Yoga",
      description: "Easy yoga class",
      start_date_time: new Date("5/7/2023 11:00:00"),
      length: 120,
      capacity: 10,
      price: 15
    },
    {
      name: "Painting",
      description: "Intro to painting class",
      start_date_time: new Date("5/6/2023 11:00:00"),
      length: 120,
      capacity: 10,
      price: 15
    },
  ]
}
