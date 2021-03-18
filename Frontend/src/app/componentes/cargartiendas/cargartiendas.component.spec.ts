import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CargartiendasComponent } from './cargartiendas.component';

describe('CargartiendasComponent', () => {
  let component: CargartiendasComponent;
  let fixture: ComponentFixture<CargartiendasComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CargartiendasComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CargartiendasComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
