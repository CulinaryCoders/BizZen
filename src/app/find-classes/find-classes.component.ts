import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { Service } from '../service';
import { User } from '../user';

@Component({
  selector: 'app-find-classes',
  templateUrl: './find-classes.component.html',
  styleUrls: ['./find-classes.component.scss'],
})


export class FindClassesComponent {

  testService:Service = new Service("123", "Test Service", "Test Service description", 
                                new Date("4/11/2023 11:00:00"), 120, 10, 15);
  
  testService2:Service = new Service("123", "Test Service 2", "Test Service 2 description", 
                                new Date("4/11/2023 11:00:00"), 120, 10, 15);


  testArray:Service[] = [this.testService, this.testService2, this.testService, this.testService2];

  constructor(private router:Router){}
 
  user:User = {} as User;

  ngOnInit()
  {
    if(history.state != null)
      this.user = history.state.user;
  }

  routeToService(serviceToPass:Service)
  {

    this.router.navigateByUrl('/class-summary', {state: {user:this.user, service:serviceToPass}});

  }
  routeToUserPage()
  {
    this.router.navigateByUrl('/profile', {state: {user: this.user }});

  }
}
