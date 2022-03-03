import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router, UrlSegment } from '@angular/router';
import { serviceDisplay } from '../models/serviceDisplay.models';
import { Services } from '../ngservices/mockdata';
import { ServiceDisplayService } from '../ngservices/service-display.service';


@Component({
  selector: 'app-service',
  templateUrl: './service.component.html',
  styleUrls: ['./service.component.css'],
  providers:[ServiceDisplayService]
})
export class ServiceComponent implements OnInit {
  data: serviceDisplay[] = [];
  mappingData:any = {
    labels: [],
    datasets: []
  };
   
  serviceId!: { id: number; };
  hey!: number;
  
  red: number | undefined;
  constructor( private dataService:ServiceDisplayService, private router: ActivatedRoute) {
    this.hey = this.router.snapshot.params['id'];

   }

  
  

  ngOnInit(): void {
    
    this.data=this.dataService.getServicesData();
    
}


public get  test() {
  console.log(Services);
  console.log(this.red);
  return Services.find(item => (item.id) == this.hey);
}

getAppointment(){
  alert ("Appointment is Booked.")
}
}
 