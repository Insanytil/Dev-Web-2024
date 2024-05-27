import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ProfileService } from 'src/app/services/profile.service';
import { Producer } from 'src/app/models/producer.model';
import { ProductsService } from 'src/app/services/products.service';


@Component({
  selector: 'app-products-management',
  templateUrl: './products-management.component.html',
  styleUrls: ['./products-management.component.scss']
})
export class ProductsManagementComponent implements OnInit {
  public producer: Producer | undefined;
  public showAddProductForm: boolean = false;
  categories: any[] = [];
  products: any[] = [];
  productsCompany: any[] = [];
  selectedCategoryId: string = '';
  selectedProductId: string = '';
  constructor(
    private profileService: ProfileService, 
    private productsService: ProductsService, 
    private router: Router
  ) { }
  productId: string = '';
  quantity: number = 0;
  price: number = 0;

  ngOnInit(): void {
    this.getProducer();
    this.productsService.getCategories().subscribe(data => {
      this.categories = data;
    });
    this.getProductsByCompany();
  }
  getProductsByCompany(): void {
    this.productsService.getProductsByCompany().subscribe(data => {
      this.productsCompany = data;
    });
  }
  onCategoryChange(event: any): void {
    const categoryId = event.value;
    this.productsService.getProductsByCategory(categoryId).subscribe(data => {
      this.products = data;
    });
  }
  onProductChange(event: any): void {
    const productName = event.value;
    const selectedProduct = this.products.find(product => product.name === productName);
    this.productId = selectedProduct.id;
  }
  
  getProducer(): void {
    this.profileService.getProducer().subscribe(
      (dataProducer) => {
        this.producer = dataProducer;
      },
      (error) => {
        console.error('Error fetching producer:', error);
      }
    );
  }

  toggleAddProductForm(): void {
    this.showAddProductForm = !this.showAddProductForm;
  }

  addNewProduct(): void {
    this.productsService.addProduct(this.productId, this.quantity, this.price).subscribe(
      (dataProducer) => { 
        this.producer = dataProducer;
        this.toggleAddProductForm();
        this.getProductsByCompany();
      },
      (error) => {
        console.error('Error adding new product:', error);
      }
    );
  }
  getProductImageUrl(picture: string): string {
    return "../../../assets/products/" + picture;
  }
}
