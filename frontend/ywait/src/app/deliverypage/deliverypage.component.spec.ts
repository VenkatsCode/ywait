import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { DeliverypageComponent } from './deliverypage.component';

describe('DeliverypageComponent', () => {
  let component: DeliverypageComponent;
  let fixture: ComponentFixture<DeliverypageComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ DeliverypageComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(DeliverypageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
