import {Component} from '@angular/core';
import {Router} from "@angular/router";
import {formatDate} from "@angular/common";
import {Service} from "../service";
import {User} from "../user";

@Component({
  selector: 'app-busineses-dashboard',
  templateUrl: './busineses-dashboard.component.html',
  styleUrls: ['./busineses-dashboard.component.scss']
})

export class BusinesesDashboardComponent {
  // @ts-ignore
  services: any[];
  // @ts-ignore
  user: User;

  ngOnInit() {
    this.user = history.state.user;
    this.services = [
      {
        id: 1,
        name: "Yoga",
        description: "Easy yoga class",
        start_date_time: new Date("4/6/2023 11:00:00"),
        length: 120,
        capacity: 10,
        price: 15
      },
      {
        id: 2,
        name: "Painting",
        description: "Intro to painting class",
        start_date_time: new Date("4/6/2023 10:00:00"),
        length: 120,
        capacity: 10,
        price: 15
      },
      {
        id: 3,
        name: "Weightlifting",
        description: "Do you even lift",
        start_date_time: new Date("4/11/2023 15:00:00"),
        length: 120,
        capacity: 10,
        price: 15
      },
      {
        id: 4,
        name: "Computers",
        description: "Learn how to use computer applications like Excel",
        start_date_time: new Date("4/22/2023 13:00:00"),
        length: 120,
        capacity: 10,
        price: 15
      },
    ];
  }

  constructor(private router: Router) {};
  businessOwnerView = history.state.user.accountType === "business";
  // TODO: read from db
  business = {
    id: 1,
    name: "CEN Recreational Center",
    bio: "Welcome to our community! Browse our classes for a variety of enriching opportunities.",
    created_at: new Date(),
    opening_time: "11:00",
    closing_time: "19:00"
  }

  formatDate(day: Date) {
    return formatDate(day, "MMM dd, yyyy", 'en')
  }

  formatTime(time: Date) {
    return formatDate(time, "HH:mm", 'en')
  }

  openAddService() {
    this.router.navigateByUrl("/create-service", {state: {user: history.state.user}})
  }

  goToServicePage(serviceToPass: any) {
    this.router.navigateByUrl('/class-summary', {state: {user: history.state.user, service:serviceToPass}});
  }

  routeToHome() {
    this.router.navigate(['/home']);
  }
}
