import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Appointment } from './appointment';
import {User} from './user';
import { ServiceAppointment } from './service-appointment';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor(private http: HttpClient) { }

  private apiUrl = 'http://localhost:8080/register';
  private getUserURL = 'http://localhost:8080/user/';
  private apptUrl = 'http://localhost:8080/appointment';

  addUser(firstName: string, lastName: string, email: string, password: string, accountType: string) : Promise<User>{
    return this.http.post<User>(this.apiUrl, {
        first_name: firstName, last_name: lastName, email: email, password: password, account_type: accountType
    }).toPromise().then();

  }

  //to access user obj in other code (console.log example)
  // call login.then( (user) => { console.log(user); });
  login(email: string, password: string) : Promise<void | User> {
    return this.http.post<User>('http://localhost:8080/login', {
      email, password
    }).toPromise().then();
  }

  getUser(id: string, password: string) : Promise<User>{
    return this.http.get<User>(this.getUserURL+id, {

    }).toPromise().then();
  }

  addService(service_id: string, user_id:string) : Promise<Appointment>
  {
    return this.http.post<Appointment>(this.apptUrl, {
      service_id: service_id, user_id: user_id
    }).toPromise().then();
  }

  /*
  cancelAppointment(appointment_id: string) : Promise<Appointment>
  {
    return 
  }*/

  getUserServices(user_id: string) : Promise<ServiceAppointment[]>
  {
    return this.http.get<ServiceAppointment[]>(this.getUserURL+user_id+'/service-appointments').toPromise().then();
  }

}
