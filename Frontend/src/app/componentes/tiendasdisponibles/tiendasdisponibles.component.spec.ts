import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TiendasdisponiblesComponent } from './tiendasdisponibles.component';

describe('TiendasdisponiblesComponent', () => {
  let component: TiendasdisponiblesComponent;
  let fixture: ComponentFixture<TiendasdisponiblesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ TiendasdisponiblesComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(TiendasdisponiblesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
