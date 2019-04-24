import { Component, OnInit } from '@angular/core';
import { DeliveryService } from 'src/app/services/delivery.service';
import { ActivatedRoute } from '@angular/router';

export interface OrderInfo {
  orderId: string,
  storeCoordinates: string,
  deliveryCoordinates: string
}

@Component({
  selector: 'app-delivery-page',
  templateUrl: './delivery-page.component.html',
  styleUrls: ['./delivery-page.component.css']
})
export class DeliverypageComponent implements OnInit {
  orderInfo: OrderInfo;
  orderId: string;
  deliveryId: string;

  constructor(
    public deliveryService: DeliveryService,
    private route: ActivatedRoute
  ) { }

  ngOnInit() {
    this.route.queryParams.subscribe(params => {
      this.orderId = params['orderId'];
      this.deliveryId = params['deliveryId']

      console.log("this.orderId" + this.orderId);
      console.log("this.deliveryId" + this.deliveryId);

      this.deliveryService.getOrderInfo(this.orderId).subscribe((orderInfo: OrderInfo) => {
        console.log("retrieved order info");
        console.log(orderInfo);
        this.orderInfo = orderInfo;
      });
    })
  }

  accept() {
    console.log("accept()");
    this.deliveryService.acceptDelivery(this.orderId, this.deliveryId).subscribe(() => {
      console.log("accepted");
    })
  }

  delivered() {
    console.log("delivered()");
    this.deliveryService.delivered(this.orderId, this.deliveryId).subscribe(() => {
      console.log("delivered");
    })
  }
}
