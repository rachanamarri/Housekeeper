import { Injectable } from '@angular/core';
import { serviceDisplay } from '../models/serviceDisplay.models';

import { Services } from './mockdata';

@Injectable({
  providedIn: 'root'
})
export class ServiceDisplayService {

  constructor() { }

  getServicesData(): serviceDisplay[]{
    return Services;
  }
     
}
