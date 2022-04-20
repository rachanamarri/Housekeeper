import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';

@Injectable()
export class SharedService {

  private isLoggedIn = new BehaviorSubject(false);
  sharedData = this.isLoggedIn.asObservable();

  constructor() { }

  changeBoolean(bool: boolean) {
    this.isLoggedIn.next(bool)
  }


  private data: any = {};  
  
 setOption(option: any, value: any) { 
     console.log('data set');
     this.data[option] = value;  
     console.log(this.data);   
  }  
  
  getOption() {
    console.log('in service', this.data);
    return this.data;  
  }
  
}