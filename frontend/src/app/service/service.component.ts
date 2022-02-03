import { Component, OnInit } from '@angular/core';
import { ServiceDisplayService } from '../ngservices/service-display.service';

@Component({
  selector: 'app-service',
  templateUrl: './service.component.html',
  styleUrls: ['./service.component.css'],
  providers:[ServiceDisplayService]
})
export class ServiceComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }

}
