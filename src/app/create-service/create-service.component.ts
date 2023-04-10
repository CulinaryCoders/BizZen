import { Component } from '@angular/core';
import {Router} from "@angular/router";
import {FormControl, FormGroup} from "@angular/forms";
import {User} from "../user";

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
    startDateTime: new FormControl(),
    length: new FormControl(),
    capacity: new FormControl(),
    pricePerUnit: new FormControl(),
    cancellationFee: new FormControl(),
  })

  verifyFields() {
    let errorMsg = "";
    let name = this.newService.value.name
    let description = this.newService.value.description
    let startDateTime = this.newService.value.startDateTime
    let length = this.newService.value.length
    let capacity = this.newService.value.capacity
    let pricePerUnit = this.newService.value.pricePerUnit
    let cancellationFee = this.newService.value.cancellationFee
    if (!name || name === "") {
      errorMsg += "ERROR Business Name Required -- "
    }
    if (!description || description === "") {
      errorMsg += "ERROR Business Description Required -- "
    }
    if (!startDateTime) {
      errorMsg += "ERROR Start Date & Time Required -- "
    }
    if (!length || length === 0) {
      errorMsg += "ERROR Length of Service Required -- "
    }
    if (!capacity || capacity === 0) {
      errorMsg += "ERROR Please specify how many participants -- "
    }
    if (!pricePerUnit) {
      errorMsg += "ERROR Please specify a price per unit -- "
    }
    if (!cancellationFee) {
      errorMsg += "ERROR Please specify a cancellation fee -- "
    }
    return errorMsg;
  }

  onSubmit() {
    this.errorMsg = this.verifyFields();
    if (this.errorMsg === "") {
      // this.newService.value.tags = this.selectedTags;

      // CONNECT BACKEND this.newService.value has all the info needed to add to DB User object
      let user = new User("","","","",[]);

      console.log(this.newService.value);

      this.router.navigate(['/home']);
    }
  }

  routeToHome() {
    this.router.navigate(['/']);
  }
}
