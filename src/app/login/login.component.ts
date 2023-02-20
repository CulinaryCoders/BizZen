import {Component} from '@angular/core';
import {User} from '../user';

//passes info between components
import {Router, ActivatedRoute, ParamMap} from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']

})

export class LoginComponent {

  //userTypes = ["Customer", "Business"];

  model = new User("", "", false);

  constructor(private router:Router){}

  submitted = false;

  onSubmit() {
    this.submitted = true;

    if(!this.model.isBusiness)
      this.router.navigate(['/profile']);
  }
  routeToRegister() {
    this.router.navigate(['/register']);
  }
  routeToHome() {
    this.router.navigate(['/']);
  }
}
