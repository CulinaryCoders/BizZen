import { Component } from '@angular/core';
import {Router} from "@angular/router";
import {FormControl, FormGroup} from "@angular/forms";
import {User} from "../user";

@Component({
  selector: 'app-business-onboarding',
  templateUrl: './business-onboarding.component.html',
  styleUrls: ['./business-onboarding.component.scss']
})
export class BusinessOnboardingComponent {
  constructor(private router: Router) {};

  errorMsg = "";

  onboardingForm = new FormGroup({
    businessName: new FormControl(''),
    businessDescription: new FormControl(''),
    openingTime: new FormControl(''),
    closingTime: new FormControl(''),
    tags: new FormControl(),
  })

  tags = ["Travel", "Cooking", "Yoga", "Technology", "Art"]
  selectedTags: string[] = [];

  toggleTags(tag: string) {
    if (this.selectedTags.includes(tag)) {
      this.selectedTags.splice(this.selectedTags.indexOf(tag), 1)
    } else {
      this.selectedTags.push(tag)
    }
  }

  validStartEndTime() {
    let start = this.onboardingForm.value.openingTime;
    let end = this.onboardingForm.value.closingTime;
    if (start && end) {
      let startJSDate = new Date();
      startJSDate.setHours(Number(start[0]+start[1]));
      startJSDate.setMinutes(Number(start[3]+start[4]));

      let endJSDate = new Date();
      endJSDate.setHours(Number(end[0]+end[1]));
      endJSDate.setMinutes(Number(end[3]+end[4]));

      return startJSDate < endJSDate;
    }
    return false;
  }

  onSubmit() {
    this.errorMsg = "";
    let businessName = this.onboardingForm.value.businessName
    let businessDescription = this.onboardingForm.value.businessDescription
    let openingTime = this.onboardingForm.value.openingTime
    let closingTime = this.onboardingForm.value.closingTime
    console.log("name open close: ", businessName, openingTime, closingTime)
    if (!businessName || businessName === "") {
      this.errorMsg += "ERROR Business Name Required -- "
    }
    if (!businessDescription || businessDescription === "") {
      this.errorMsg += "ERROR Business Description Required -- "
    }
    if (!openingTime || openingTime === "") {
      this.errorMsg += "ERROR Opening Time Required -- "
    }
    if (!closingTime || closingTime === "") {
      this.errorMsg += "ERROR Closing Time Required -- "
    }
    if (!this.validStartEndTime()) {
      this.errorMsg += "ERROR Closing Time must be after Opening -- "
    }
    if (this.selectedTags.length === 0) {
      this.errorMsg += "ERROR Please add at least 1 tag "
    }
    if (this.errorMsg === "") {
      this.onboardingForm.value.tags = this.selectedTags;

      // CONNECT BACKEND this.onboardingForm.value has all the info needed to add to DB User object
      console.log(this.onboardingForm.value);

      let user = new User("","", "", "", "", "", [])
      this.router.navigate(['/profile'], {state: {user: user }});
    }
  }

  routeToHome() {
    this.router.navigate(['/']);
  }
}
