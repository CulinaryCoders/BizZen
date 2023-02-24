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

  userModel = new User("12345","", "", false);
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

  passwordMatch(password: string, confirmPassword: string):ValidatorFn {
    return (formGroup: AbstractControl):{ [key: string]: any } | null => {
      const passwordControl = formGroup.get(password);
      const confirmPasswordControl = formGroup.get(confirmPassword);

      if (!passwordControl || !confirmPasswordControl) {
        return null;
      }

      if (
        confirmPasswordControl.errors &&
        !confirmPasswordControl.errors['passwordMismatch']
      ) {
        return null;
      }

      if (passwordControl.value !== confirmPasswordControl.value) {
        confirmPasswordControl.setErrors({ passwordMismatch: true });
        return { passwordMismatch: true }
      } else {
        confirmPasswordControl.setErrors(null);
        return null;
      }
    };
  }

  allFieldsFilled() {
    // return this.registerForm.value.username && this.registerForm.value.username !== ""
    //   && this.registerForm.value.email && this.registerForm.value.email !== ""
    //   && this.registerForm.value.password?.pass && this.registerForm.value.password.pass !== ""
    //   && this.registerForm.value.password?.confPass && this.registerForm.value.password.confPass !== ""
    return this.userModel.username && this.userModel.username !== ""
      && this.userModel.userId && this.userModel.userId !== ""
      && this.userModel.password && this.userModel.password !== ""
      // && this.userModel.password?.confPass && this.userModel.password.confPass !== ""
  }

  async addUser(){
    // Promise interface
    this.userService.addUser(this.userModel.userId, this.userModel.username, this.userModel.password, this.userModel.isBusiness).then(
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
    this.errorMsg = "";
    // console.log(this.registerForm.value);
    console.log(this.userModel);
    let user;
    if (this.allFieldsFilled() && this.registerForm.value.password?.pass === this.registerForm.value.password?.confPass) {
      // Note: since first & last name required, might need to test with dummy data
      user = new User(this.registerForm.value.email || "test", this.registerForm.value.username || "test", this.registerForm.value.password?.pass || "pass", this.registerForm.value.isBusiness || false);

      // Update userModel to be sent
      this.userModel.username = this.registerForm.value.username || "test";

      // Send to db!
      this.addUser();
    } else {
      this.errorMsg = "ERROR Please fill out all fields correctly!"
    }
  }

  routeToLogin() {
    this.router.navigate(['/login']);
  }
  routeToHome() {
    this.router.navigate(['/']);
  }
}
