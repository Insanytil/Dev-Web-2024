import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RegisterProducersComponent } from './register-producers.component';

describe('RegisterProducersComponent', () => {
  let component: RegisterProducersComponent;
  let fixture: ComponentFixture<RegisterProducersComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [RegisterProducersComponent]
    });
    fixture = TestBed.createComponent(RegisterProducersComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
