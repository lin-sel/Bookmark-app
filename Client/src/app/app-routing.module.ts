import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { LoginComponent } from './component/login/login.component';
import { RegisterComponent } from './component/register/register.component';
import { PageNotFoundComponent } from './component/page-not-found/page-not-found.component';



const appRoutes: Routes = [
      { path: 'login', component: LoginComponent },
      { path: 'register', component: RegisterComponent },
      // {
      //       path: 'heroes',
      //       component: HeroListComponent,
      //       data: { title: 'Heroes List' }
      // },
      {
            path: '',
            redirectTo: '/login',
            pathMatch: 'full'
      },
      { path: '**', component: PageNotFoundComponent }
];

@NgModule({
      imports: [RouterModule.forRoot(appRoutes)],
      exports: [RouterModule]
})
export class AppRoutingModule { }
