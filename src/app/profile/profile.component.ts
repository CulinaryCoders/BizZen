import { Component } from '@angular/core';
import { User } from '../user';
import { Router, ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.scss']
})
export class ProfileComponent {

    userIdParameter : string = "";

    constructor(private router: Router, private route: ActivatedRoute){}

    ngOnInit()
    {
      //this.route.params.subscribe((params:Params) => this.userIdParameter = params['idToPass'])
        if(history.state.user != null)
          this.userIdParameter = history.state.user.first_name;
        else
          this.userIdParameter = "ERROR: userIdParameter is null; no user was passed!";
    }
    routeToHome() {
      this.router.navigate(['/']);
    }
    routeToClasses() {

      this.router.navigateByUrl('/home', {state: {user: history.state.user}});
    } 
    routeToAppointments() {

      this.router.navigateByUrl('/view-appointments', {state: {user: history.state.user}});
    }

}
