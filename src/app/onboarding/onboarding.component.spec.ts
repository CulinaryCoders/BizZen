import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Router } from '@angular/router';
import { RouterTestingModule } from '@angular/router/testing';
import { FormsModule } from '@angular/forms';
import { ReactiveFormsModule } from '@angular/forms';
import { OnboardingComponent } from './onboarding.component';

describe('OnboardingComponent', () => {
  let component: OnboardingComponent;
  let fixture: ComponentFixture<OnboardingComponent>;
  let router : Router;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ OnboardingComponent ],

      imports: [
        RouterTestingModule,
        FormsModule,
        ReactiveFormsModule
      ]

    })
    .compileComponents();

    router = TestBed.inject(Router);

    fixture = TestBed.createComponent(OnboardingComponent);
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

  it('toggles whether interest is selected', () => {
    const interestToAdd = "Technology";
    component.toggleInterest(interestToAdd);
    expect(component.selectedInterests.includes(interestToAdd)).toBeTruthy();
    component.toggleInterest(interestToAdd);
    expect(component.selectedInterests.includes(interestToAdd)).toBeFalse();
  });
});
