import { ComponentFixture, TestBed } from '@angular/core/testing';
import { MockComponent } from 'ng2-mock-component';
import { EventEmitter } from '@angular/core';
import { ResetNodeKeyComponent } from './reset-node-key.component';
import { StoreModule, Store } from '@ngrx/store';
import { ngrxReducers, runtimeChecks, NgrxStateAtom } from 'app/ngrx.reducers';
import { HttpErrorResponse } from '@angular/common/http';
import { HttpStatus } from 'app/types/types';
import { ResetKeyClientSuccess,
  ResetKeySuccessPayload,
  ResetKeyClientFailure } from 'app/entities/clients/client.action';

describe('ResetNodeKeyComponent', () => {
  let component: ResetNodeKeyComponent;
  let fixture: ComponentFixture<ResetNodeKeyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        MockComponent({ selector: 'chef-modal', inputs: ['visible'] }),
        MockComponent({ selector: 'chef-button', inputs: ['disabled'] }),
        MockComponent({ selector: 'chef-loading-spinner'}),
        ResetNodeKeyComponent ],
        imports: [
          StoreModule.forRoot(ngrxReducers, { runtimeChecks })
        ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ResetNodeKeyComponent);
    component = fixture.componentInstance;
    component.openEvent = new EventEmitter();
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  describe('#ResetNodeKeyComponentData', () => {
    let store: Store<NgrxStateAtom>;

    const responseData: ResetKeySuccessPayload = {
      name: 'test_name',
      client_key: {
        name: 'default',
        public_key: 'test_public_key',
        expiration_date: 'test_expiration_date',
        private_key: 'test_private_key'
      }
    };

    beforeEach(() => {
      store = TestBed.inject(Store);
    });

    it('hide modal after reset key.', () => {
      component.resetKeyNode();

      store.dispatch(new ResetKeyClientSuccess(
        {name: responseData.name, client_key: responseData.client_key}));
      expect(component.visible).toBe(false);
    });


    it('on reset error, modal is closed (because error is handled by failure banner)', () => {

      const error = <HttpErrorResponse>{
        status: HttpStatus.INTERNAL_SERVER_ERROR,
        ok: false
      };
      store.dispatch(new ResetKeyClientFailure(error));
      expect(component.visible).toBe(false);
    });

  });
});
