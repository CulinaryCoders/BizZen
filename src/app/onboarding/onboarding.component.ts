import { Component } from '@angular/core';
import {Router} from "@angular/router";
import {FormControl, FormGroup} from "@angular/forms";

@Component({
  selector: 'app-onboarding',
  templateUrl: './onboarding.component.html',
  styleUrls: ['./onboarding.component.scss'],
})
export class OnboardingComponent {
  constructor(private router: Router) {};

  errorMsg = "";

  onboardingForm = new FormGroup({
    firstName: new FormControl(''),
    lastName: new FormControl(''),
    interests: new FormControl(),
  })

  interests = ["Travel", "Cooking", "Yoga", "Technology", "Art"]
  selectedInterests: string[] = [];

  toggleInterest(interest: string) {
    if (this.selectedInterests.includes(interest)) {
      this.selectedInterests.splice(this.selectedInterests.indexOf(interest), 1)
    } else {
      this.selectedInterests.push(interest)
    }
  }

  onSubmit() {
    this.errorMsg = "";
    let fname = this.onboardingForm.value.firstName
    let lname = this.onboardingForm.value.lastName
    console.log("fname lname: ", fname, lname)
    if (!fname || fname === "") {
      this.errorMsg += "ERROR First Name Required -- "
    }
    if (!lname || lname === "") {
      this.errorMsg += "ERROR Last Name Required -- "
    }
    if (this.selectedInterests.length === 0) {
      this.errorMsg += "ERROR Please add at least 1 interest "
    }
    if (this.errorMsg === "") {
      this.onboardingForm.value.interests = this.selectedInterests;

      // CONNECT BACKEND this.newApptForm.value has all the info needed to add to DB User object
      console.log(this.onboardingForm.value);

      this.router.navigateByUrl('/profile', {state: {idToPass: fname }});
    }
  }

  routeToHome() {
    this.router.navigate(['/']);
  }
}
