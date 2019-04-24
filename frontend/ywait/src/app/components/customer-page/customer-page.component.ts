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
      name: ['', Validators.required]
    });
  }

  placeOrder() {
    console.log("placing order");
    console.log("customer name: " + this.form.value.name);
    this.deliveryService.createOrder(this.form.value.name).subscribe(() => {
      console.log("placed order");
    })
  }
}
