import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-find-classes',
  templateUrl: './find-classes.component.html',
  styleUrls: ['./find-classes.component.scss']
})
export class FindClassesComponent {

    testArray:String[] = ["A", "B", "C", "D", "E"];

  constructor(private router:Router){}

  routeToService()
  {
    this.router.navigate(['class-summary']);

  }
  routeToUserPage()
  {
    this.router.navigate(['profile']);

  }
}
