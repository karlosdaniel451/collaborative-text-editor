import { Injectable } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {User} from "../model/user.model";
import {Observable} from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class UserService {

  urlBas = '/api/users'

  constructor(private httpClient: HttpClient) { }

  postUser(user: User): Observable<User>  {
    console.log(this.urlBas)
    return this.httpClient.post<User>(this.urlBas, user)
  }

  getById(id: number): Observable<User> {
    const url = this.urlBas + '/' + id;
    return this.httpClient.get<User>(url);
  }

}
