import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { User } from '../user';
import { ServicePageComponent } from './service-page.component';
import { RouterTestingModule } from '@angular/router/testing';
import { Service } from '../service';
import { ServiceOffering } from '../service-offering';

import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import {HttpClient, HttpResponse} from '@angular/common/http'

import { NavbarComponent } from '../navbar/navbar.component';
import { FormsModule } from '@angular/forms'
import { UserService } from '../user.service';
import { Appointment } from '../appointment';


describe('ServicePageComponent', () => {
  let component: ServicePageComponent;
  let fixture: ComponentFixture<ServicePageComponent>;
  let router : Router;
  let userService: UserService;
  let httpClient: HttpClient;
  let httpTestController:HttpTestingController;

  let service:Service = new Service("123", "Test Service", "Test Service desc", 
    new Date("4/11/2023 11:00:00"), 120, 10, 15);

  let testUser = new User("12345","firstname", "lastname", "email", "pass", "User", [service]);

  beforeEach(async () => {
    window.history.pushState({user: testUser, service: service}, '');

    await TestBed.configureTestingModule({
      declarations: [ ServicePageComponent, NavbarComponent ],
      imports: [ 
        RouterTestingModule, 
        HttpClientTestingModule, 
        FormsModule
      ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ServicePageComponent);
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

  it('should navigate to find classes', () => {

    const navigateSpy = spyOn(router, 'navigateByUrl');

    component.routeToFindClass();
    expect(navigateSpy).toHaveBeenCalledWith('/home', {state:{user:component.currentUser}});

  });
  
  it('should join class', () => {

    component.currentUser = testUser;
    component.service = service;

    component.joinClass();
    expect(component.userJoined).toBeTruthy();

    const req = httpTestController.expectOne("http://localhost:8080/appointment");
    expect(req.request.method).toEqual("POST");
    

  });

  it('should edit', () => {

    component.editService();
    expect(component.isEditing).toBeTruthy();

  });


//TODO: fix
  it('should leave class after joining', () => {

    component.currentUser = testUser;

    component.joinClass();
    component.leaveClass();

    let index:number = component.currentUser.classes.findIndex((findService) => component.service.ID == findService.ID);

    expect(component.userJoined).toBeFalsy();
    expect(index).toBe(-1);


  });

  
});
