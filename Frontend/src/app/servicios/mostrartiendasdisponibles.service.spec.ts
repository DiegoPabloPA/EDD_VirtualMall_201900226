import { TestBed } from '@angular/core/testing';

import { MostrartiendasdisponiblesService } from './mostrartiendasdisponibles.service';

describe('MostrartiendasdisponiblesService', () => {
  let service: MostrartiendasdisponiblesService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(MostrartiendasdisponiblesService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
