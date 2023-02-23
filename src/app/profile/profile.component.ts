import { Component } from '@angular/core';
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
        this.userIdParameter = history.state.idToPass;
    }
    routeToHome() {
      this.router.navigate(['/']);
    }

}
