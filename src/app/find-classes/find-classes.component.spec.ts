import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { RouterTestingModule } from '@angular/router/testing';
import { FindClassesComponent } from './find-classes.component';
import { ScrollingModule } from '@angular/cdk/scrolling';
import {FormsModule, ReactiveFormsModule} from "@angular/forms";

describe('FindClassesComponent', () => {
  let component: FindClassesComponent;
  let fixture: ComponentFixture<FindClassesComponent>;
  let router : Router;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ FindClassesComponent ],
      imports: [
        FormsModule,
        ReactiveFormsModule,
        ScrollingModule
      ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(FindClassesComponent);
    component = fixture.componentInstance;
    router = TestBed.inject(Router);


    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should navigate to summary', () => {

    const navigateSpy = spyOn(router, 'navigateByUrl');

    component.routeToService(component.testService);
    expect(navigateSpy).toHaveBeenCalledWith('/class-summary', {state:{user:component.user, service:component.testService}});

  });

  it('should navigate to profile', () => {

    const navigateSpy = spyOn(router, 'navigateByUrl');

    component.routeToUserPage();
    expect(navigateSpy).toHaveBeenCalledWith('/profile', {state:{user:component.user}});

  });
});
