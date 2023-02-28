import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Router } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { RouterTestingModule } from '@angular/router/testing';
import { HttpClientTestingModule } from '@angular/common/http/testing';

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
        HttpClientTestingModule
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
    component.userModel.username = "test"
    component.userModel.userId = "test@example.com"
    component.userModel.password = "pass123"

    const allFilled = component.allFieldsFilled();
    expect(allFilled).toBe(true);
  });

  it('should catch when not all fields filled in', () => {
    component.userModel.username = ""
    component.userModel.userId = ""
    component.userModel.password = ""

    const allFilled = component.allFieldsFilled();
    expect(allFilled).toBe(false);
  });

  it('check for matching passwords', () => {
    component.userModel.username = "test"
    component.userModel.userId = "test@example.com"
    component.userModel.password = "pass123"
    component.confPass = "pass123"

    const passMatch = component.passwordsMatch();
    expect(passMatch).toBe(true);
  });

  it('check for MISmatching passwords', () => {
    component.userModel.username = "test"
    component.userModel.userId = "test@example.com"
    component.userModel.password = "pass123"
    component.confPass = "pass1234"

    const passMatch = component.passwordsMatch();
    expect(passMatch).toBe(false);
  });
});
