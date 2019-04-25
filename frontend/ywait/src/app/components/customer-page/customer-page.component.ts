import { Component } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { DeliveryService } from 'src/app/services/delivery.service';
import { MatSnackBar } from '@angular/material';
import { MapService } from 'src/app/services/map.service';

export interface LatLng {
  lat: number;
  lng: number;
}

@Component({
  selector: 'app-customer-page',
  templateUrl: './customer-page.component.html',
  styleUrls: ['./customer-page.component.css']
})
export class CustomerpageComponent {
  form: FormGroup;
  stores: string[] = [
    'H3K2C3', 'H4N1J4', 'H7P5P4'
  ];
  latLng: LatLng;
  orderPlaced: boolean = false;
  orderId: string;

  constructor(
    private fb: FormBuilder,
    private deliveryService: DeliveryService,
    private snackBar: MatSnackBar,
    private mapService: MapService
  ) {
    this.form = this.fb.group({
      nameFormCtrl: ['', Validators.required],
      addressFormCtrl: ['', Validators.required],
      phoneFormCtrl: ['', Validators.required]
    });
  }

  placeOrder() {
    console.log("placing order");
    console.log("customer name: " + this.form.value.nameFormCtrl);
    console.log("customer address: " + this.form.value.addressFormCtrl);
    console.log("customer phone: " + this.form.value.phoneFormCtrl);
    this.mapService.getLatLng(this.form.value.addressFormCtrl).subscribe((jsonRes) => {
      console.log("retrieved map coordinates");
      this.latLng = jsonRes.results[0].locations[0].latLng;
      console.log(this.latLng);
      this.deliveryService.createOrder(
        this.form.value.nameFormCtrl,
        this.form.value.addressFormCtrl,
        this.form.value.phoneFormCtrl,
        this.latLng
      ).subscribe((res) => {
        console.log("placed order");
        console.log(res);
        this.orderId = res;
        this.orderPlaced = true;
        this.openSnackBar('Your order was placed successfully!', 'success');
      })
    })
  }

  openSnackBar(message: string, action: string) {
    this.snackBar.open(message, action, {duration: 5000})
  }
}
