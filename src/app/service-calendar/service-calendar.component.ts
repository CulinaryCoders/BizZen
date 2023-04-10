import {ChangeDetectionStrategy, Component, Input, OnInit} from '@angular/core';
import {CalendarEvent, CalendarView} from 'angular-calendar';
import {startOfDay} from 'date-fns';
import {Router} from "@angular/router";
import {Service} from "../service";
import {User} from "../user";

@Component({
  selector: 'app-service-calendar-component',
  templateUrl: './service-calendar.component.html',
  styleUrls: ['./service-calendar.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
  styles: [
    `
      h3 {
        margin: 0 0 10px;
      }

      pre {
        background-color: #f5f5f5;
        padding: 15px;
      }
    `,
  ],
})
export class ServiceCalendarComponent implements OnInit{
  constructor(private router: Router) {};
  // @ts-ignore
  @Input() services: any[];
  // @ts-ignore
  @Input() user: User;

  viewDate: Date = new Date();
  view: CalendarView = CalendarView.Month;
  CalendarView = CalendarView;
  // @ts-ignore
  events: CalendarEvent[] = [];

  ngOnInit(): void {
    this.services.forEach((service) => {
      this.events.push({
        start: service.start_date_time,
        title: service.name,
        meta: {serviceObj: service},
      })
    })
  }

  setView(view: CalendarView) {
    this.view = view;
  }

  dayClicked({ date, events }: { date: Date; events: CalendarEvent[] }): void {
    console.log(date);
    this.viewDate = date;
    this.view = CalendarView.Day;
  }

  eventClicked({ event }: { event: CalendarEvent }): void {
    this.router.navigateByUrl('/class-summary', {state: {user: this.user, service: event.meta.serviceObj}});
  }
}
