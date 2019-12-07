import { Component, OnInit } from '@angular/core';
import { ModalController } from '@ionic/angular';

@Component({
  selector: 'app-post',
  templateUrl: './post.page.html',
  styleUrls: ['./post.page.scss'],
})
export class PostPage implements OnInit {

  postData: {
    title: string;
    time: string,
    body: string
  } = {title: null, time: null, body: null};

  constructor(
    public modal: ModalController,
  ) { }

  ngOnInit() {
  }

  modalDismiss() {
    this.modal.dismiss();
  }

  onPostClick() {
    console.log(this.postData)
  }

}
