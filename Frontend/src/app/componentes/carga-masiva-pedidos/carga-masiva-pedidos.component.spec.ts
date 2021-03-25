import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CargaMasivaPedidosComponent } from './carga-masiva-pedidos.component';

describe('CargaMasivaPedidosComponent', () => {
  let component: CargaMasivaPedidosComponent;
  let fixture: ComponentFixture<CargaMasivaPedidosComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CargaMasivaPedidosComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CargaMasivaPedidosComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
