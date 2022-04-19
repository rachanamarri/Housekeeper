import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ProviderBoardComponent } from './provider-board.component';

describe('ProviderBoardComponent', () => {
  let component: ProviderBoardComponent;
  let fixture: ComponentFixture<ProviderBoardComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ProviderBoardComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ProviderBoardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
