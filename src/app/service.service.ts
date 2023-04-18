import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import {Service} from './service';
import {FormGroup} from "@angular/forms";
import { User } from './user';

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
      name: service.value.name, desc: service.value.description, start_date_time: new Date(service.value.startDateTime), length: service.value.length, capacity: service.value.capacity, price: service.value.price
    }).toPromise().then();
  }

  getService(ID: string) : Promise<Service> {
    return this.http.get<Service>(this.apiUrl+'/'+ID).toPromise().then();
  }

  // Gets all services in db
  getServices() : Promise<Service[]>{
    return this.http.get<Service[]>(this.getAllServices).toPromise().then();
  }

  //update a particular service id based on given information
  updateService(ID: string, name: string, desc: string, start_date_time: Date, length: number, capacity:number, price:number) : Promise<Service>{
    return this.http.put<Service>(this.apiUrl+'/'+ID, {
      name: name, desc: desc, start_date_time: start_date_time, length: length, capacity: capacity, price: price
    }).toPromise().then();
  }

  //get all users attached to service with id given
  getUsers(ID: string) : Promise<User[]>{
    return this.http.get<User[]>(this.apiUrl+'/'+ID+'/users').toPromise().then();
  }
}
