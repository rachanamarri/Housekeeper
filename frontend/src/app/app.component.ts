import { Component } from '@angular/core';
import { SharedService } from './ngservices/shared-service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  
})
export class AppComponent {
  title = 'frontend';

  // isLoggedIn:boolean=false;

  constructor(private sharedService: SharedService) { }

  // ngOnInit() {
  //   this.sharedService.sharedData.subscribe(isLoggedIn => this.isLoggedIn = isLoggedIn)
  // }

  // test() {
  //   console.log('in test');
  //   console.log(this.sharedService.getOption());
  //   let t = false;
  //    this.sharedService.sharedData.subscribe(isLoggedIn => {
  //     t = isLoggedIn;
  //   })
  //   return t;
  //   // return this.isLoggedIn;
  // }

  get isLoggedIn() {
    console.log('in getter -----', this.sharedService.getOption());
    return this.sharedService.getOption().isLogged;
  }

  logOut() {
    this.sharedService.setOption('isLogged', false);
  }
}
