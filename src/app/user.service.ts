import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import {User} from './user';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor(private http: HttpClient) { }

  private apiUrl = 'http://localhost:8080/register';
  private getUserURL = 'http://localhost:8080/user/';

  addUser(firstName: string, lastName: string, email: string, password: string, accountType: string) : Promise<User>{
    return this.http.post<User>(this.apiUrl, {
        first_name: firstName, last_name: lastName, email: email, password: password, account_type: accountType
    }).toPromise().then();

  }

  //to access user obj in other code (console.log example)
  // call login.then( (user) => { console.log(user); });
  login(username: string, password: string) : Promise<void | User> {
    return this.http.post<User>('http://localhost:8080/login', {
      username, password
    }).toPromise().then(); 
  }
  
  getUser(id: string, password: string) : Promise<User>{
    return this.http.get<User>(this.getUserURL+id, {

    }).toPromise().then();
  }

}
