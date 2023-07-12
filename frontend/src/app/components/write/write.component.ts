import {Component, OnInit} from '@angular/core';
import {DocumentService} from "../../service/document.service";
import {Document} from "../../model/document.model";

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

  constructor(private documentService: DocumentService) {}

  ngOnInit(): void {
    this.documentService.getDocuments().subscribe(
      (data: Document[]) => {
        this.document = data[0];
      }
    )
  }


}
