import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FindClassesComponent } from './find-classes.component';

describe('FindClassesComponent', () => {
  let component: FindClassesComponent;
  let fixture: ComponentFixture<FindClassesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ FindClassesComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(FindClassesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
