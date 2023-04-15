import {ChangeDetectionStrategy, Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {
  CalendarEvent,
  CalendarMonthViewBeforeRenderEvent,
  CalendarView,
  CalendarWeekViewBeforeRenderEvent
} from 'angular-calendar';
import {Router} from "@angular/router";
import {User} from "../user";
import {ServiceService} from "../service.service";

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
  constructor(private router: Router, private serviceService: ServiceService) {};
  // @ts-ignore
  @Input() services: any[];
  // @ts-ignore
  @Input() user: User;
  @Output() newDateRangeEvent = new EventEmitter<any[]>(true);

  viewDateRange: any[2] = [
    new Date().setDate(1), // first of the month
    new Date(new Date().getFullYear(), new Date().getMonth()+1, 0) // last day of the month
  ];

  viewDate: Date = new Date();
  monthNames = ["January", "February", "March", "April", "May", "June",
    "July", "August", "September", "October", "November", "December"
  ];
  weekDays = ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"]
  view: CalendarView = CalendarView.Month;
  CalendarView = CalendarView;
  // @ts-ignore
  events: CalendarEvent[] = [];

  ngOnInit(): void {
    this.services.forEach((service) => {
      this.events.push({
        start: new Date(service.start_date_time),
        title: service.name,
        meta: {serviceObj: service},
      })
    });
  }

  setView(view: CalendarView) {
    this.view = view;
  }

  viewMonthChange(e: CalendarMonthViewBeforeRenderEvent) {
    this.viewDateRange = [e.period.start, e.period.end];
    this.newDateRangeEvent.emit(this.viewDateRange);
  }

  viewWeekChange(e: CalendarWeekViewBeforeRenderEvent) {
    this.viewDateRange = [e.period.start, e.period.end];
    this.newDateRangeEvent.emit(this.viewDateRange);
  }

  viewDayChange(e: CalendarWeekViewBeforeRenderEvent) {
    this.viewDateRange = [e.period.start, e.period.end];
    this.newDateRangeEvent.emit(this.viewDateRange);
  }

  dayClicked({ date, events }: { date: Date; events: CalendarEvent[] }): void {
    this.viewDate = date;
    this.view = CalendarView.Day;
  }

  eventClicked({ event }: { event: CalendarEvent }): void {
    this.router.navigateByUrl('/class-summary', {state: {user: history.state.user, service: event.meta.serviceObj}});
  }
}
