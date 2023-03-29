import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { Service } from '../service';
import { User } from '../user';
import { ServiceOffering } from '../service-offering';

@Component({
  selector: 'app-find-classes',
  templateUrl: './find-classes.component.html',
  styleUrls: ['./find-classes.component.scss'],
})


export class FindClassesComponent {

  testService:Service = new Service("001", "Test Service", "A service for testing", 
    new ServiceOffering("01/01/2023", "02/01/2023", 25));
  
  testService2:Service = new Service("002", "Another Test Service", "A second service for testing", 
  new ServiceOffering("02/01/2023", "03/01/2023", 15));


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
