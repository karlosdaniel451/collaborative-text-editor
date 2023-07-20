import { Injectable } from '@angular/core';
import {Observable} from "rxjs";
import {Document} from "../model/document.model";
import {HttpClient} from "@angular/common/http";

@Injectable({
  providedIn: 'root'
})
export class DocumentService {

  urlBase = '/documents';
  constructor(private httpClient: HttpClient) { }

  getDocuments(): Observable<Document[]> {
    return this.httpClient.get<Document[]>(this.urlBase);
  }

  getById(id: number): Observable<Document> {
    const url = this.urlBase + '/' + id;
    return this.httpClient.get<Document>(url);
  }
}
