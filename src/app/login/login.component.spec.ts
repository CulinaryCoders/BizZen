import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RouterTestingModule } from '@angular/router/testing';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { ActivatedRoute, Router } from '@angular/router';

import { FormsModule } from '@angular/forms';
import { LoginComponent } from './login.component';
import { routes } from '../app-routing.module';
import { DebugElement } from '@angular/core';

describe('LoginComponent', () => {
  let component: LoginComponent;
  let fixture: ComponentFixture<LoginComponent>;
  let router : Router;
  let submitButtonElement : DebugElement;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LoginComponent ],

      imports: [
        RouterTestingModule.withRoutes(routes),
        HttpClientTestingModule,
        FormsModule
      ],
    })
    .compileComponents();

    //set up router
    router = TestBed.inject(Router);

    //main parts of component
    fixture = TestBed.createComponent(LoginComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();

    //html elements
    submitButtonElement = fixture.debugElement;
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should navigate to register', () => {

      const navigateSpy = spyOn(router, 'navigate');

      component.routeToRegister();
      expect(navigateSpy).toHaveBeenCalledWith(['/register']);
      
      expect(navigateSpy).not.toHaveBeenCalledWith(['/profile']);
      expect(navigateSpy).not.toHaveBeenCalledWith(['/']);

  });

  it('should navigate to home', () => {

      const navigateSpy = spyOn(router, 'navigate');

      component.routeToHome();
      expect(navigateSpy).toHaveBeenCalledWith(['/']);


  });

  it('should navigate to profile', () => {

      const navigateSpy = spyOn(router, 'navigateByUrl');

      component.onSubmit();

      expect(navigateSpy).toHaveBeenCalledWith('/profile', {state: {user:component.model}});

  });

});
