import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateServiceComponent } from './create-service.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import {HttpClientTestingModule} from '@angular/common/http/testing'
import { NavbarComponent } from '../navbar/navbar.component';
import { RouterTestingModule } from '@angular/router/testing';
import {User} from "../user";
import {Router} from "@angular/router";

describe('CreateServiceComponent', () => {
  let component: CreateServiceComponent;
  let fixture: ComponentFixture<CreateServiceComponent>;
  let router : Router;

  let testUser = new User("12345","firstname", "lastname", "email", "pass", "Business", []);

  beforeEach(async () => {
    window.history.pushState({user: testUser}, '');

    await TestBed.configureTestingModule({
      declarations: [ CreateServiceComponent, NavbarComponent],
      imports: [
        FormsModule,
        ReactiveFormsModule,
        HttpClientTestingModule,
        RouterTestingModule
      ]
    })
    .compileComponents();

    router = TestBed.inject(Router);
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
    component.newService.value.startDateTime = new Date();
    component.newService.value.length = 12;
    component.newService.value.capacity = 5;
    component.newService.value.price = 15;
    component.newService.value.cancellationFee = 15;

    const allFilled = component.verifyFields();
    // Returns error message, if empty string, no errors
    expect(allFilled).toBe("");
  });

  it('adds to error message when not all fields filled in', () => {
    component.newService.value.name = "";
    component.newService.value.description = "";
    component.newService.value.startDateTime = "11:12";
    component.newService.value.length = "12:11";

    const allFilled = component.verifyFields();
    expect(allFilled).not.toBe("");
  });

  it('Routes to landing', () => {
    const navigateSpy = spyOn(router, 'navigate');
    component.routeToHome();
    expect(navigateSpy).toHaveBeenCalledWith(['/']);

  });

  it('Routes to Business Dashboard', () => {
    const navigateSpy = spyOn(router, 'navigateByUrl');
    component.routeToDash();
    expect(navigateSpy).toHaveBeenCalledWith('/home', {state: {user: history.state.user}});

  });
});
