import { ComponentFixture, TestBed } from '@angular/core/testing';
import { SignupComponent } from './signup.component';
import { By } from '@angular/platform-browser';

describe('SignupComponent', () => {
  let component: SignupComponent;
  let fixture: ComponentFixture<SignupComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [SignupComponent],
    });
    fixture = TestBed.createComponent(SignupComponent);
    component = fixture.componentInstance;
    component.ngOnInit();
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('form invalid when empty', () => {
    component.signupForm.get('usernameValid')?.setValue('');
    component.signupForm.get('passwordValid')?.setValue('');
    component.signupForm.get('emailValid')?.setValue('');
    expect(component.signupForm.valid).toBeFalsy();
  });

  it('form valid when filled', () => {
    component.signupForm.get('usernameValid')?.setValue('test');
    component.signupForm.get('passwordValid')?.setValue('password123');
    component.signupForm.get('emailValid')?.setValue('test@example.com');
    expect(component.signupForm.valid).toBeTruthy();
  });

  it('form invalid when email is invalid', () => {
    component.signupForm.get('usernameValid')?.setValue('test');
    component.signupForm.get('passwordValid')?.setValue('password123');
    component.signupForm.get('emailValid')?.setValue('test');
    expect(component.signupForm.valid).toBeFalsy();
  });

  it('form invalid when password is too short', () => {
    component.signupForm.get('usernameValid')?.setValue('test');
    component.signupForm.get('passwordValid')?.setValue('pass');
    component.signupForm.get('emailValid')?.setValue('test@example.com');
    expect(component.signupForm.valid).toBeFalsy();
  });

  it('form invalid when password is missing', () => {
    component.signupForm.get('usernameValid')?.setValue('test');
    component.signupForm.get('passwordValid')?.setValue('');
    component.signupForm.get('emailValid')?.setValue('test@example.com');
    expect(component.signupForm.valid).toBeFalsy();
  });

  it('form invalid when username is missing', () => {
    component.signupForm.get('usernameValid')?.setValue('');
    component.signupForm.get('passwordValid')?.setValue('password123');
    component.signupForm.get('emailValid')?.setValue('test@example.com');
    expect(component.signupForm.valid).toBeFalsy();
  });

  it('form invalid when email is missing', () => {
    component.signupForm.get('usernameValid')?.setValue('test');
    component.signupForm.get('passwordValid')?.setValue('password123');
    component.signupForm.get('emailValid')?.setValue('');
    expect(component.signupForm.valid).toBeFalsy();
  });

  it('form invalid when email and password are missing', () => {
    component.signupForm.get('usernameValid')?.setValue('test');
    component.signupForm.get('passwordValid')?.setValue('');
    component.signupForm.get('emailValid')?.setValue('');
    expect(component.signupForm.valid).toBeFalsy();
  });

  it('form invalid when username and email are missing', () => {
    component.signupForm.get('usernameValid')?.setValue('');
    component.signupForm.get('passwordValid')?.setValue('password123');
    component.signupForm.get('emailValid')?.setValue('');
    expect(component.signupForm.valid).toBeFalsy();
  });

  it('form invalid when username and password are missing', () => {
    component.signupForm.get('usernameValid')?.setValue('');
    component.signupForm.get('passwordValid')?.setValue('');
    component.signupForm.get('emailValid')?.setValue('test@example.com');
    expect(component.signupForm.valid).toBeFalsy();
  });

  it('should call onSubmit method', () => {
    spyOn(component, 'signup');
    let el = fixture.debugElement.query(By.css('button')).nativeElement;
    el.click();
    expect(component.signup).toHaveBeenCalledTimes(1);
  });
});
