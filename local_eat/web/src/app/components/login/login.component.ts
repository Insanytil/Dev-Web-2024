import { Component } from '@angular/core';
import { HttpClient } from "@angular/common/http";
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent {
  username: string = '';
  password: string = '';
  errorMessage: string = '';

  constructor(private http: HttpClient, private router: Router) { }
  async login() {
    const userData = {
      "username": this.username,
      "password": this.password,
    };

    const requestOptions = {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(userData)
    };

    try {
      const response = await fetch('http://localhost:8080/api/auth/login', requestOptions);
      const data = await response.json();
      if (response.ok) {
        const token = data.accessToken; 
        console.log(data);
        console.log(token);
        localStorage.setItem('token', token);
        await this.router.navigate(['/']);
      } else {
        console.error('Error:', data.error);
      }
    } catch (error) {
      console.error('Error during request:', error);
    }
  }
}