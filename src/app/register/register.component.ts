import { Component } from '@angular/core';
import {FormGroup, FormControl, Validator, AbstractControl, ValidationErrors, ValidatorFn} from '@angular/forms';
import {Router} from "@angular/router";

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.scss']
})
export class RegisterComponent {
  constructor(private router: Router) {}
  // public registerForm : FormGroup;

  registerForm = new FormGroup({
    firstName: new FormControl(''),
    lastName: new FormControl(''),
    password: new FormGroup({
      pass: new FormControl(''),
      confPass: new FormControl('')
    }),
    isBusiness: new FormControl('')
  }
  // , {validators: this.passwordMatch}
  )

  // comparisonValidator() : ValidatorFn{
  //   return (registerForm: FormGroup): ValidationErrors => {
  //     const control1 = registerForm.controls['password'].pass;
  //     const control2 = registerForm.controls['password'].confPass;
  //     if (control1.value !== control2.value) {
  //       control2.setErrors({notEquivalent: true});
  //     } else {
  //       control2.setErrors(null);
  //     }
  //     return;
  //   };
  // }

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

  onSubmit() {
    console.log(this.registerForm.value);
  }

  routeToLogin() {
    this.router.navigate(['/login']);
  }
  routeToHome() {
    this.router.navigate(['/']);
  }
}
