import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CargainventariosComponent } from './cargainventarios.component';

describe('CargainventariosComponent', () => {
  let component: CargainventariosComponent;
  let fixture: ComponentFixture<CargainventariosComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CargainventariosComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CargainventariosComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
