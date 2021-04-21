import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SubidagrafoComponent } from './subidagrafo.component';

describe('SubidagrafoComponent', () => {
  let component: SubidagrafoComponent;
  let fixture: ComponentFixture<SubidagrafoComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SubidagrafoComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(SubidagrafoComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
