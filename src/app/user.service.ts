import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import {User} from './user';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor(private http: HttpClient) { }

  private apiUrl = '/server';

  addUser(userId: string, username: string, password: string, isBusiness:boolean) : Promise<User>{
    
    return this.http.post<User>(this.apiUrl, {
        username, password, isBusiness
    }).toPromise().then();

  }

}
