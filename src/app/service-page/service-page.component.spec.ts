import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { ServicePageComponent } from './service-page.component';
import { RouterTestingModule } from '@angular/router/testing';

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

    const navigateSpy = spyOn(router, 'navigate');

    component.routeToFindClass();
    expect(navigateSpy).toHaveBeenCalledWith(['find-classes']);

  });
  
  it('should join class', () => {

    component.joinClass();
    expect(component.userJoined).toBeTruthy();

  });

  it('should leave class after joining', () => {

    component.joinClass();
    component.leaveClass();
    expect(component.userJoined).toBeFalsy();

  });

  
});
