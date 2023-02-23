import {Component} from '@angular/core';
import {User} from '../user';

//passes info between components
import {Router, ActivatedRoute} from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']

})

export class LoginComponent {

  model = new User("", "", false);
  userId : String = "12345";

  constructor(private router:Router, private activatedRoute:ActivatedRoute){}

  submitted = false;

  onSubmit() {
    this.submitted = true;
  
    if(!this.model.isBusiness)
      this.router.navigateByUrl('/profile', {state: {idToPass: this.model.username }});
  }
  routeToRegister() {
    this.router.navigate(['/register']);
  }
  routeToHome() {
    this.router.navigate(['/']);
  }
}
