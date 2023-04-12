import { Component } from '@angular/core';
import { Service } from '../service';
import { User } from '../user';
import {formatDate} from '@angular/common'

import { Router } from '@angular/router';

@Component({
  selector: 'app-service-page',
  templateUrl: './service-page.component.html',
  styleUrls: ['./service-page.component.scss']
})
export class ServicePageComponent {


  service:Service = new Service("123", "Test Service", "Test Service description", 
    new Date("4/11/2023 11:00:00"), 120, 10, 15);

  userJoined : boolean = false;
  isBusiness : boolean = false;
  isEditing : boolean = false;

  constructor(private router:Router){}

  //empty user
  currentUser:User = {} as User;

  //for canceling edits
  backupService:Service = {} as Service; 

  //make sure the user has been passed throughout the routing
  ngOnInit()
  {
    //check that the user is not already signed up
    if(history.state != null)
    {
      this.currentUser = history.state.user;
      this.service = history.state.service;
      
      this.backupService = this.copyService(this.service);

      //set isBusiness boolean based on current user
      if(this.currentUser.accountType.toLowerCase() == "user")
      {
        this.isBusiness = false;
      }
      else
      {
        this.isBusiness = true;
      }

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
    //this.router.navigateByUrl('find-classes', {state:{user: this.currentUser}});
    this.router.navigateByUrl('home', {state:{user: this.currentUser}});
  }

  joinClass()
  {
    this.userJoined = true;
    this.currentUser.classes.push(this.service);
    
  }
  

  leaveClass()
  {
    this.userJoined = false;
    let index:number = this.currentUser.classes.findIndex((findService) => this.service.serviceId == findService.serviceId);
    
    //removes the service
    this.currentUser.classes.splice(index, 1);

  }

  editService()
  {
    this.isEditing = true;

  }

  //TODO: update the service in DB
  saveEdit()
  {
    this.backupService = this.copyService(this.service);
    this.isEditing = false;
  }

  cancelEdit()
  {
    this.service = this.copyService(this.backupService);
    this.isEditing = false;
  }

  debug()
  {
    this.isBusiness = !this.isBusiness;
  }

  formatDate(day: Date) {
    return formatDate(day, "MMM dd, yyyy", 'en')
  }

  copyService(input:Service)
  {
      var returnService:Service = {} as Service;

      returnService.serviceId = input.serviceId;
      returnService.name = input.name;
      returnService.description = input.description;
      returnService.start_date_time = new Date(input.start_date_time.getTime());
      returnService.length = input.length;
      returnService.capacity = input.capacity;
      returnService.price = input.price;

      return returnService;
  }
  
}
