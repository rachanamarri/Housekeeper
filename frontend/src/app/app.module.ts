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
import { AboutusComponent } from './aboutus/aboutus.component';
import { HttpClientModule } from '@angular/common/http';
import { RegisterComponent } from './register/register.component';
import { ProviderBoardComponent } from './provider-board/provider-board.component';
import { FooterComponent } from './footer/footer.component';


const routes: Routes = [
  {path:'', component: ServiceviewComponent},
  {path: 'service/:id', component: ServiceComponent },
  { path: 'login', component: LoginComponent },
  { path: 'register', component: RegisterComponent },
  { path: 'aboutUs', component: AboutusComponent }
];

@NgModule({
  declarations: [
    AppComponent,
    ServiceComponent,
    ServiceviewComponent,
    LoginComponent,
    AboutusComponent,
    RegisterComponent,
    ProviderBoardComponent,
    FooterComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    MatSliderModule,
    MatToolbarModule,
    RouterModule.forRoot(routes),
    FormsModule,
    HttpClientModule,
    
    
  ],
  providers: [],
  bootstrap: [AppComponent],

  exports: [RouterModule]
  
})
export class AppModule { }
