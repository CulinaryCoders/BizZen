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
  constructor(private router: Router, private serviceService: ServiceService) {};
  // @ts-ignore
  services: any[];

  srv: any[] = [];


  // @ts-ignore
  user: User;

  ngOnInit() {
    this.user = history.state.user;
    this.serviceService.getServices().then((res) => {
      for (let i=0; i<res.length; i++) {
        this.srv.push(res[i]);
      }
      this.services = this.srv;
    });
    this.services = this.srv;
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
