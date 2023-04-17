import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewAppointmentsComponent } from './view-appointments.component';
import { Router } from '@angular/router';
import { RouterTestingModule } from '@angular/router/testing';
import { Service } from '../service';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { NavbarComponent } from '../navbar/navbar.component';
import { CdkVirtualScrollViewport, ScrollingModule } from '@angular/cdk/scrolling';


describe('ViewAppointmentsComponent', () => {
  let component: ViewAppointmentsComponent;
  let fixture: ComponentFixture<ViewAppointmentsComponent>;
  let router : Router;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ViewAppointmentsComponent, NavbarComponent ],

      imports: [RouterTestingModule, HttpClientTestingModule, ScrollingModule]
    })
    .compileComponents();

    router = TestBed.inject(Router);

    fixture = TestBed.createComponent(ViewAppointmentsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should navigate to profile', () => {

    const navigateSpy = spyOn(router, 'navigateByUrl');

    component.routeToProfile();
    expect(navigateSpy).toHaveBeenCalledWith('/profile', {state: {user: component.user}});


  });

  it('should navigate to a passed in service', () => {

    const navigateSpy = spyOn(router, 'navigateByUrl');
    let testService = new Service("","","", new Date("4/17/2023"), 120,10,15);

    component.routeToService(testService);

    expect(navigateSpy).toHaveBeenCalledWith('/class-summary', {state: {user: component.user, service: testService, prevPage: '/view-appointments'}});


  });

});
