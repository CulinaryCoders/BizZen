import { ComponentFixture, TestBed } from '@angular/core/testing';
import { FormsModule, ReactiveFormsModule } from "@angular/forms";
import { BusinessOnboardingComponent } from './business-onboarding.component';

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

  it('changes which business tag is selected', () => {
    const tagToAdd = "Technology";
    component.toggleTags(tagToAdd);
    expect(component.selectedTags.includes(tagToAdd)).toBeTruthy();
    component.toggleTags(tagToAdd);
    expect(component.selectedTags.includes(tagToAdd)).toBeFalse();
  });

  it('checks that the stores opening time is before closing', () => {
    component.onboardingForm.value.openingTime = "11:12";
    component.onboardingForm.value.closingTime = "12:12";
    expect(component.validStartEndTime()).toBeTruthy();
  });

  it('returns error if closing time is before opening', () => {
    component.onboardingForm.value.openingTime = "13:12";
    component.onboardingForm.value.closingTime = "11:12";
    expect(component.validStartEndTime()).toBeFalsy();
  });
});
