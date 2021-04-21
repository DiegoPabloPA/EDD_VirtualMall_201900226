import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CargaUsuariosComponent } from './carga-usuarios.component';

describe('CargaUsuariosComponent', () => {
  let component: CargaUsuariosComponent;
  let fixture: ComponentFixture<CargaUsuariosComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CargaUsuariosComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CargaUsuariosComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
