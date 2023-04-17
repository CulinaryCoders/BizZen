import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Router } from '@angular/router';
import { RouterTestingModule } from '@angular/router/testing';
import { ProfileComponent } from './profile.component';
import { NavbarComponent } from '../navbar/navbar.component';
import { User } from '../user';

describe('ProfileComponent', () => {
  let component: ProfileComponent;
  let fixture: ComponentFixture<ProfileComponent>;
  let router : Router;
  let testUser = new User("","first name","","","","", []);

  beforeEach(async () => {
    window.history.pushState({user:testUser}, '');
    await TestBed.configureTestingModule({
      declarations: [ ProfileComponent, NavbarComponent ],
      
      imports: [
        RouterTestingModule
      ],
      
    })
    .compileComponents();

    router = TestBed.inject(Router);

    fixture = TestBed.createComponent(ProfileComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();


  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should navigate to home', () => {

    const navigateSpy = spyOn(router, 'navigate');

    component.routeToHome();
    expect(navigateSpy).toHaveBeenCalledWith(['/']);


  });
  
  it('should navigate to classes', () => {

    const navigateSpy = spyOn(router, 'navigateByUrl');

    component.routeToClasses();
    expect(navigateSpy).toHaveBeenCalledWith('/home', {state: {user: history.state.user}});


  });

  it('should navigate to view appointments', () => {

    const navigateSpy = spyOn(router, 'navigateByUrl');

    component.routeToAppointments();
    expect(navigateSpy).toHaveBeenCalledWith('/view-appointments', {state: {user: history.state.user}});


  });

  it('should give the testUser\'s first name', () => {

    component.ngOnInit();
    expect(component.userIdParameter).toBe("first name");

  });

});
