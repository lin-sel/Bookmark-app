import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { LoginComponent } from './component/login/login.component';
import { RegisterComponent } from './component/register/register.component';
import { PageNotFoundComponent } from './component/page-not-found/page-not-found.component';
import { HomeComponent } from './component/home/home.component';
import { CategoryComponent } from './component/category/category.component';
import { EditbookmarkComponent } from './component/editbookmark/editbookmark.component';
import { AddbookmarkComponent } from './component/addbookmark/addbookmark.component';



const appRoutes: Routes = [
      { path: 'login', component: LoginComponent },
      { path: 'register', component: RegisterComponent },
      { path: 'bookmark', component: HomeComponent },
      {
            path: '',
            redirectTo: '/login',
            pathMatch: 'full'
      },
      { path: 'addbookmark/:id', component: AddbookmarkComponent },
      { path: 'editbookmark/:id', component: EditbookmarkComponent },
      { path: 'category', component: CategoryComponent },
      { path: '**', component: PageNotFoundComponent }
];

@NgModule({
      imports: [RouterModule.forRoot(appRoutes)],
      exports: [RouterModule]
})
export class AppRoutingModule { }
