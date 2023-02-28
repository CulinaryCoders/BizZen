import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import {User} from './user';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor(private http: HttpClient) { }

  private apiUrl = 'http://localhost:8080/register';

  addUser(userId: string, username: string, password: string, accountType:string) : Promise<User>{

    return this.http.post<User>(this.apiUrl, {
        username, password, accountType
    }).toPromise().then();

  }

}
