import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Router } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { RouterTestingModule } from '@angular/router/testing';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { RegisterComponent } from './register.component';

describe('RegisterComponent', () => {
  let component: RegisterComponent;
  let fixture: ComponentFixture<RegisterComponent>;
  let router: Router;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RegisterComponent ],

      imports: [
        FormsModule,
        RouterTestingModule,
        HttpClientTestingModule,
        ReactiveFormsModule
      ],
    })
    .compileComponents();

    router = TestBed.inject(Router);

    fixture = TestBed.createComponent(RegisterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should navigate to home', () => {

    const navigateSpy = spyOn(router, 'navigate');

    component.routeToHome();
    expect(navigateSpy).toHaveBeenCalledWith(['/']);

  });

  it('should navigate to login', () => {

    const navigateSpy = spyOn(router, 'navigate');

    component.routeToLogin();
    expect(navigateSpy).toHaveBeenCalledWith(['/login']);


  });

  it('should navigate to home', () => {
    const navigateSpy = spyOn(router, 'navigate');

    component.routeToHome();
    expect(navigateSpy).toHaveBeenCalledWith(['/']);
  });

  it('should check that all fields are filled in', () => {
    component.registerForm.value.first_name = "test"
    component.registerForm.value.email = "test@example.com"
    component.registerForm.value.password = "pass123"

    const allFilled = component.allFieldsFilled();
    expect(allFilled).toBe(true);
  });

  it('should catch when not all fields filled in', () => {
    component.registerForm.value.first_name = ""
    component.registerForm.value.email = ""
    component.registerForm.value.password = ""

    const allFilled = component.allFieldsFilled();
    expect(allFilled).toBe(false);
  });

  it('check for matching passwords', () => {
    component.registerForm.value.first_name = "test"
    component.registerForm.value.email = "test@example.com"
    component.registerForm.value.password = "pass123"
    component.confPass = "pass123"

    const passMatch = component.passwordsMatch();
    expect(passMatch).toBe(true);
  });

  it('check for MISmatching passwords', () => {
    component.registerForm.value.first_name = "test"
    component.registerForm.value.email = "test@example.com"
    component.registerForm.value.password = "pass123"
    component.confPass = "pass1234"

    const passMatch = component.passwordsMatch();
    expect(passMatch).toBe(false);
  });
});
