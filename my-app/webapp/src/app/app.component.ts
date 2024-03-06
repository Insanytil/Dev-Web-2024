import { Component, OnInit } from '@angular/core';
import { PRODUCERS } from "./mock-producer-list";
import { Producer } from "./models/producer.model";

@Component({
  selector: 'app-root',
  templateUrl: 'app.component.html',
  styleUrls: ['app.component.scss']
})


export class AppComponent implements OnInit{
  title = 'LocalEat';
  producerList : Producer[] = PRODUCERS;
  producerSelected : Producer|undefined;
  ngOnInit(){
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
