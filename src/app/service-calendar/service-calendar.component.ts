import {ChangeDetectionStrategy, Component, Input, OnInit} from '@angular/core';
import {CalendarEvent, CalendarView} from 'angular-calendar';
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
      console.log("services received: ", service)
      this.events.push({
        start: new Date(service.start_date_time),
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
