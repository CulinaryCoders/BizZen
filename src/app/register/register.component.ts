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
  constructor(private router: Router, private activatedRoute:ActivatedRoute, private userService:UserService) {}

  userModel = new User("12345","", "", "business", []);
  isBusiness = false;
  confPass = "";
  errorMsg = "";

  registerForm = new FormGroup({
    username: new FormControl(''),
    email: new FormControl(''),
    password: new FormGroup({
      pass: new FormControl(''),
      confPass: new FormControl('')
    }),
    isBusiness: new FormControl(false)
  }
  )

  passwordsMatch() {
    console.log("confpass: ", this.confPass)
    return this.userModel.password === this.confPass;
  }

  allFieldsFilled() {
    // return this.registerForm.value.username && this.registerForm.value.username !== ""
    //   && this.registerForm.value.email && this.registerForm.value.email !== ""
    //   && this.registerForm.value.password?.pass && this.registerForm.value.password.pass !== ""
    //   && this.registerForm.value.password?.confPass && this.registerForm.value.password.confPass !== ""
    return this.userModel.username && this.userModel.username !== ""
      && this.userModel.userId && this.userModel.userId !== ""
      && this.userModel.password && this.userModel.password !== "" || false
      // && this.userModel.password?.confPass && this.userModel.password.confPass !== ""
  }

  async addUser(){
    // Promise interface

    this.userService.addUser(this.userModel.userId, this.userModel.username, this.userModel.password, this.userModel.accountType).then(
      user => {
        console.log("ADDING USER")
        this.userModel = user;
        console.log("success");
      }, err => {
        console.log("ERROR");
        console.log(err);
      }
    );
  }

  onSubmit() {
    console.log("===confpass: ", this.confPass)

    this.errorMsg = "";
    // console.log(this.registerForm.value);
    console.log(this.userModel);
    let user;

    if (this.allFieldsFilled() && this.registerForm.value.password?.pass === this.registerForm.value.password?.confPass) {
      if (!this.passwordsMatch()) {
        this.errorMsg = "ERROR Passwords must match"
      } else {
        // Note: since first & last name required, might need to test with dummy data
        user = new User(this.registerForm.value.email || "test", this.registerForm.value.username || "test", this.registerForm.value.password?.pass || "pass", "user", []);

        if (this.isBusiness)
          user.accountType = "business";
        else
          user.accountType = "user";

        // Update userModel to be sent
        this.userModel.username = this.registerForm.value.username || "test";

        // Send to db!
        this.addUser();
        if (this.isBusiness) {
          this.router.navigate(["/onboarding"])
        } else {
          this.router.navigate(["/businessOnboarding"])
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
