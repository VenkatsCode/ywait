/// <reference types="@types/googlemaps" />

import { Component, OnInit } from '@angular/core';
import { DeliveryService } from 'src/app/services/delivery.service';
import { ActivatedRoute } from '@angular/router';

import { ViewChild } from '@angular/core';
// import { } from '@types/googlemaps';

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
  accepted: boolean = false;
  isDelivered: boolean = false;

  @ViewChild('gmap') gmapElement: any;
  map: google.maps.Map;
  latitude:number;
  longitude:number;

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

    var mapProp = {
      center: new google.maps.LatLng(45.5344459,-73.6562009),
      zoom: 15,
      mapTypeId: google.maps.MapTypeId.ROADMAP
    };
    this.map = new google.maps.Map(this.gmapElement.nativeElement, mapProp);
    this.setMapType('roadmap')
  }

  accept() {
    console.log("accept()");
    this.accepted = true;
    this.deliveryService.acceptDelivery(this.orderId, this.deliveryId).subscribe(() => {
      console.log("accepted");
    })
  }

  delivered() {
    console.log("delivered()");
    this.isDelivered = true;
    this.deliveryService.delivered(this.orderId, this.deliveryId).subscribe(() => {
      console.log("delivered");
    })
  }

  setMapType(mapTypeId: string) {
    this.map.setMapTypeId(mapTypeId)
  }

  setCenter(e:any){
    e.preventDefault();
    this.map.setCenter(new google.maps.LatLng(this.latitude, this.longitude));
  }
}
