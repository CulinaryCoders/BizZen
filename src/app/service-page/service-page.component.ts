import { Component } from '@angular/core';
import { Service } from '../service';
import { User } from '../user';
import { ServiceService } from '../service.service';
import {formatDate} from '@angular/common'

import { Router } from '@angular/router';
import { UserService } from '../user.service';

@Component({
  selector: 'app-service-page',
  templateUrl: './service-page.component.html',
  styleUrls: ['./service-page.component.scss']
})
export class ServicePageComponent {


  service:Service = new Service("123", "Test Service", "Test Service desc", 
    new Date("4/11/2023 11:00:00"), 120, 10, 15);

  userJoined : boolean = false;
  isBusiness : boolean = false;
  isEditing : boolean = false;

  constructor(private router:Router, private serviceService : ServiceService, private userService:UserService){}

  //empty user
  currentUser:User = {} as User;

  //for canceling edits
  backupService:Service = {} as Service; 

  usersSignedUp : User[] = [];

  //make sure the user has been passed throughout the routing
  ngOnInit()
  {
    //check that the user is not already signed up
    if(history.state != null)
    {
      this.currentUser = history.state.user;
      this.service = history.state.service;

      this.backupService = this.copyService(this.service);

      //get all users attached to the current service (for business view)
      this.serviceService.getUsers(this.service.ID)
       .then((users) => {this.usersSignedUp = users})   //success


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
      let index:number = this.currentUser.classes.findIndex((findService) => this.service.ID == findService.ID);

      //the user has already joined if the class was found
      if(index != -1)
        this.userJoined = true;

    }
    else
    {
      console.log("ERROR: the browser state is null. DID you pass the user correctly?");
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
    this.userService.addService(this.service.ID, this.currentUser.ID).then();

  }
  

  leaveClass()
  {
    this.userJoined = false;
    let index:number = this.currentUser.classes.findIndex((findService) => this.service.ID == findService.ID);
    
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

    //backend connection
    this.serviceService.updateService(
      this.service.ID, this.service.name, this.service.desc, 
      this.service.start_date_time, this.service.length, this.service.length, this.service.price).then();
    this.isEditing = false;

    history.state.service = this.service;
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

  //performs a deep copy of an input service
  copyService(input:Service)
  {
      var returnService:Service = {} as Service;

      returnService.ID = input.ID;
      returnService.name = input.name;
      returnService.desc = input.desc;
      returnService.start_date_time = new Date(input.start_date_time);
      returnService.length = input.length;
      returnService.capacity = input.capacity;
      returnService.price = input.price;

      return returnService;

  }
  
}
