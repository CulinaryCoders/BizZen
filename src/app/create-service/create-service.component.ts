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

  onSubmit() {
    this.errorMsg = "";
    let name = this.newService.value.name
    let description = this.newService.value.description
    let startTime = this.newService.value.startTime
    let endTime = this.newService.value.endTime
    let numParticipants = this.newService.value.numParticipants
    let pricePerUnit = this.newService.value.pricePerUnit
    console.log("name open close: ", name, startTime, endTime)
    if (!name || name === "") {
      this.errorMsg += "ERROR Business Name Required -- "
    }
    if (!description || description === "") {
      this.errorMsg += "ERROR Business Description Required -- "
    }
    if (!startTime || startTime === "") {
      this.errorMsg += "ERROR Opening Time Required -- "
    }
    if (!endTime || endTime === "") {
      this.errorMsg += "ERROR Closing Time Required -- "
    }
    if (numParticipants === 0) {
      this.errorMsg += "ERROR Please specify how many participants -- "
    }
    if (!pricePerUnit) {
      this.errorMsg += "ERROR Please specify a price per unit -- "
    }
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
