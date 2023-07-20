import { Injectable } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {User} from "../model/user.model";
import {Observable} from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class UserService {

  urlBas = 'http://localhost:8080/users'

  constructor(private httpClient: HttpClient) { }

  postUser(user: User): Observable<User>  {
    return this.httpClient.post<User>(this.urlBas, user)
  }

  getById(id: number): Observable<User> {
    const url = this.urlBas + '/' + id;
    return this.httpClient.get<User>(url);
  }

}
