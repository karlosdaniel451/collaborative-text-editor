import { Component } from '@angular/core';
import {FormBuilder, FormGroup, Validators} from "@angular/forms";
import {UserService} from "../../service/user.service";
import {Router} from "@angular/router";

@Component({
  selector: 'app-cadastra-user',
  templateUrl: './cadastra-user.component.html',
  styleUrls: ['./cadastra-user.component.css']
})
export class CadastraUserComponent {
  cadastroForm: FormGroup

  constructor(private FormBuilder: FormBuilder, private UserService: UserService, private router: Router) {
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
        this.router.navigate(['/write/' + data['id']])
      }
    )
  }

}
