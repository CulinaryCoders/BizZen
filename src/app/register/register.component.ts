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
  // userModel = new User("","","", "", "", "", []);
  isBusiness = false;
  confPass = "";
  errorMsg = "";

  registerForm = new FormGroup({
    first_name: new FormControl(''),
    last_name: new FormControl(''),
    email: new FormControl(''),
    password: new FormControl(''),
    accountType: new FormControl("user")
  })

  passwordsMatch() {
    return this.registerForm.value.password === this.confPass;
  }

  allFieldsFilled() {
    return this.registerForm.value.first_name && this.registerForm.value.first_name !== ""
      && this.registerForm.value.last_name && this.registerForm.value.last_name !== ""
      && this.registerForm.value.email && this.registerForm.value.email !== ""
      && this.registerForm.value.password && this.registerForm.value.password !== "" || false
  }

  onSubmit() {
    this.errorMsg = "";
    if (this.allFieldsFilled()) {
      if (!this.passwordsMatch()) {
        this.errorMsg = "ERROR Passwords must match"
      } else {
        if (this.isBusiness) {
          this.registerForm.value.accountType = "business"
        } else {
          this.registerForm.value.accountType = "user"
        }

        // Creates user in DB
        this.userService.addUser(this.registerForm.value.first_name!, this.registerForm.value.last_name!, this.registerForm.value.email!, this.registerForm.value.password!, this.registerForm.value.accountType)
          .then((result) =>
            //route to profile
            this.router.navigateByUrl('/profile', {state: {user: result.user }})
          ).catch((err) => {
            console.error("ERROR creating user: ", err)
        });
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
