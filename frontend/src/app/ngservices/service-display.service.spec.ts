import { TestBed } from '@angular/core/testing';

import { ServiceDisplayService } from './service-display.service';

describe('ServiceDisplayService', () => {
  let service: ServiceDisplayService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ServiceDisplayService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
