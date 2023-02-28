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
});
