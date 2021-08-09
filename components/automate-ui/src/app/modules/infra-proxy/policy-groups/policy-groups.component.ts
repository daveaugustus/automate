import {
  Component,
  Input,
  OnInit,
  OnDestroy,
  EventEmitter,
  Output } from '@angular/core';
import { Store } from '@ngrx/store';
import { combineLatest, Subject } from 'rxjs';
import { takeUntil } from 'rxjs/operators';
import { isNil } from 'lodash/fp';
import { LayoutFacadeService, Sidebar } from 'app/entities/layout/layout.facade';
import { NgrxStateAtom } from 'app/ngrx.reducers';
import { EntityStatus } from 'app/entities/entities';
import { GetPolicyGroups } from 'app/entities/policy-files/policy-file.action';
import { PolicyFile } from 'app/entities/policy-files/policy-file.model';
import {
  getGroupsStatus,
  policyFile
} from 'app/entities/policy-files/policy-group.selectors';

@Component({
  selector: 'app-policy-groups',
  templateUrl: './policy-groups.component.html',
  styleUrls: ['./policy-groups.component.scss']
})
export class PolicyGroupsComponent implements OnInit, OnDestroy {
  @Input() serverId: string;
  @Input() orgId: string;
  @Output() resetKeyRedirection = new EventEmitter<boolean>();

  private isDestroyed = new Subject<boolean>();
  public policyFiles: PolicyFile[] = [];
  public policyGroupsListLoading = true;
  public authFailure = false;
  public searching = false;
  public searchValue = '';
  public searchArr;
  public groupList = [];

  constructor(
    private store: Store<NgrxStateAtom>,
    private layoutFacade: LayoutFacadeService
  ) { }

  ngOnInit(): void {
    this.layoutFacade.showSidebar(Sidebar.Infrastructure);

    this.getPolicyGropus();

    combineLatest([
      this.store.select(getGroupsStatus),
      this.store.select(policyFile)
    ]).pipe(takeUntil(this.isDestroyed))
    .subscribe(([ getPolicyFilesSt, policyFilesState]) => {
      if (getPolicyFilesSt === EntityStatus.loadingSuccess && !isNil(policyFilesState)) {
        this.policyFiles = policyFilesState;
        this.groupList = this.policyFiles;
        this.policyGroupsListLoading = false;
      } else if (getPolicyFilesSt === EntityStatus.loadingFailure) {
        this.policyGroupsListLoading = false;
        this.authFailure = true;
      }
    });
  }

  ngOnDestroy(): void {
    this.isDestroyed.next(true);
    this.isDestroyed.complete();
  }

  resetKeyTabRedirection(resetLink: boolean) {
    this.resetKeyRedirection.emit(resetLink);
  }

  getPolicyGropus() {
    const payload = {
      server_id: this.serverId,
      org_id: this.orgId
    };
    this.store.dispatch(new GetPolicyGroups(payload));
  }

  searchPolicyFiles(searchText: string): void {
    this.searching = true;
    this.searchValue = searchText;
    if (!this.policyFiles || !searchText) {
      this.groupList = this.policyFiles;
    } else {
      const list = this.policyFiles.filter((key) => {
        if (key) {
          return key.policy_group.includes(this.searchValue);
        }
      });
      this.groupList = list;
    }
    this.searchArr = this.groupList;
    this.searching = false;
  }
}