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
});
