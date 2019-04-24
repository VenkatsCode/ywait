import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CustomerpageComponent } from './customer-page.component';

describe('CustomerpageComponent', () => {
  let component: CustomerpageComponent;
  let fixture: ComponentFixture<CustomerpageComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CustomerpageComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CustomerpageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
