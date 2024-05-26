import { HttpClient } from '@angular/common/http';
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
    GET_PRODUCTS_BY_CATEGORY_URL = '/products?categoryId='
    POST_NEW_PRODUCT_URL = '/add-product'
    constructor(private http: HttpClient) { }

    getProducts(): Observable<Product[]> {
        return this.http.get<Product[]>(this.url);
    }
    addProduct(productData: any): Observable<any> {
        return this.http.post<any>(this.url + this.POST_NEW_PRODUCT_URL, productData);
    }
    getCategories(): Observable<any> {
        return this.http.get(`${this.url}/categories`);
    }
    
    getProductsByCategory(categoryId: string): Observable<any> {
        return this.http.get(`${this.url}/products?categoryId=${categoryId}`);
    }
}
