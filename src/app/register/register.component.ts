import { Component } from '@angular/core';
import {FormGroup, FormControl, Validator, AbstractControl, ValidationErrors, ValidatorFn} from '@angular/forms';
import {Router} from "@angular/router";
import {User} from "../user";

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.scss']
})
export class RegisterComponent {
  constructor(private router: Router) {}

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

  toggleAccountType() {
    if (this.registerForm.value.isBusiness) {
      this.registerForm.value.isBusiness = false;
    } else {
      this.registerForm.value.isBusiness = true;
    }
  }

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
    return this.registerForm.value.username && this.registerForm.value.username !== ""
      && this.registerForm.value.email && this.registerForm.value.email !== ""
      && this.registerForm.value.password?.pass && this.registerForm.value.password.pass !== ""
      && this.registerForm.value.password?.confPass && this.registerForm.value.password.confPass !== ""
  }

  onSubmit() {
    console.log(this.registerForm.value);
    let user;
    if (this.allFieldsFilled() && this.registerForm.value.password?.pass === this.registerForm.value.password?.confPass) {
      // Note: since first & last name required, might need to test with dummy data
      user = new User(this.registerForm.value.email || "test", this.registerForm.value.username || "test", this.registerForm.value.password?.pass || "pass", this.registerForm.value.isBusiness || false);
      // Send to db!
    }
  }

  routeToLogin() {
    this.router.navigate(['/login']);
  }
  routeToHome() {
    this.router.navigate(['/']);
  }
}
