import { Component } from '@angular/core';
import { Service } from '../service';
import { ServiceOffering } from '../service-offering';

import { Router } from '@angular/router';

@Component({
  selector: 'app-service-page',
  templateUrl: './service-page.component.html',
  styleUrls: ['./service-page.component.scss']
})
export class ServicePageComponent {


  serviceOffer : ServiceOffering = new ServiceOffering("01/01/2023", "02/01/2023", 25);
  service:Service = new Service("00000001", "Test Service", "An example of a description for a test service.", this.serviceOffer);

  userJoined : boolean = false;

  constructor(private router:Router){}

  routeToFindClass()
  {
    this.router.navigate(['find-classes']);
  }

  joinClass()
  {
    this.userJoined = true;
    console.log("class joined");
  }
  
  leaveClass()
  {
    this.userJoined = false;
    console.log("class joined");
  }

}
