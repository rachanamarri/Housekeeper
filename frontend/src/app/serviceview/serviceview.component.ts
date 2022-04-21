import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-serviceview',
  templateUrl: './serviceview.component.html',
  styleUrls: ['./serviceview.component.css']
})
export class ServiceviewComponent implements OnInit {

  constructor(private http: HttpClient) { }

  ngOnInit(): void {
  }
  config = {};
  test() {
    console.log('in api');
    return this.http.get<any>("http://localhost:8080/services/:serviceId");
  }

  callApi() {
    this.test()
      .subscribe((data: any) => this.config = {
          heroesUrl: data.heroesUrl,
          textfile:  data.textfile,
          date: data.date,
      });
  }
}
