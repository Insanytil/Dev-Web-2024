import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from '../../services/auth.service';
import { HttpResponse } from '@angular/common/http';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.scss']
})
export class SignupComponent {
  username: string = '';
  password: string = '';
  email: string = '';
  constructor(private authService: AuthService, private router: Router) { }

  signup() {
    this.authService.signup(this.email, this.password, this.username).subscribe(
      (res: HttpResponse<any>) => {
        console.log('response from server:', res);
        console.log('response headers', res.headers.keys());
        if (res.ok) {
          this.router.navigate(['/login']);
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