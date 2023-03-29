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

  constructor(private router:Router){}

  serviceOffer : ServiceOffering = new ServiceOffering("01/01/2023", "02/01/2023", 25);
  service:Service = new Service("00000001", "Test Service", "An example of a description for a test service.", this.serviceOffer);


  routeToFindClass()
  {
    this.router.navigate(['find-classes']);
  }

}
