/// <reference types="@types/googlemaps" />

import { Component, OnInit } from '@angular/core';
import { DeliveryService } from 'src/app/services/delivery.service';
import { ActivatedRoute } from '@angular/router';
import { MatSnackBar } from '@angular/material';

import { ViewChild } from '@angular/core';
import { LatLng } from '../customer-page/customer-page.component';

export interface OrderInfo {
  orderId: string,
  storeCoordinates: LatLng,
  deliveryCoordinates: LatLng
}

export interface Order {
  orderId: string,
  status: number,
  customer: {
    customerId: string,
    name: string,
    phone: string,
    deliveryLocation: {
      lat: number,
      lng: number
    }
  },
  storeLocation: {
    lat: number,
    lng: number
  }
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
  accepted: boolean = false;
  isDelivered: boolean = false;
  order: Order;


  @ViewChild('gmap') gmapElement: any;
  map: google.maps.Map;
  latitude:number;
  longitude:number;

  constructor(
    public deliveryService: DeliveryService,
    private route: ActivatedRoute,
    private snackBar: MatSnackBar
  ) { }

  ngOnInit() {
    this.route.queryParams.subscribe(params => {
      this.orderId = params['orderId'];
      this.deliveryId = params['deliveryId']

      console.log("this.orderId" + this.orderId);
      console.log("this.deliveryId" + this.deliveryId);

      this.deliveryService.getOrderInfo(this.orderId).subscribe((order: Order) => {
        this.order = order;

        var mapProp = {
          //center: new google.maps.LatLng(45.5344459,-73.6562009),
          center: new google.maps.LatLng(this.order.customer.deliveryLocation.lat, this.order.customer.deliveryLocation.lng),
          zoom: 15,
          mapTypeId: google.maps.MapTypeId.ROADMAP
        };
        this.map = new google.maps.Map(this.gmapElement.nativeElement, mapProp);
      });
    })

    // var mapProp = {
    //   //center: new google.maps.LatLng(45.5344459,-73.6562009),
    //   center: new google.maps.LatLng(this.order.customer.deliveryLocation.lat, this.order.customer.deliveryLocation.lng),
    //   zoom: 15,
    //   mapTypeId: google.maps.MapTypeId.ROADMAP
    // };
    // this.map = new google.maps.Map(this.gmapElement.nativeElement, mapProp);
  }

  accept() {
    console.log("accept()");
    this.accepted = true;
    this.deliveryService.acceptDelivery(this.orderId, this.deliveryId).subscribe(() => {
      console.log("accepted");
      this.openSnackBar('You have accepted the order!', 'success');
    })
  }

  delivered() {
    console.log("delivered()");
    this.isDelivered = true;
    this.deliveryService.delivered(this.orderId, this.deliveryId).subscribe(() => {
      console.log("delivered");
      this.openSnackBar('You have delivered the order!', 'success');
    })
  }

  setMapType(mapTypeId: string) {
    this.map.setMapTypeId(mapTypeId)
  }

  setCenter(e:any){
    e.preventDefault();
    this.map.setCenter(new google.maps.LatLng(this.latitude, this.longitude));
  }

  openSnackBar(message: string, action: string) {
    this.snackBar.open(message, action, {duration: 5000})
  }
}
