import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { MatSliderModule } from '@angular/material/slider';
import {MatToolbarModule } from '@angular/material/toolbar'
import { AppComponent } from './app.component';
import { Routes, RouterModule } from '@angular/router'; 
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { ServiceComponent } from './service/service.component';
import { ServiceviewComponent } from './serviceview/serviceview.component';
import { LoginComponent } from './login/login.component';
import { FormsModule } from '@angular/forms';

const routes: Routes = [
  {path:'', component: ServiceviewComponent},
  {path: 'service/:id', component: ServiceComponent },
  { path: 'login', component: LoginComponent }
];

@NgModule({
  declarations: [
    AppComponent,
    ServiceComponent,
    ServiceviewComponent,
    LoginComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    MatSliderModule,
    MatToolbarModule,
    RouterModule.forRoot(routes),
    FormsModule,
    
    
  ],
  providers: [],
  bootstrap: [AppComponent],

  exports: [RouterModule]
  
})
export class AppModule { }
