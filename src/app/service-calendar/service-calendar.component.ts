import {ChangeDetectionStrategy, Component,} from '@angular/core';
import {CalendarEvent, CalendarView} from 'angular-calendar';
import {startOfDay} from 'date-fns';
import {Router} from "@angular/router";

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
export class ServiceCalendarComponent {
  constructor(private router: Router) {};

  viewDate: Date = new Date();
  view: CalendarView = CalendarView.Month;
  CalendarView = CalendarView;

  setView(view: CalendarView) {
    this.view = view;
  }

  events: CalendarEvent[] = [
    {
      start: startOfDay(new Date()),
      title: 'Yoga Class',
    },
    {
      start: startOfDay(new Date()),
      title: 'Painting Class',
    }
  ];

  dayClicked({ date, events }: { date: Date; events: CalendarEvent[] }): void {
    console.log(date);
    this.viewDate = date;
    this.view = CalendarView.Day;
  }

  eventClicked({ event }: { event: CalendarEvent }): void {
    alert("Navigating to " + event.title + " page");
    this.router.navigate(['/service/'+event.meta.id]);
  }
}
