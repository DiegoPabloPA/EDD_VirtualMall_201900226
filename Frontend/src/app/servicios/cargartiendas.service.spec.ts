import { TestBed } from '@angular/core/testing';

import { CargartiendasService } from './cargartiendas.service';

describe('CargartiendasService', () => {
  let service: CargartiendasService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(CargartiendasService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
