import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import {Service} from './service';

@Injectable({
  providedIn: 'root'
})
export class ServiceService {
  constructor(private http: HttpClient) { }

  private apiUrl = 'http://localhost:8080/service';
  private getAllServices = 'http://localhost:8080/services';

  // Adds service to DB with specified properties
  addService(name : string, desc : string, start_date_time: Date, length: number, capacity: number, price: number) : Promise<Service>{
    return this.http.post<Service>(this.apiUrl, {
      name: name, desc: desc, start_date_time: start_date_time, length: length, capacity: capacity, price: price
    }).toPromise().then();

  }
  // Gets all services in db
  getServices() : Promise<Service>{
    return this.http.get<Service>(this.getAllServices).toPromise().then();
  }
}
