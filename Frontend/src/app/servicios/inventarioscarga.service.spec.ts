import { TestBed } from '@angular/core/testing';

import { InventarioscargaService } from './inventarioscarga.service';

describe('InventarioscargaService', () => {
  let service: InventarioscargaService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(InventarioscargaService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
