import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.scss']
})
export class SignupComponent {
  username: string = '';
  password: string = '';
  email: string = '';
  errorMessage: string = '';
  constructor(private http: HttpClient, private router: Router) { }
  async signup() {
    const userData = {
      "email": this.email,
      "password": this.password,
      "username": this.username
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
      const response = await fetch('http://localhost:8080/api/auth/signup', requestOptions);
      const data = await response.json();
      if (response) {
        await this.router.navigate(['/login']);
      }
    } catch (error) {
      console.error('Error during request:', error);
    }
  }

}