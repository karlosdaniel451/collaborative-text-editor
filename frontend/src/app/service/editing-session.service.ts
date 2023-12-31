import { Injectable } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";
import {EditingSession} from "../model/editing-session.model";

@Injectable({
  providedIn: 'root'
})
export class EditingSessionService {

  urlBase = '/api/editing-sessions';
  constructor(private httpClient: HttpClient) { }

  getEditingSessions(): Observable<EditingSession[]> {
    return this.httpClient.get<EditingSession[]>(this.urlBase);
  }

  putEditingSession(editingSession: EditingSession): Observable<EditingSession> {
    const url = this.urlBase + '/' + editingSession.user_id + '/' + editingSession.document_id;
    return this.httpClient.put<EditingSession>(url, editingSession);
  }

  postEditingSession(editingSession: EditingSession, novoTexto: string): Observable<EditingSession> {
    const url = this.urlBase + '/' + editingSession.user_id + '/' + editingSession.document_id;
    return this.httpClient.post<EditingSession>(url, novoTexto);
  }

  postEditingSession2(editingSession: EditingSession): Observable<EditingSession> {
    return this.httpClient.post<EditingSession>(this.urlBase, editingSession);
  }
  deleteEditingSession(editingSession: EditingSession): Observable<EditingSession> {
    const url = this.urlBase + '/' + editingSession.user_id + '/' + editingSession.document_id + '/' + 1;
    return this.httpClient.delete<EditingSession>(url);
  }
}
