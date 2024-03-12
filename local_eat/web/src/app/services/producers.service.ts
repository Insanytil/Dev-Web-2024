import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable, of } from 'rxjs';
import { Producer } from '../models/producer.model';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class ProducersService {
  private url = `${environment.apiUrl}/api/producers`
  constructor(private http:HttpClient) { }

  getProducers(): Observable<Producer[]>{
    return this.http.get<Producer[]>(this.url);
  }
}
