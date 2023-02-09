import { ComponentFixture, TestBed } from '@angular/core/testing';

import { IssuerenamingspecificComponent } from './issuerenamingspecific.component';

describe('IssuerenamingspecificComponent', () => {
  let component: IssuerenamingspecificComponent;
  let fixture: ComponentFixture<IssuerenamingspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ IssuerenamingspecificComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(IssuerenamingspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
