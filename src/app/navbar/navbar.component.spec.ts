import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NavbarComponent } from './navbar.component';
import {RouterTestingModule} from "@angular/router/testing";
import {Router} from "@angular/router";
import {User} from "../user";

describe('NavbarComponent', () => {
  let component: NavbarComponent;
  let fixture: ComponentFixture<NavbarComponent>;
  let router : Router;
  let testUser = new User("12345","firstname", "lastname", "email", "pass", "User", []);

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ NavbarComponent ],
      imports: [RouterTestingModule]
    })
    .compileComponents();

    router = TestBed.inject(Router);
    fixture = TestBed.createComponent(NavbarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('Should navigate to Profile page', () => {
    const navigateSpy = spyOn(router, 'navigateByUrl');

    component.routeToProfile();
    expect(navigateSpy).toHaveBeenCalledWith('/profile', {state: {user: testUser}});
  });

  it('Should navigate to Landing page', () => {
    const navigateSpy = spyOn(router, 'navigate');

    component.routeToHome();
    expect(navigateSpy).toHaveBeenCalledWith(['/']);
  });

  it('Should navigate to Business Dashboard page', () => {
    const navigateSpy = spyOn(router, 'navigateByUrl');

    component.routeToDash();
    expect(navigateSpy).toHaveBeenCalledWith('/home', {state: {user: testUser}});
  });
});
