import { Component } from '@angular/core';
import { takeUntil } from 'rxjs';
import { Producer } from 'src/app/models/producer.model';
import { ProducersService } from 'src/app/services/producers.service';
import { Unsub } from 'src/app/utils/unsub';

@Component({
  selector: 'app-producers',
  templateUrl: './producers.component.html',
  styleUrls: ['./producers.component.scss']
})
export class ProducersComponent extends Unsub{
  constructor(private producersService: ProducersService) {
    super();
  }

  title = 'LocalEat';
  producerList : Producer[] = [];
  producerSelected : Producer|undefined;

  ngOnInit(){
    this.producersService.getProducers().pipe(takeUntil(this.unsubscribe$)).subscribe((producers) => {
        this.producerList = producers;
        this.producerList.forEach(producer => {
            producer.created = new Date(producer.created);
            producer.picture = `assets/${producer.picture}`;
        });
    });
  }
  
  selectProducer(producerId: string) {
      const id = +producerId;
      const producer: Producer|undefined = this.producerList.find(producer => producer.id == +producerId);
      if(producer){
          console.log(`Vous avez cliqué sur le producteur ${producer.name}`);
          this.producerSelected = producer;
      }else{
          console.log(`Producteur inexistant`);
          this.producerSelected = producer;
      }
  }
}
