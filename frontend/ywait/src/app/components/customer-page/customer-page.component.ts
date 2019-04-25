import { Component } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { DeliveryService } from 'src/app/services/delivery.service';

@Component({
  selector: 'app-customer-page',
  templateUrl: './customer-page.component.html',
  styleUrls: ['./customer-page.component.css']
})
export class CustomerpageComponent {
  form: FormGroup;

  constructor(private fb: FormBuilder, private deliveryService: DeliveryService) {
    this.form = this.fb.group({
      nameFormCtrl: ['', Validators.required],
      addressFormCtrl: ['', Validators.required],
      storeFormCtrl: ['', Validators.required]
    });
  }

  placeOrder() {
    console.log("placing order");
    console.log("customer name: " + this.form.value.nameFormCtrl);
    this.deliveryService.createOrder(this.form.value.nameFormCtrl).subscribe(() => {
      console.log("placed order");
    })
  }
}
