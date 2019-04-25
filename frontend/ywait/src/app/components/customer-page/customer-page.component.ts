import { Component } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { DeliveryService } from 'src/app/services/delivery.service';

export const DUMMY_STORES = [
  'H3K2C3', 'H4N1J4', 'H7P5P4'
]

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

  constructor(private fb: FormBuilder, private deliveryService: DeliveryService) {
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
    this.deliveryService.createOrder(this.form.value.nameFormCtrl, this.form.value.addressFormCtrl, this.form.value.phoneFormCtrl).subscribe(() => {
      console.log("placed order");
    })
  }
}
