import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LoginComponent } from './login.component';
import { By } from '@angular/platform-browser';

describe('LoginComponent', () => {
  let component: LoginComponent;
  let fixture: ComponentFixture<LoginComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [LoginComponent]
    });
    fixture = TestBed.createComponent(LoginComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('form invalid when empty', () => {
    component.loginForm.get('usernameValid')?.setValue('');
    component.loginForm.get('passwordValid')?.setValue('');
    expect(component.loginForm.valid).toBeFalsy();
  });

  it('form valid when filled', () => {
    component.loginForm.get('usernameValid')?.setValue('test');
    component.loginForm.get('passwordValid')?.setValue('password123');
    expect(component.loginForm.valid).toBeTruthy();
  });

  it('form invalid when password is too short', () => {
    component.loginForm.get('usernameValid')?.setValue('test');
    component.loginForm.get('passwordValid')?.setValue('pass');
    expect(component.loginForm.valid).toBeFalsy();
  });

  it('form invalid when password is missing', () => {
    component.loginForm.get('usernameValid')?.setValue('test');
    component.loginForm.get('passwordValid')?.setValue('');
    expect(component.loginForm.valid).toBeFalsy();
  });

  it('form invalid when username is missing', () => {
    component.loginForm.get('usernameValid')?.setValue('');
    component.loginForm.get('passwordValid')?.setValue('password123');
    expect(component.loginForm.valid).toBeFalsy();
  });

  it('form invalid when username and password are missing', () => {
    component.loginForm.get('usernameValid')?.setValue('');
    component.loginForm.get('passwordValid')?.setValue('');
    expect(component.loginForm.valid).toBeFalsy();
  });

  it('should call onSubmit method', () => {
    spyOn(component, 'login');
    let el = fixture.debugElement.query(By.css('button')).nativeElement;
    el.click();
    expect(component.login).toHaveBeenCalledTimes(1);
  });
});
