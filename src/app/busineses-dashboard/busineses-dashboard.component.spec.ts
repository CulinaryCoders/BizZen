import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BusinesesDashboardComponent } from './busineses-dashboard.component';

describe('BusinesesDashboardComponent', () => {
  let component: BusinesesDashboardComponent;
  let fixture: ComponentFixture<BusinesesDashboardComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ BusinesesDashboardComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(BusinesesDashboardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('fetches list of services from db', () => {
    expect(component).toBeTruthy();
  });

});
