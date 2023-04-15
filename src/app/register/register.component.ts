import { Component } from '@angular/core';
import {FormGroup, FormControl, Validator, AbstractControl, ValidationErrors, ValidatorFn} from '@angular/forms';
import {ActivatedRoute, Router} from "@angular/router";
import {User} from "../user";
import {UserService} from "../user.service"
@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.scss']
})
export class RegisterComponent {
  constructor(private router: Router, private activatedRoute:ActivatedRoute, private userService:UserService) {}
  userModel = new User("","","", "", "", "", []);
  isBusiness = false;
  confPass = "";
  errorMsg = "";

  registerForm = new FormGroup({
    first_name: new FormControl(''),
    last_name: new FormControl(''),
    email: new FormControl(''),
    password: new FormGroup({
      pass: new FormControl(''),
      confPass: new FormControl('')
    }),
    isBusiness: new FormControl(false)
  }
  )

  passwordsMatch() {
    return this.userModel.password === this.confPass;
  }

  allFieldsFilled() {
    return this.userModel.first_name && this.userModel.first_name !== ""
      && this.userModel.last_name && this.userModel.last_name !== ""
      && this.userModel.email && this.userModel.email !== ""
      && this.userModel.password && this.userModel.password !== "" || false
  }

  onSubmit() {
    this.errorMsg = "";
    if (this.allFieldsFilled()) {
      if (!this.passwordsMatch()) {
        this.errorMsg = "ERROR Passwords must match"
      } else {
        if (this.isBusiness) {
          this.userModel.accountType = "business"
          this.router.navigateByUrl('/profile', {state: {user: this.userModel }});
        } else {
          this.userModel.accountType = "user"
          this.router.navigateByUrl('/profile', {state: {user: this.userModel }});
        }

        // Creates user in DB
        this.userService.addUser(this.userModel.first_name, this.userModel.last_name, this.userModel.email, this.userModel.password, this.userModel.accountType)
          .then((result) =>
            //route to profile
            this.router.navigateByUrl('/profile', {state: {user: result }})
          );

        // Routes to profile
        //this.router.navigateByUrl('/profile', {state: {user: userResponse }});
      }
    }
  }

  routeToLogin() {
    this.router.navigate(['/login']);
  }
  routeToHome() {
    this.router.navigate(['/']);
  }
}
