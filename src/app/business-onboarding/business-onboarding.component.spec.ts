import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BusinessOnboardingComponent } from './business-onboarding.component';

describe('BusinessOnboardingComponent', () => {
  let component: BusinessOnboardingComponent;
  let fixture: ComponentFixture<BusinessOnboardingComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ BusinessOnboardingComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(BusinessOnboardingComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
