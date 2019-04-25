import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { LatLng } from '../components/customer-page/customer-page.component';

@Injectable({
  providedIn: 'root'
})
export class DeliveryService {
  protected baseUrl: string;

  constructor(private http: HttpClient) {
    this.baseUrl = 'http://localhost:3000/order/';
  }

  acceptDelivery(orderId: string, deliveryId: string): Observable<any> {
    const url = this.baseUrl + 'accept';
    return this.http.post(url, {orderId: orderId, deliveryId: deliveryId});
  }

  delivered(orderId: string, deliveryId: string): Observable<any> {
    const url = this.baseUrl + 'delivered';
    return this.http.post(url, {orderId: orderId, deliveryId: deliveryId});
  }

  getOrderInfo(orderId: string): Observable<any> {
    const url = this.baseUrl + orderId;
    return this.http.get(url);
  }

  createOrder(name: string, address: string, phone: string, latLng: LatLng): Observable<any> {
    const url = this.baseUrl + 'place';
    return this.http.post(url, {name: name, address: address, phone: phone, latLng: latLng});
  }
}
