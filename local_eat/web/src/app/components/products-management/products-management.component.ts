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
  selectedCategoryId: string = '';
  newProduct: any = {
    name: '',
    categoryId: '',
    quantity: 0,
    price: 0,
    picture: ''
  };
  constructor(
    private profileService: ProfileService, 
    private productsService: ProductsService, 
    private router: Router
  ) { }

  ngOnInit(): void {
    this.getProducer();
    this.productsService.getCategories().subscribe(data => {
      this.categories = data;
    });
    
  }
  onCategoryChange(event: any): void {
    const categoryId = event.target.value;
    console.log('Selected Category ID:', categoryId); // Log selected category ID
    this.productsService.getProductsByCategory(categoryId).subscribe(data => {
      this.products = data;
      console.log('Products:', this.products); // Verify products are fetched
    });
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

  addNewProduct(productData: any): void {
    this.productsService.addProduct(productData).subscribe(
      (dataProducer) => { 
        this.producer = dataProducer;
        this.toggleAddProductForm();
      },
      (error) => {
        console.error('Error adding new product:', error);
      }
    );
  }
}
