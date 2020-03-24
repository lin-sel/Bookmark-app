import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginComponent } from './component/login/login.component';
import { RegisterComponent } from './component/register/register.component';
import { PageNotFoundComponent } from './component/page-not-found/page-not-found.component';
import { ReactiveFormsModule } from '@angular/forms'
import { HttpClientModule } from '@angular/common/http'
import { Constant } from './service/constant';
import { HomeComponent } from './component/home/home.component';
import { NavbarComponent } from './component/navbar/navbar.component';
import { PopmodelComponent } from './model/popmodel/popmodel.component';
import { ViewComponent } from './model/view/view.component';
import { CategoryComponent } from './component/category/category.component';
import { EditbookmarkComponent } from './component/editbookmark/editbookmark.component';
import { AddbookmarkComponent } from './component/addbookmark/addbookmark.component';


@NgModule({
      declarations: [
            AppComponent,
            LoginComponent,
            RegisterComponent,
            PageNotFoundComponent,
            HomeComponent,
            NavbarComponent,
            PopmodelComponent,
            ViewComponent,
            CategoryComponent,
            EditbookmarkComponent,
            AddbookmarkComponent,
      ],
      imports: [
            BrowserModule,
            AppRoutingModule,
            ReactiveFormsModule,
            HttpClientModule
      ],
      providers: [
            Constant,
      ],
      bootstrap: [AppComponent]
})
export class AppModule { }
