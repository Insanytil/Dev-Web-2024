import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';
import { Product } from '../models/product.model';

@Injectable({
    providedIn: 'root'
})
export class ProductsService {
    private url: string = `${environment.apiUrl}/products`;
    GET_CATEGORY_URL = '/categories';
    POST_NEW_PRODUCT_URL = '/add-product'
    GET_PRODUCT_COMPANY = '/products-by-company'
    constructor(private http: HttpClient) { }

    getProducts(): Observable<Product[]> {
        return this.http.get<Product[]>(this.url);
    }
    addProduct(productId : string, quantity: number, price : number): Observable<any> {
        const productData = {
            "productId": productId,
            "price": price,
            "quantity": quantity,
        };
        const httpOptions = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
            withCredentials: true,
            observe: 'response' as 'response'
        };
        return this.http.post<any>(this.url + this.POST_NEW_PRODUCT_URL, productData, httpOptions);
    }
    getCategories(): Observable<any> {
        return this.http.get(this.url + this.GET_CATEGORY_URL);
    }
    
    getProductsByCategory(categoryId: string): Observable<any> {
        return this.http.get(this.url+`/by-category?categoryId=${categoryId}`);
    }
    getProductsByCompany(): Observable<any> {
        return this.http.get(this.url+ this.GET_PRODUCT_COMPANY);
    }
}
