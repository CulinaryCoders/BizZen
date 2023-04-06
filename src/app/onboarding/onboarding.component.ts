import { Component } from '@angular/core';
import {Router} from "@angular/router";
import {FormControl, FormGroup} from "@angular/forms";
import {User} from "../user";
import {UserService} from '../user.service';

@Component({
  selector: 'app-onboarding',
  templateUrl: './onboarding.component.html',
  styleUrls: ['./onboarding.component.scss'],
})
export class OnboardingComponent {
  constructor(private router: Router, private userService:UserService) {};

  errorMsg = "";
  userModel = history.state.user;

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

  async addUser(){
    console.log("user model: ", this.userModel)
    // Promise interface
    // this.userService.addUser(this.userModel.userId, this.userModel.username, this.userModel.password, this.userModel.accountType).then(
    //   user => {
    //     console.log("ADDING USER")
    //     this.userModel = user;
    //     console.log("success");
    //   }, err => {
    //     console.log("ERROR");
    //     console.log(err);
    //   }
    // );
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
      // Send to db!
      this.addUser().then(() => {
        console.log(this.onboardingForm.value);
        // let user = new User("12345", fname || "Gatey", "pass", "user", [])
        // this.router.navigateByUrl('/profile', {state: {user: user}});
        // this.router.navigateByUrl('/profile', {state: {user: this.userModel}});
      }).catch((err) => {
        console.log("error: ", err)
      });
    }
  }

  routeToHome() {
    this.router.navigate(['/']);
  }
}
