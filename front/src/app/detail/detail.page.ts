import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, ParamMap } from '@angular/router';
import { HttpClient } from'@angular/common/http';

@Component({
  selector: 'app-detail',
  templateUrl: './detail.page.html',
  styleUrls: ['./detail.page.scss'],
})
export class DetailPage implements OnInit {

  id: number;
  item: { 
    id: number, title: string, time: string, body: string
   } = {
    // id: null, title: null, time: null, body: null
    id: 1, title: 'null', time: '2005-06-17T11:06Z', body: 'null'
  };

  constructor(
    public route: ActivatedRoute,
    public http: HttpClient,
  ) { }

  ngOnInit() {
    this.route.paramMap
      .subscribe((params: ParamMap) => {
        this.id = parseInt(params.get('detailId'), 10);
    });
  }

  // ionViewDidEnter() {
  //   this.http.get<{id: number, title: string, time: string, body: string}>('')
  //     .subscribe(data => {
  //       this.item = data;
  //     })
  // }

}
