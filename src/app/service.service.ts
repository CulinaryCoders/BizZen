import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import {Service} from './service';
import {FormGroup} from "@angular/forms";

@Injectable({
  providedIn: 'root'
})
export class ServiceService {
  constructor(private http: HttpClient) { }

  private apiUrl = 'http://localhost:8080/service';
  private getAllServices = 'http://localhost:8080/services';

  // Adds service to DB with specified properties
  addService(service: FormGroup) : Promise<Service>{
    return this.http.post<Service>(this.apiUrl, {
      name: service.value.name, desc: service.value.description, start_date_time: service.value.start_date_time, length: service.value.length, capacity: service.value.capacity, price: service.value.price
    }).toPromise().then();
  }

  // Gets all services in db
  getServices() : Promise<Service[]>{
    return this.http.get<Service[]>(this.getAllServices).toPromise().then();
  }
}
