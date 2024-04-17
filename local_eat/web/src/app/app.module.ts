import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { RouterModule, Routes } from '@angular/router'
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { AppComponent } from './app.component';
import { ProducersComponent } from './components/producers/producers.component';
import { NotFoundComponent } from './components/not-found/not-found.component';
import { LoginComponent } from './components/login/login.component';
import { SignupComponent } from './components/signup/signup.component';
import { NavComponent} from "./components/nav/nav.component";
import { CatchphraseComponent } from './components/catchphrase/catchphrase.component';
import { FormsModule } from '@angular/forms';
import { CustomInterceptor } from './interceptors/custom-interceptor';

const routes: Routes = [
  { path: '', redirectTo: 'home', pathMatch: 'full'},
  { path: 'home', component: NavComponent },
  { path: 'catchphrase', component: CatchphraseComponent},
  { path: 'producers', component: ProducersComponent },
  { path: 'login', component: LoginComponent},
  { path: 'signup', component: SignupComponent},
  { path: '**', component: NotFoundComponent},
];

@NgModule({
  declarations: [
    AppComponent,
    ProducersComponent,
    NotFoundComponent,
    LoginComponent,
    SignupComponent,
    CatchphraseComponent,
    NavComponent
  ],
  imports: [
    BrowserModule,
    RouterModule.forRoot(routes),
    HttpClientModule,
    FormsModule
  ],
  providers: [
    {provide: HTTP_INTERCEPTORS, useClass: CustomInterceptor, multi: true}
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
