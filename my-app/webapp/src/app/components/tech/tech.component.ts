import { Component, OnInit } from '@angular/core';
import { Technology } from '../../models/tech.model';
import { TechService } from '../../services/tech.service';

@Component({
  selector: 'app-tech',
  templateUrl: './tech.component.html'
})
export class TechComponent implements OnInit {

  technologies: Technology[] = [];

  constructor(private readonly techService: TechService) { }

  ngOnInit() {
    this.techService.getTechnologies().subscribe(value => {
      this.technologies = value;
    });
  }
}
