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

  interests: string[] = [];

  toggleInterest(interest: string) {
    this.interests.push(interest)
    // if (this.interests.find(interest)) {
    //   this.interests.remove(interest)
    // } else {
    //   this.interests.push(interest)
    // }
  }

  onSubmit() {
    console.log(this.onboardingForm.value);
    // let text = document.getElementById("selectedInterests").textContent || ''

    const selected: HTMLElement = document.getElementById("selectedInterests")!
    // selected?.textContent("hi")
  }

  routeToHome() {
    this.router.navigate(['/']);
  }
}
