import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { HttpResponse } from '@angular/common/http';
import { AuthService } from '../../services/auth.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  username: string;
  password: string;

  constructor(private authService: AuthService, private router: Router) { }

  login() {
    this.authService.login(this.username, this.password).subscribe(
      (res: HttpResponse<any>) => {
        console.log('response from server:', res);
        console.log('response headers', res.headers.keys());
        if (res.ok) {
          const token = res.body.accessToken;
          localStorage.setItem('token', token);
          this.router.navigate(['/']);
        } else {
          console.error('Error:', res.body.error);
        }
      },
      error => {
        console.error('Error:', error);
      }
    );
  }
}
