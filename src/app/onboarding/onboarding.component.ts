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

  onboardingForm = new FormGroup({
    firstName: new FormControl(''),
    lastName: new FormControl(''),
    interests: new FormControl([]),
  });

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
    console.log(this.onboardingForm.value);
    // let text = document.getElementById("selectedInterests").textContent || ''
    this.router.navigate(['/profile']);

  }

  routeToHome() {
    this.router.navigate(['/']);
  }
}
