import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ProfileService } from 'src/app/services/profile.service';
import { User } from 'src/app/models/user.model';
import { Company } from 'src/app/models/company.model';
import { HttpResponse } from '@angular/common/http';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.scss']
})
export class ProfileComponent implements OnInit {

  public user: User | undefined;
  public company: Company | undefined;
  public showCreateCompanyForm: boolean = false;
  public showJoinCompanyForm: boolean = false;

  constructor(private profileService: ProfileService, private router: Router) { }
  CompanyName: string = '';
  Password: string = '';
  Alias: string = '';
  Address: string = '';
  Mail: string = '';
  PhoneNum: string = '';
  VATNum: string = '';
  Description: string = '';

  ngOnInit(): void {
    this.getUser();
    this.getCompany();
  }

  getUser() {
    this.profileService.getUser().subscribe(
      (dataUser) => {
        this.user = dataUser;
      },
      (error) => {
        console.error(error);
      }
    );
  }
  
  getCompany(){
    this.profileService.getCompany().subscribe(
      (dataCompany) => {
        this.company = dataCompany;
        console.log(dataCompany)
      },
      (error) => {
        console.error(error);
      }
    );
  }
  createCompany(){
    this.profileService.CreateCompany(this.CompanyName, this.Password, this.Alias, 
    this.Address, this.Mail, this.PhoneNum, this.VATNum, this.Description).subscribe(
      (res: HttpResponse<any>) => {
        if (res.ok) {
          this.router.navigate(['/profil']);
        } else {
          console.error('Error:', res.body.error);
        }
      }
    )
  }
  joinCompany(){

  }
  onFileSelected(event: any) {
    const selectedFile = event.target.files[0];
  }
  toggleCreateCompanyForm() {
    this.showCreateCompanyForm = !this.showCreateCompanyForm;
  }
  toggleJoinCompanyForm() {
    this.showJoinCompanyForm = !this.showJoinCompanyForm;
  }
}
