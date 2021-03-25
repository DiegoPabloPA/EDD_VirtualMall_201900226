import { TestBed } from '@angular/core/testing';

import { PedidosMasivosService } from './pedidos-masivos.service';

describe('PedidosMasivosService', () => {
  let service: PedidosMasivosService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(PedidosMasivosService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
