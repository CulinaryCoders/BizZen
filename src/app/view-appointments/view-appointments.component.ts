import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { User } from '../user';
import { Service } from '../service';
import { UserService } from '../user.service';
import { Appointment } from '../appointment';
import { ServiceAppointment } from '../service-appointment';

@Component({
  selector: 'app-view-appointments',
  templateUrl: './view-appointments.component.html',
  styleUrls: ['./view-appointments.component.scss']
})
export class ViewAppointmentsComponent {

  constructor(private router:Router, private userService:UserService){}

  user:User = {} as User;
  services:Service[] = [];
  userAppointments: ServiceAppointment[] = [];
  
  hasAppointments = false;

  ngOnInit()
  {

    //console.log(history.state.user);
    if(history.state != null)
    {
      this.user = history.state.user;
      this.userService.getUserServices(this.user.ID)
        .then((result)=>{
          this.hasAppointments = true; 
          this.userAppointments = result
        })
        .catch((reason) => console.log(reason));

    }

  }

  routeToProfile()
  {
    this.router.navigateByUrl('/profile', {state: {user: this.user }});

  }
  routeToService(serviceToPass:Service)
  {
    this.router.navigateByUrl('/class-summary', {state: {user: this.user, service: serviceToPass, prevPage: '/view-appointments'}});

  }

}
