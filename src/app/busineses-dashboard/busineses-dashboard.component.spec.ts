import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BusinesesDashboardComponent } from './busineses-dashboard.component';
import {HttpClientTestingModule, HttpTestingController} from "@angular/common/http/testing";
import {HttpClient} from "@angular/common/http";
import {Router} from "@angular/router";
import {UserService} from "../user.service";
import {Service} from "../service";
import {User} from "../user";
import {RouterTestingModule} from "@angular/router/testing";
import { NavbarComponent } from '../navbar/navbar.component';

describe('BusinesesDashboardComponent', () => {
  let component: BusinesesDashboardComponent;
  let fixture: ComponentFixture<BusinesesDashboardComponent>;
  let router : Router;
  let userService: UserService;
  let httpClient: HttpClient;
  let httpTestController:HttpTestingController;

  let service:Service = new Service("123", "Test Service", "Test Service desc",
    new Date("4/11/2023 11:00:00"), 120, 10, 15);

  let testUser = new User("12345","firstname", "lastname", "email", "pass", "Business", [service]);

  beforeEach(async () => {
    window.history.pushState({user: testUser, service: service}, '');

    await TestBed.configureTestingModule({
      declarations: [ BusinesesDashboardComponent, NavbarComponent ],
      imports: [RouterTestingModule, HttpClientTestingModule]
    })
    .compileComponents();

    fixture = TestBed.createComponent(BusinesesDashboardComponent);
    component = fixture.componentInstance;
    router = TestBed.inject(Router);
    httpClient = TestBed.inject(HttpClient);
    userService = TestBed.inject(UserService);
    httpTestController = TestBed.inject(HttpTestingController);

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('Completes GET request to fetch list of services from db', () => {
    const req = httpTestController.expectOne("http://localhost:8080/services");
    expect(req.request.method).toEqual("GET");
  });

  it('Filters services by ascending date', () => {
    let dateRange = [new Date("4/13/2023T11:00:00")]
    let services = [
      new Service("123", "Test Service", "Test Service desc",
        new Date("4/13/2023 11:00:00"), 120, 10, 15),
      new Service("123", "Test Service", "Test Service desc",
        new Date("4/12/2023 11:00:00"), 120, 10, 15),
      new Service("123", "Test Service", "Test Service desc",
        new Date("4/11/2023 11:00:00"), 120, 10, 15),
      new Service("123", "Test Service", "Test Service desc",
        new Date("4/15/2023 11:00:00"), 120, 10, 15)
    ];

  });

  it('Should navigate to create service', () => {
    const navigateSpy = spyOn(router, 'navigateByUrl');
    component.openAddService();
    expect(navigateSpy).toHaveBeenCalledWith('/create-service', {state:{user:history.state.user}});
  });

  it('Returns correct Date-Time from start date and duration', () => {
    let startDateTime = new Date();
    let duration = 10;
    let endDate = component.getEndDate(startDateTime, duration);
    let expectedEnd = new Date(startDateTime.getTime() + length*60000);
    expect(endDate).toEqual(expectedEnd);
  });

  it('Should navigate to service info page', () => {
    const navigateSpy = spyOn(router, 'navigateByUrl');
    component.goToServicePage(service);
    expect(navigateSpy).toHaveBeenCalledWith('/class-summary', {state:{user:history.state.user, service: service, prevPage: "/home"}});
  });

  it('Updates the view date range for services to be filtered by', () => {
    component.updateDateRange([new Date("4/11/2023"), new Date("4/12/2023")]);
    expect(component.viewDateRange).toEqual([new Date("4/11/2023"), new Date("4/12/2023")])
  });
});
