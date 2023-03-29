import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BusinessOnboardingComponent } from './business-onboarding.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

describe('BusinessOnboardingComponent', () => {
  let component: BusinessOnboardingComponent;
  let fixture: ComponentFixture<BusinessOnboardingComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ BusinessOnboardingComponent ],
      imports: [
        FormsModule,
        ReactiveFormsModule
      ]
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
