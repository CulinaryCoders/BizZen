import { Component } from '@angular/core';
import {Router} from "@angular/router";
import {FormControl, FormGroup} from "@angular/forms";
import {ServiceService} from "../service.service";

@Component({
  selector: 'app-create-service',
  templateUrl: './create-service.component.html',
  styleUrls: ['./create-service.component.scss']
})
export class CreateServiceComponent {
  constructor(private router: Router, private serviceService:ServiceService) {};

  errorMsg = "";

  newService = new FormGroup({
    name: new FormControl(''),
    description: new FormControl(''),
    startDateTime: new FormControl(),
    length: new FormControl(),
    capacity: new FormControl(),
    price: new FormControl(),
    cancellationFee: new FormControl(),
    // businessId: history.state.user.email // TODO: email or id
  })

  verifyFields() {
    let errorMsg = "";
    let name = this.newService.value.name
    let description = this.newService.value.description
    let startDateTime = this.newService.value.startDateTime
    let length = this.newService.value.length
    let capacity = this.newService.value.capacity
    let pricePerUnit = this.newService.value.price
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
      let res = this.serviceService.addService(this.newService);
      console.log("AFTER result: ", res)
      this.router.navigateByUrl('/home', {state: {user: history.state.user}});
    }
  }

  routeToHome() {
    this.router.navigate(['/']);
  }

  routeToDash() {
    this.router.navigateByUrl('/home', {state: {user: history.state.user}});
  }
}
