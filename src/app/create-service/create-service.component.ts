import { Component } from '@angular/core';
import {Router} from "@angular/router";
import {FormControl, FormGroup} from "@angular/forms";

@Component({
  selector: 'app-create-service',
  templateUrl: './create-service.component.html',
  styleUrls: ['./create-service.component.scss']
})
export class CreateServiceComponent {
  constructor(private router: Router) {};

  errorMsg = "";

  newService = new FormGroup({
    name: new FormControl(''),
    description: new FormControl(''),
    type: new FormControl(),
    startTime: new FormControl(''),
    endTime: new FormControl(''),
    numParticipants: new FormControl(0),
    pricePerUnit: new FormControl(0),
  })

  serviceTypes = ["Class", "Lecture", "Tutoring", "Demonstration", "Other"]
  selectedServiceType: string = "";

  selectType(type: string) {
    this.selectedServiceType = type;
  }

  verifyFields() {
    let errorMsg = "";
    let name = this.newService.value.name
    let description = this.newService.value.description
    let startTime = this.newService.value.startTime
    let endTime = this.newService.value.endTime
    let numParticipants = this.newService.value.numParticipants
    let pricePerUnit = this.newService.value.pricePerUnit
    if (!name || name === "") {
      errorMsg += "ERROR Business Name Required -- "
    }
    if (!description || description === "") {
      errorMsg += "ERROR Business Description Required -- "
    }
    if (!startTime || startTime === "") {
      errorMsg += "ERROR Opening Time Required -- "
    }
    if (!endTime || endTime === "") {
      errorMsg += "ERROR Closing Time Required -- "
    }
    if (this.validStartEndTime()) {
      errorMsg += "ERROR End must be after start --"
    }
    if (numParticipants === 0) {
      errorMsg += "ERROR Please specify how many participants -- "
    }
    if (!pricePerUnit) {
      errorMsg += "ERROR Please specify a price per unit -- "
    }
    return errorMsg;
  }

  validStartEndTime() {
    let start = this.newService.value.startTime;
    let end = this.newService.value.endTime;
    if (start && end) {
      let startJSDate = new Date();
      startJSDate.setHours(Number(start[0]+start[1]));
      startJSDate.setMinutes(Number(start[3]+start[4]));

      let endJSDate = new Date();
      endJSDate.setHours(Number(end[0]+end[1]));
      endJSDate.setMinutes(Number(end[3]+end[4]));

      return startJSDate > endJSDate;
    }
    return false;
  }

  onSubmit() {
    this.errorMsg = this.verifyFields();
    if (this.errorMsg === "") {
      // this.newService.value.tags = this.selectedTags;

      // CONNECT BACKEND this.newService.value has all the info needed to add to DB User object
      console.log(this.newService.value);

      this.router.navigate(['/profile']);
    }
  }

  routeToHome() {
    this.router.navigate(['/']);
  }
}
