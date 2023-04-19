import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LandingComponent } from './landing.component';
import {RouterTestingModule} from "@angular/router/testing";
import {Router} from "@angular/router";

describe('LandingComponent', () => {
  let component: LandingComponent;
  let fixture: ComponentFixture<LandingComponent>;
  let router : Router;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LandingComponent ],
      imports: [RouterTestingModule]
    })
    .compileComponents();

    router = TestBed.inject(Router);
    fixture = TestBed.createComponent(LandingComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('Should navigate to Register page', () => {
    const navigateSpy = spyOn(router, 'navigate');

    component.routeToRegister();
    expect(navigateSpy).toHaveBeenCalledWith(['/register']);
  });

  it('Should navigate to Login page', () => {
    const navigateSpy = spyOn(router, 'navigate');

    component.routeToLogin();
    expect(navigateSpy).toHaveBeenCalledWith(['/login']);
  });
});
