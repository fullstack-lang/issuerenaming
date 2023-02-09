import { TestBed } from '@angular/core/testing';

import { IssuerenamingspecificService } from './issuerenamingspecific.service';

describe('IssuerenamingspecificService', () => {
  let service: IssuerenamingspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(IssuerenamingspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
