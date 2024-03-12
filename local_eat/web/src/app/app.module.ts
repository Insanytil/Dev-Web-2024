import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { RouterModule, Routes } from '@angular/router'
import { HttpClientModule } from '@angular/common/http';
import { AppComponent } from './app.component';
import { ProducersComponent } from './components/producers/producers.component';
import { NotFoundComponent } from './components/not-found/not-found.component';

const routes: Routes = [
  { path: '', component: ProducersComponent },
  { path: '**', component: NotFoundComponent}
];

@NgModule({
  declarations: [
    AppComponent,
    ProducersComponent,
    NotFoundComponent,
  ],
  imports: [
    BrowserModule,
    RouterModule.forRoot(routes),
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
