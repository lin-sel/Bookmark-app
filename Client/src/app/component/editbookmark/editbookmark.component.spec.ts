import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { EditbookmarkComponent } from './editbookmark.component';

describe('EditbookmarkComponent', () => {
  let component: EditbookmarkComponent;
  let fixture: ComponentFixture<EditbookmarkComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ EditbookmarkComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(EditbookmarkComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
