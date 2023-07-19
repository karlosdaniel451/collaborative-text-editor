import {Component, OnInit} from '@angular/core';
import {DocumentService} from "../../service/document.service";
import {Document} from "../../model/document.model";
import {EditingSession} from "../../model/editing-session.model";
import {EditingSessionService} from "../../service/editing-session.service";
import {catchError, EMPTY, empty, interval, take} from "rxjs";
import {ActivatedRoute} from "@angular/router";

@Component({
  selector: 'app-write',
  templateUrl: './write.component.html',
  styleUrls: ['./write.component.css']
})
export class WriteComponent implements OnInit {

  document: Document = {
    id: 0,
    name: '',
    content: ''
  }

  editingSession: EditingSession = {
    current_position: 0,
    document_id: 0,
    is_editing_active: false,
    user_id: 0
  }

  constructor(private documentService: DocumentService, private editingSessionService: EditingSessionService,
              private route: ActivatedRoute) {
    const source = interval(4000)

    source.pipe().subscribe(() => {
      this.verificaSeExisteNovoConteudo()
    })
  }

  ngOnInit(): void {
    const id: string = this.route.snapshot.paramMap.get('id') as string;

    this.verificaSeExisteNovoConteudo()

    this.editingSessionService.getEditingSessions().pipe(
      catchError(err => {
        return EMPTY
      })).subscribe(
      (data: EditingSession[]) => {
        data.forEach((editingSession: EditingSession) => {
          if (editingSession.user_id == parseInt(id)) {
            this.editingSession = editingSession;
          }
        })
      }
    )
  }

  verificaSeExisteNovoConteudo(): void {
    this.documentService.getById(1).pipe(
      catchError(err => {
        window.alert('Erro ao buscar documento: ' + err.message)
        return EMPTY
      })
    ).subscribe(
      (data: Document) => {
        if (data.content != this.document.content) {
          this.document = data;
        }
      }
    )
  }

  atualizaPosicaoCorrent(event: MouseEvent | any) {
    const textarea = event.target as HTMLTextAreaElement;
    this.editingSession.current_position = textarea.selectionStart;

    this.editingSessionService.putEditingSession(this.editingSession).pipe(
      catchError(err => {
        return EMPTY
      })
    ).subscribe()
  }

  atualizaConteudo(event: KeyboardEvent) {
    const letraDigitada = event.key;

    if (letraDigitada.length > 1) {
      return;
    }
    this.editingSessionService.postEditingSession(this.editingSession, letraDigitada).pipe(
      catchError(err => {
        window.alert('Erro ao atualizar conteúdo: ' + err.error.detail)
        return EMPTY
      })
    ).subscribe();
  }

  deletaConteudo(event: any) {
    this.editingSessionService.deleteEditingSession(this.editingSession).pipe(
      catchError(err => {
        window.alert('Erro ao deletar conteúdo: ' + err.error.detail)
        return EMPTY
      })
    ).subscribe();
  }
}
