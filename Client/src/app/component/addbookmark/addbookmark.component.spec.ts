import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { AddbookmarkComponent } from './addbookmark.component';

describe('AddbookmarkComponent', () => {
  let component: AddbookmarkComponent;
  let fixture: ComponentFixture<AddbookmarkComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AddbookmarkComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AddbookmarkComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
