import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {WriteComponent} from "./components/write/write.component";
import {CadastraUserComponent} from "./components/cadastra-user/cadastra-user.component";

const routes: Routes = [
  {
    path: '',
    redirectTo: 'user',
    pathMatch: 'full'
  },
  {
    path: 'user',
    component: CadastraUserComponent
  },
  {
    path: 'write/:id',
    component: WriteComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
