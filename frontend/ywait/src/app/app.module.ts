import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { DeliverypageComponent } from './components/delivery-page/delivery-page.component';
import { CustomerpageComponent } from './components/customer-page/customer-page.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { RouterModule, Routes } from '@angular/router';
import { HttpClientModule } from '@angular/common/http';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatButtonModule } from '@angular/material/button';

const appRoutes: Routes = [
  { path: 'order/place', component: CustomerpageComponent },
  { path: 'order', component: DeliverypageComponent },
  { path: '', redirectTo: '/order', pathMatch: 'full' },
];
@NgModule({
  declarations: [
    AppComponent,
    DeliverypageComponent,
    CustomerpageComponent
  ],
  imports: [
    RouterModule.forRoot(appRoutes),
    BrowserModule,
    AppRoutingModule,
    HttpClientModule, 
    FormsModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatButtonModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
