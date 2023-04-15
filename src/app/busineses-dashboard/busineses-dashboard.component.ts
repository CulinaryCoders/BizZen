import {Component} from '@angular/core';
import {Router} from "@angular/router";
import {formatDate} from "@angular/common";
import {User} from "../user";
import {ServiceService} from "../service.service";

@Component({
  selector: 'app-busineses-dashboard',
  templateUrl: './busineses-dashboard.component.html',
  styleUrls: ['./busineses-dashboard.component.scss']
})

export class BusinesesDashboardComponent {
  // @ts-ignore
  services: any[]; // services to be displayed (filtered)

  allServices: any[] = [];

  srv: any[] = [];

  // @ts-ignore
  user: User;

  viewDateRange: any[2]
  = [
    new Date(new Date().setDate(1)), // first of the month
    new Date(new Date().getFullYear(), new Date().getMonth()+1, 0) // last day of the month
  ];

  constructor(private router: Router, private serviceService: ServiceService) {};

  ngOnInit() {
    this.user = history.state.user;
    this.serviceService.getServices().then((res) => {
      for (let i=0; i<res?.length; i++) {
        this.srv.push(res[i]);
      }
      this.allServices = this.srv.sort((a,b) => new Date(a.start_date_time).getTime() - new Date(b.start_date_time).getTime());
      this.services = this.allServices;
    });
  }

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

  filterByDateRange(e: any[]) {
    this.viewDateRange = e;
    this.services = this.allServices.filter((s) => {
      let endDateTime = this.getEndDate(new Date(s.start_date_time), s.length);
      if (new Date(s.start_date_time) >= this.viewDateRange[0] && endDateTime <= this.viewDateRange[1]) {
        return s;
      }
    });
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

  // Returns end date from startDateTime and length of service
  getEndDate(startDateTime: Date, length: number) {
    return new Date(startDateTime.getTime() + length*60000);
  }

  // Gets child data from calendar for view range
  updateDateRange(e: any[]) {
    this.viewDateRange = e;
    this.filterByDateRange(e);
  }

  routeToHome() {
    this.router.navigate(['/home']);
  }
}
