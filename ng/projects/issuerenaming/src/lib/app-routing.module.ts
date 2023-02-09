import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { FoosTableComponent } from './foos-table/foos-table.component'
import { FooDetailComponent } from './foo-detail/foo-detail.component'

import { WaldosTableComponent } from './waldos-table/waldos-table.component'
import { WaldoDetailComponent } from './waldo-detail/waldo-detail.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'github_com_fullstack_lang_issuerenaming_go-foos', component: FoosTableComponent, outlet: 'github_com_fullstack_lang_issuerenaming_go_table' },
	{ path: 'github_com_fullstack_lang_issuerenaming_go-foo-adder', component: FooDetailComponent, outlet: 'github_com_fullstack_lang_issuerenaming_go_editor' },
	{ path: 'github_com_fullstack_lang_issuerenaming_go-foo-adder/:id/:originStruct/:originStructFieldName', component: FooDetailComponent, outlet: 'github_com_fullstack_lang_issuerenaming_go_editor' },
	{ path: 'github_com_fullstack_lang_issuerenaming_go-foo-detail/:id', component: FooDetailComponent, outlet: 'github_com_fullstack_lang_issuerenaming_go_editor' },

	{ path: 'github_com_fullstack_lang_issuerenaming_go-waldos', component: WaldosTableComponent, outlet: 'github_com_fullstack_lang_issuerenaming_go_table' },
	{ path: 'github_com_fullstack_lang_issuerenaming_go-waldo-adder', component: WaldoDetailComponent, outlet: 'github_com_fullstack_lang_issuerenaming_go_editor' },
	{ path: 'github_com_fullstack_lang_issuerenaming_go-waldo-adder/:id/:originStruct/:originStructFieldName', component: WaldoDetailComponent, outlet: 'github_com_fullstack_lang_issuerenaming_go_editor' },
	{ path: 'github_com_fullstack_lang_issuerenaming_go-waldo-detail/:id', component: WaldoDetailComponent, outlet: 'github_com_fullstack_lang_issuerenaming_go_editor' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
