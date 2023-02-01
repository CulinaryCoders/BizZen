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

  userTypes = ["Customer", "Business"];

  model = new User("", "", "");

  constructor(private router:Router){}
  
  submitted = false;

  onSubmit() { 
    this.submitted = true; 
    
    if(this.model.type === "customer")
      this.router.navigate(['/profile']);
  }

}