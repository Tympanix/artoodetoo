import { Component, OnInit } from '@angular/core';
import { Meta } from '../meta';
import { Task, Unit } from '../task';
import { UnitService } from '../unit.service';
import { TaskService} from '../task.service';
import { ApiService } from '../api.service';

import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-administration',
  templateUrl: './administration.component.html',
  styles: []
})
export class AdministrationComponent implements OnInit {
  metas: Meta[]
  tasks: Task[]
  event: Unit
  actions: Unit[]
  task: Task

  constructor(private api: ApiService, private taskService: TaskService, private route: ActivatedRoute) {
    api.metas.subscribe((metas) => this.metas = metas)
    api.tasks.subscribe((tasks) => this.tasks = tasks)
  }

  ngOnInit() {
    this.route.data.subscribe((data: {task: Task}) => this.task = data.task)
  }

  // Return units with an input type mathcing the given argument
  getUnitsByType(type: string) {
    let typeUnits: Meta[];
    typeUnits =  this.metas.filter(meta => meta.input.find(x => x.type === type));
    console.log(typeUnits);
    return typeUnits;
  }

  createTask(): void {
    this.taskService.createTask(this.task);
  }

  // For test purpose only
  createMockTask():void {
    this.taskService.createMockTask()
  }


}
