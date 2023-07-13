import {Component, OnInit} from '@angular/core';
import {DocumentService} from "../../service/document.service";
import {Document} from "../../model/document.model";
import {EditingSession} from "../../model/editing-session.model";
import {EditingSessionService} from "../../service/editing-session.service";
import {interval, take} from "rxjs";

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

  constructor(private documentService: DocumentService, private editingSessionService: EditingSessionService) {}

  ngOnInit(): void {

    this.documentService.getDocuments().subscribe(
      (data: Document[]) => {
        this.document = data[0];
      }
    )

    this.editingSessionService.getEditingSessions().subscribe(
      (data: EditingSession[]) => {
        this.editingSession = data[0];
      }
    )
  }

  atualizaPosicaoCorrent(event: MouseEvent) {
    const textarea = event.target as HTMLTextAreaElement;
    this.editingSession.current_position = textarea.selectionStart;
    this.editingSessionService.putEditingSession(this.editingSession).subscribe();
  }

  atualizaConteudo(event: KeyboardEvent) {
    const letraDigitada = event.key;
    this.editingSessionService.postEditingSession(this.editingSession, letraDigitada).subscribe();
  }
}
