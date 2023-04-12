import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { User } from '../user';
import { ServicePageComponent } from './service-page.component';
import { RouterTestingModule } from '@angular/router/testing';
import { Service } from '../service';
import { ServiceOffering } from '../service-offering';

describe('ServicePageComponent', () => {
  let component: ServicePageComponent;
  let fixture: ComponentFixture<ServicePageComponent>;
  let router : Router;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ServicePageComponent ],
      imports: [ RouterTestingModule]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ServicePageComponent);
    component = fixture.componentInstance;
    router = TestBed.inject(Router);

    fixture.detectChanges();

  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should navigate to find classes', () => {

    const navigateSpy = spyOn(router, 'navigateByUrl');

    component.routeToFindClass();
    expect(navigateSpy).toHaveBeenCalledWith('find-classes', {state:{user:component.currentUser}});

  });
  
  it('should join class', () => {

    component.currentUser = new User("firstname", "lastname", "email", "pass", "User", 
      [new Service("123", "Test Service", "Test Service description", new Date("4/11/2023 11:00:00"), 120, 10, 15)]);

    component.joinClass();
    expect(component.userJoined).toBeTruthy();
    
    //checks that the class was added to user
    let index:number = component.currentUser.classes.findIndex((findService) => component.service.serviceId == findService.serviceId);
    expect(index).not.toBe(-1);
  });

  it('should leave class after joining', () => {

    component.currentUser = new User("firstname", "lastname", "email", "pass", "User", 
      [new Service("123", "Test Service", "Test Service description", new Date("4/11/2023 11:00:00"), 120, 10, 15)]);


    component.joinClass();
    component.leaveClass();

    let index:number = component.currentUser.classes.findIndex((findService) => component.service.serviceId == findService.serviceId);

    expect(component.userJoined).toBeFalsy();
    expect(index).toBe(-1);


  });

  
});
