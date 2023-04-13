import {Component} from '@angular/core';
import {User} from '../user';
import {UserService} from '../user.service';

//passes info between components
import {Router, ActivatedRoute} from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']

})

export class LoginComponent {


  model = new User("","", "", "", "", []);

  checkbox : boolean = false;
  userExists : boolean = true;

  constructor(private router:Router, private activatedRoute:ActivatedRoute, private userService:UserService){}

  submitted = false;

  onSubmit() {
    this.submitted = true;

    if(this.checkbox)
    {
      this.model.accountType = "business";
    }
    else
    {
      this.model.accountType = "user";
    }

  
    this.userService.login(this.model.email, this.model.password)
      .then((user) => {this.successfulLogin(user)})   //success
      .catch(()=>this.unsuccessfulLogin());           //failure
   
  }
  successfulLogin(returnedUser : void|User)
  {
    this.model = returnedUser as User;
    this.userExists = true;

    this.router.navigateByUrl('/profile', {state: {user: returnedUser as User }});
  }

  unsuccessfulLogin()
  {
    this.userExists = false;
  }

  routeToRegister() {
    this.router.navigate(['/register']);
  }
  routeToHome() {
    this.router.navigate(['/']);
  }
}
