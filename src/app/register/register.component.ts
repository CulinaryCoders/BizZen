import { Component } from '@angular/core';
import {FormGroup, FormControl, Validator, AbstractControl, ValidationErrors, ValidatorFn} from '@angular/forms';
import {ActivatedRoute, Router} from "@angular/router";
import {User} from "../user";
import {UserService} from '../user.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.scss']
})
export class RegisterComponent {
  constructor(private router: Router, private activatedRoute:ActivatedRoute) {}
  userModel = new User("","", "", "", "", []);
  isBusiness = false;
  confPass = "";
  errorMsg = "";

  registerForm = new FormGroup({
    firstName: new FormControl(''),
    lastName: new FormControl(''),
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
    return this.userModel.firstName && this.userModel.firstName !== ""
      && this.userModel.lastName && this.userModel.lastName !== ""
      && this.userModel.email && this.userModel.email !== ""
      && this.userModel.password && this.userModel.password !== "" || false
  }

  onSubmit() {
    this.errorMsg = "";
    console.log(this.userModel);
    let user;

    if (this.allFieldsFilled()) {
      if (!this.passwordsMatch()) {
        this.errorMsg = "ERROR Passwords must match"
      } else {
        if (this.isBusiness) {
          console.log("is business")

          this.userModel.accountType = "business"
          this.router.navigateByUrl('/profile', {state: {user: this.userModel }});

        } else {
          console.log("is NOTbusiness")

          this.userModel.accountType = "user"
          this.router.navigateByUrl('/profile', {state: {user: this.userModel }});

        }
        // TODO: add user to db

        if (this.isBusiness) {
          // this.router.navigate(["/business-onboarding"]);
          this.router.navigateByUrl('/profile', {state: {user: this.userModel }});

        } else {
          // this.router.navigateByUrl('/profile', {state: {user: this.userModel }});
          this.router.navigateByUrl('/profile', {state: {user: this.userModel }});
        }
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
