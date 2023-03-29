import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateServiceComponent } from './create-service.component';
import {FormsModule, ReactiveFormsModule} from "@angular/forms";

describe('CreateServiceComponent', () => {
  let component: CreateServiceComponent;
  let fixture: ComponentFixture<CreateServiceComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CreateServiceComponent ],
      imports: [
        FormsModule,
        ReactiveFormsModule
      ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CreateServiceComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  // it('should check that all fields are filled in', () => {
  //   component.userModel.username = "test"
  //   component.userModel.userId = "test@example.com"
  //   component.userModel.password = "pass123"
  //
  //   const allFilled = component.allFieldsFilled();
  //   expect(allFilled).toBe(true);
  // });
  //
  // it('should catch when not all fields filled in', () => {
  //   component.userModel.username = ""
  //   component.userModel.userId = ""
  //   component.userModel.password = ""
  //
  //   const allFilled = component.allFieldsFilled();
  //   expect(allFilled).toBe(false);
  // });
  //
  // it('checks that the specified start is before the end', () => {
  //   component.toggleInterest(interestToAdd);
  //   expect(component.selectedInterests.includes(interestToAdd)).toBeTruthy();
  //   component.toggleInterest(interestToAdd);
  //   expect(component.selectedInterests.includes(interestToAdd)).toBeFalse();
  // });
});
