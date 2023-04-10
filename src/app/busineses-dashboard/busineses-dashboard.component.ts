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

  ngOnInit() {
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
  // TODO: read from db
  businessOwnerView = true; // user.account_type
  view = "list";
  // TODO: read from db
  business = {
    id: 1,
    name: "CEN Recreational Center",
    bio: "Welcome to our community! Browse our classes for a variety of enriching opportunities.",
    created_at: new Date(),
    opening_time: "11:00",
    closing_time: "19:00"
  }


  services1 = [
    {
      id: 1,
      name: "Yoga",
      description: "Easy yoga class",
      start_date_time: new Date("5/7/2023 11:00:00"),
      length: 120,
      capacity: 10,
      price: 15
    },
    {
      id: 2,
      name: "Painting",
      description: "Intro to painting class",
      start_date_time: new Date("5/6/2023 11:00:00"),
      length: 120,
      capacity: 10,
      price: 15
    },
  ]

  formatDate(day: Date) {
    return formatDate(day, "MMM dd, yyyy", 'en')
  }

  formatTime(time: Date) {
    return formatDate(time, "HH:mm", 'en')
  }

  openAddService() {
    this.router.navigate(["/create-service"])
  }

  toggleView(type: string) {
    let listBtn = document.getElementById("list-view") || document.createElement("<p>");
    let calendarBtn = document.getElementById("calendar-view") || document.createElement("<p>");
    if (type === "calendar" && this.view === "list") {
      this.view = "calendar";
      listBtn.classList.remove("btn-primary");
      listBtn.classList.add("btn-secondary");
      calendarBtn.classList.add("btn-primary");
      calendarBtn.classList.remove("btn-secondary");
    } else if (type === "list" && this.view === "calendar") {
      this.view = "list";
      listBtn.classList.add("btn-primary");
      listBtn.classList.remove("btn-secondary");
      calendarBtn.classList.remove("btn-primary");
      calendarBtn.classList.add("btn-secondary");
    }
  }

  // TODO: feed in current user
  user = new User("","","","","", []);
  goToServicePage(serviceToPass: any) {
    this.router.navigateByUrl('/class-summary', {state: {user:this.user, service:serviceToPass}});
  }

  routeToHome() {
    this.router.navigate(['/home']);
  }
}
