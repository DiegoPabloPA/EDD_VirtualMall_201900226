import { TestBed } from '@angular/core/testing';

import { InformesPedidoService } from './informes-pedido.service';

describe('InformesPedidoService', () => {
  let service: InformesPedidoService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(InformesPedidoService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
