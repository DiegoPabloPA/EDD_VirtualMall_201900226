import { TestBed } from '@angular/core/testing';

import { UsuariosmasivoService } from './usuariosmasivo.service';

describe('UsuariosmasivoService', () => {
  let service: UsuariosmasivoService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(UsuariosmasivoService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
