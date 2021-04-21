import { TestBed } from '@angular/core/testing';

import { DatosgrafoService } from './datosgrafo.service';

describe('DatosgrafoService', () => {
  let service: DatosgrafoService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(DatosgrafoService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
