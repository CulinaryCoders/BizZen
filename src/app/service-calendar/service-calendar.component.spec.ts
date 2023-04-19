import {ComponentFixture, TestBed} from '@angular/core/testing';

import {ServiceCalendarComponent} from './service-calendar.component';
import {CalendarView} from 'angular-calendar';
import { HttpClientTestingModule } from '@angular/common/http/testing'

describe('ServiceCalendarComponent', () => {
  let component: ServiceCalendarComponent;
  let fixture: ComponentFixture<ServiceCalendarComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ServiceCalendarComponent ],
      imports: [HttpClientTestingModule, CalendarView]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ServiceCalendarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('Should set calendar view to given size', () => {
    component.setView(CalendarView.Month);
    expect(component.view).toBe(CalendarView.Month);

    component.setView(CalendarView.Week);
    expect(component.view).toBe(CalendarView.Week);

    component.setView(CalendarView.Day);
    expect(component.view).toBe(CalendarView.Day);
  });
});
