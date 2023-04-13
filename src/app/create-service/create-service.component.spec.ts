import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateServiceComponent } from './create-service.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

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

  it('verifies that all fields are entered', () => {
    component.newService.value.name = "test";
    component.newService.value.description = "test descr";
    component.newService.value.startTime = "11:10";
    component.newService.value.endTime = "12:10";
    component.newService.value.numParticipants = 5;
    component.newService.value.price = 15;

    const allFilled = component.verifyFields();
    // Returns error message, if empty string, no errors
    expect(allFilled).toBe("");
  });

  it('adds to error message when not all fields filled in', () => {
    component.newService.value.name = "";
    component.newService.value.description = "";
    component.newService.value.startTime = "11:12";
    component.newService.value.endTime = "12:11";

    const allFilled = component.verifyFields();
    expect(allFilled).not.toBe("");
  });

  it('checks that the specified start is before the end', () => {
    component.newService.value.startTime = "11:12";
    component.newService.value.endTime = "12:12";
    expect(component.validStartEndTime()).toBeTruthy();
  });

  it('returns error if end time is before start', () => {
    component.newService.value.startTime = "13:12";
    component.newService.value.endTime = "11:12";
    expect(component.validStartEndTime()).toBeFalsy();
  });
});
