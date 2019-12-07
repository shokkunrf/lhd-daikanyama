import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { ModalController } from '@ionic/angular';
import { PostPage } from '../post/post.page';

@Component({
  selector: 'app-home',
  templateUrl: 'home.page.html',
  styleUrls: ['home.page.scss'],
})

export class HomePage {

  items: { id: number, title: string, time: string, body: string }[] = [];

  constructor(
    public http: HttpClient,
    public modal: ModalController,
  ) { }

  ionViewDidEnter() {
    this.items = [
      {
        id: 1,
        title: 'hoge',
        time: '2005-06-17T11:06Z',
        body: 'hgeoghehgoehgoehgoehgogoheogheoghoehehgoehgoehoghoehgoehgehgheohgoehogheo',
      },
      {
        id: 2,
        title: 'hoge',
        time: '2005-06-17T11:06Z',
        body: 'hgeoghehgoehgoehgoehgogoheogheoghoehehgoehgoehoghoehgoehgehgheohgoehogheo',
      },
      {
        id: 3,
        title: 'hoge',
        time: '2005-06-17T11:06Z',
        body: 'hgeoghehgoehgoehgoehgogoheogheoghoehehgoehgoehoghoehgoehgehgheohgoehogheo',
      },
      {
        id: 4,
        title: 'hoge',
        time: '2005-06-17T11:06Z',
        body: 'hgeoghehgoehgoehgoehgogoheogheoghoehehgoehgoehoghoehgoehgehgheohgoehogheo',
      }
    ]
  }

  async postReminder() {
    const modal = await this.modal.create({
      component: PostPage,
    });
    await modal.present();
  }
}
