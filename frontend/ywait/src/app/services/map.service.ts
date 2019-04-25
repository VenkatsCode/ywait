import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class MapService {
  baseUrl: string;
  apiKey: string = 'E5IEoVs6K7zHHdDTzbFm5azArH2gaEsk';
  location: string = '1015%20Rue%20Du%20March%C3%A9%20Central';

  constructor(private http: HttpClient) {
    //this.baseUrl = 'http://open.mapquestapi.com/geocoding/v1/address?key=E5IEoVs6K7zHHdDTzbFm5azArH2gaEsk&location=1015%20Rue%20Du%20March%C3%A9%20Central';
    this.baseUrl = 'http://open.mapquestapi.com/geocoding/v1/address';
  }

  getLatLng(address: string): Observable<any> {
    const url = this.baseUrl + '?' + 'key=' + this.apiKey + '&' + 'location=' + encodeURI(address);
    return this.http.get(url);
  }
}
