import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ConfirmarcarritoComponent } from './confirmarcarrito.component';

describe('ConfirmarcarritoComponent', () => {
  let component: ConfirmarcarritoComponent;
  let fixture: ComponentFixture<ConfirmarcarritoComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ConfirmarcarritoComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ConfirmarcarritoComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
