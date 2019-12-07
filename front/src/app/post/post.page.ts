import { Component, OnInit } from '@angular/core';
import { ModalController } from '@ionic/angular';

@Component({
  selector: 'app-post',
  templateUrl: './post.page.html',
  styleUrls: ['./post.page.scss'],
})
export class PostPage implements OnInit {

  constructor(
    public modal: ModalController,
  ) { }

  ngOnInit() {
  }

  modalDismiss() {
    console.log('hoge');
    this.modal.dismiss();
  }

}
