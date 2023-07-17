import { Component } from '@angular/core';
import {FormBuilder, FormGroup, Validators} from "@angular/forms";
import {UserService} from "../../service/user.service";
import {Router} from "@angular/router";
import {EditingSessionService} from "../../service/editing-session.service";
import {EditingSession} from "../../model/editing-session.model";

@Component({
  selector: 'app-cadastra-user',
  templateUrl: './cadastra-user.component.html',
  styleUrls: ['./cadastra-user.component.css']
})
export class CadastraUserComponent {
  cadastroForm: FormGroup
  editingSession: EditingSession = {
    current_position: 0,
    document_id: 1,
    is_editing_active: true,
    user_id: 0
  }

  constructor(private FormBuilder: FormBuilder, private UserService: UserService, private router: Router, private editingSessionService: EditingSessionService) {
    this.cadastroForm = this.FormBuilder.group({
      user_name: ['', Validators.required],
    })
  }

  onSubmit() {
    if (this.cadastroForm.invalid) {
      return
    }

    this.UserService.postUser(this.cadastroForm.value).subscribe(
      data => {
        this.editingSession.user_id = data['id']
        this.editingSessionService.postEditingSession2(this.editingSession).subscribe()
        this.router.navigate(['/write/' + data['id']])
      }
    )
  }
}
