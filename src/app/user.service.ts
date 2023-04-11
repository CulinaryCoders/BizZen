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

  addUser(userId: string, email: string, password: string, accountType:string) : Promise<User>{
    return this.http.post<User>(this.apiUrl, {
        email, password, accountType
    }).toPromise().then();

  }

  getUser(email: string, password: string) : Promise<User>{
    return this.http.get<User>(this.getUserURL+email, {

    }).toPromise().then();
  }

}
