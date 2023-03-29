import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { RouterTestingModule } from '@angular/router/testing';
import { FindClassesComponent } from './find-classes.component';
import { ScrollingModule } from '@angular/cdk/scrolling';

describe('FindClassesComponent', () => {
  let component: FindClassesComponent;
  let fixture: ComponentFixture<FindClassesComponent>;
  let router : Router;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ FindClassesComponent ],
      imports: [
        RouterTestingModule,
        ScrollingModule
      ],
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

    const navigateSpy = spyOn(router, 'navigate');

    component.routeToService();
    expect(navigateSpy).toHaveBeenCalledWith(['class-summary']);

  });

  it('should navigate to profile', () => {

    const navigateSpy = spyOn(router, 'navigate');

    component.routeToUserPage();
    expect(navigateSpy).toHaveBeenCalledWith(['profile']);

  });
});
