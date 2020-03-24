import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PopmodelComponent } from './popmodel.component';

describe('PopmodelComponent', () => {
  let component: PopmodelComponent;
  let fixture: ComponentFixture<PopmodelComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PopmodelComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PopmodelComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
