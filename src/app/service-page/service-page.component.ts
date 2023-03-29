import { Component } from '@angular/core';
import { Service } from '../service';
import { ServiceOffering } from '../service-offering';
import { User } from '../user';

import { Router } from '@angular/router';

@Component({
  selector: 'app-service-page',
  templateUrl: './service-page.component.html',
  styleUrls: ['./service-page.component.scss']
})
export class ServicePageComponent {


  serviceOffer : ServiceOffering = new ServiceOffering("01/01/2023", "02/01/2023", 25);
  service:Service = new Service("123", "Test Service", "An example of a description for a test service.", this.serviceOffer);

  userJoined : boolean = false;

  constructor(private router:Router){}

  //empty user
  currentUser:User = {} as User;

  //make sure the user has been passed throughout the routing
  ngOnInit()
  {
    //check that the user is not already signed up
    if(history.state != null)
    {
      this.currentUser = history.state.user;
      this.service = history.state.service;

      //find a service so that it matches this service
      let index:number = this.currentUser.classes.findIndex((findService) => this.service.serviceId == findService.serviceId);

      //the user has already joined if the class was found
      if(index != -1)
        this.userJoined = true;

    }
    else
    {
      console.log("ERROR: the browser state is null. Did you pass the user correctly?");
    }
  }

  routeToFindClass()
  {
    this.router.navigateByUrl('find-classes', {state:{user: history.state.user}});
  }

  joinClass()
  {
    this.userJoined = true;
    history.state.user.classes.push(this.service);
    
  }

  leaveClass()
  {
    this.userJoined = false;
    let index:number = this.currentUser.classes.findIndex((findService) => this.service.serviceId == findService.serviceId);
    
    //removes the service
    history.state.user.classes.splice(index, 1);

  }

}
